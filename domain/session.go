package domain

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"strings"

	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/lexid"
	"github.com/kpango/fastime"
	"github.com/mojura/enkodo"
	"github.com/segmentio/fasthash/fnv1a"
	"github.com/zeebo/xxh3"

	"street/conf"
	"street/model/mAuth/rqAuth"
	"street/model/mAuth/wcAuth"
)

type Session struct {
	UserId    uint64
	ExpiredAt int64 // in seconds
	Email     string
}

func (u *Session) MarshalEnkodo(enc *enkodo.Encoder) (err error) {
	_ = enc.Uint64(u.UserId)
	_ = enc.Int64(u.ExpiredAt)
	_ = enc.String(u.Email)
	return
}

func (u *Session) UnmarshalEnkodo(dec *enkodo.Decoder) (err error) {
	if u.UserId, err = dec.Uint64(); err != nil {
		return
	}
	if u.ExpiredAt, err = dec.Int64(); err != nil {
		return
	}
	if u.Email, err = dec.String(); err != nil {
		return
	}
	return
}

func createHash(key1, key2 string) string {
	res := xxh3.HashString128(key1 + conf.PROJECT_NAME + key2) // PROJECT_NAME = salt, if you change this, all token will be invalidated
	const x = 256
	return string([]byte{
		byte(res.Hi >> (64 - 8) % x),
		byte(res.Hi >> (64 - 16) % x),
		byte(res.Hi >> (64 - 24) % x),
		byte(res.Hi >> (64 - 32) % x),
		byte(res.Hi >> (64 - 40) % x),
		byte(res.Hi >> (64 - 48) % x),
		byte(res.Hi >> (64 - 56) % x),
		byte(res.Hi >> (64 - 64) % x), // nolint: staticcheck
		byte(res.Lo >> (64 - 8) % x),
		byte(res.Lo >> (64 - 16) % x),
		byte(res.Lo >> (64 - 24) % x),
		byte(res.Lo >> (64 - 32) % x),
		byte(res.Lo >> (64 - 40) % x),
		byte(res.Lo >> (64 - 48) % x),
		byte(res.Lo >> (64 - 56) % x),
		byte(res.Lo >> (64 - 64) % x), // nolint: staticcheck
	})
}

const TokenSeparator = `|`

func (s *Session) Encrypt(userAgent string) string {
	key1 := lexid.NanoID()
	key2 := S.EncodeCB63(fnv1a.HashString64(userAgent), 1)
	block, err := aes.NewCipher([]byte(createHash(key1, key2)))
	L.PanicIf(err, `aes.NewCipher`)
	gcm, err := cipher.NewGCM(block)
	L.PanicIf(err, `cipher.NewGCM`)
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	L.PanicIf(err, `io.ReadFull`)
	buffer := bytes.NewBuffer(nil)
	w := enkodo.NewWriter(buffer)
	err = w.Encode(s)
	L.PanicIf(err, `w.Encode`)
	cipherText := gcm.Seal(nonce, nonce, buffer.Bytes(), nil)
	return key1 + TokenSeparator + hex.EncodeToString(cipherText) + TokenSeparator + key2
}

func (s *Session) Decrypt(sessionToken, userAgent string) bool {
	strs := strings.Split(sessionToken, TokenSeparator)
	tokenLen := len(strs)
	if tokenLen != 3 {
		L.Print(`sessionToken length mismatch: ` + I.ToStr(tokenLen))
		return false
	}
	uaHash := S.EncodeCB63(fnv1a.HashString64(userAgent), 1)
	if strs[2] != uaHash {
		L.Print(`userAgent mismatch: ` + strs[2])
		return false
	}
	data, err := hex.DecodeString(strs[1])
	if L.IsError(err, `hex.DecodeString`) {
		return false
	}
	key := []byte(createHash(strs[0], strs[2]))
	block, err := aes.NewCipher(key)
	if L.IsError(err, `aes.NewCipher`) {
		return false
	}
	gcm, err := cipher.NewGCM(block)
	if L.IsError(err, `cipher.NewGCM`) {
		return false
	}
	noncecSize := gcm.NonceSize()
	if len(data) < noncecSize {
		L.Print(`len(data) < noncecSize`)
		return false
	}
	nonce, cipherText := data[:noncecSize], data[noncecSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if L.IsError(err, `gcm.Open`) {
		return false
	}
	err = enkodo.Unmarshal(plainText, s)
	return !L.IsError(err, `enkodo.Unmarshal`)
}

func (d *Domain) createSession(userId uint64, email string, userAgent string) *wcAuth.SessionsMutator {
	session := wcAuth.NewSessionsMutator(d.AuthOltp)
	session.UserId = userId
	sess := Session{
		UserId:    userId,
		ExpiredAt: fastime.Now().AddDate(0, 0, conf.CookieDays).Unix(),
		Email:     email,
	}
	session.SessionToken = sess.Encrypt(userAgent)
	session.ExpiredAt = sess.ExpiredAt
	return session
}

func (d *Domain) expireSession(token string, out *ResponseCommon) int64 {
	session := wcAuth.NewSessionsMutator(d.AuthOltp)
	session.SessionToken = token
	now := fastime.UnixNow()
	if session.FindBySessionToken() {
		out.SessionToken = conf.CookieLogoutValue
		if session.ExpiredAt > now {
			session.SetExpiredAt(now)
			if !session.DoUpdateBySessionToken() {
				out.SetError(500, ErrUserSessionRemovalFailed)
				return 0
			}
			return now
		}
		return session.ExpiredAt
	}
	return 0
}

const (
	ErrSessionTokenEmpty     = `sessionToken empty`
	ErrSessionTokenInvalid   = `sessionToken invalid`
	ErrSessionTokenExpired   = `sessionToken expired`
	ErrSessionTokenNotFound  = `sessionToken not found`
	ErrSessionTokenLoggedOut = `sessionToken already logged out`
)

func (d *Domain) mustLogin(in RequestCommon, out *ResponseCommon) (res *Session) {
	if in.SessionToken == `` {
		out.SetError(403, ErrSessionTokenEmpty)
		return nil
	}
	defer func() {
		if res == nil {
			// force user to clear cookie
			out.SessionToken = conf.CookieLogoutValue
		}
	}()
	sess := Session{}
	if !sess.Decrypt(in.SessionToken, in.UserAgent) {
		out.SetError(403, ErrSessionTokenInvalid)
		return nil
	}
	now := fastime.UnixNow()
	if sess.ExpiredAt < now {
		out.SetError(403, ErrSessionTokenExpired)
		return nil
	}

	session := rqAuth.NewSessions(d.AuthOltp)
	session.SessionToken = in.SessionToken
	if !session.FindBySessionToken() {
		out.SetError(403, ErrSessionTokenNotFound)
		return nil
	}
	if session.ExpiredAt <= now {
		out.SetError(403, ErrSessionTokenLoggedOut)
		return nil
	}
	return &sess
}
