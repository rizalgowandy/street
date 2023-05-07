package domain

import (
	"testing"

	"github.com/kokizzu/id64"
	"github.com/kpango/fastime"
	"github.com/stretchr/testify/require"
	"github.com/zeebo/assert"
)

func TestLogout(t *testing.T) {
	d := Domain{
		AuthOltp: testTt,
		AuthOlap: testCh,
	}

	email := id64.SID() + `@login`
	const pass = `012345678901`

	// register first
	in := GuestRegisterIn{
		Email:    email,
		Password: pass,
	}
	out := d.GuestRegister(&in)
	assert.Equal(t, out.Error, "")
	require.NotZero(t, out.User.Id)

	t.Run(`emptySessionToken`, func(t *testing.T) {
		in := UserProfileIn{}
		out := d.UserProfile(&in)
		assert.Equal(t, out.Error, ErrSessionTokenEmpty)
	})

	t.Run(`invalidSessionToken`, func(t *testing.T) {
		in := UserProfileIn{
			RequestCommon: RequestCommon{
				SessionToken: `invalid`,
			},
		}
		out := d.UserProfile(&in)
		assert.Equal(t, out.Error, ErrSessionTokenInvalid)
	})

	t.Run(`expiredSessionToken`, func(t *testing.T) {
		const userAgent = ``
		sess := Session{
			UserId:    1,
			ExpiredAt: fastime.UnixNow() - 1,
			Email:     email,
		}

		in := UserProfileIn{
			RequestCommon: RequestCommon{
				SessionToken: sess.Encrypt(userAgent),
			},
		}
		out := d.UserProfile(&in)
		assert.Equal(t, out.Error, ErrSessionTokenExpired)
	})

	t.Run(`login`, func(t *testing.T) {
		in := GuestLoginIn{
			Email:    email,
			Password: pass,
		}
		out := d.GuestLogin(&in)
		assert.Equal(t, out.Error, "")
		require.NotZero(t, out.User.Id)
		require.NotEmpty(t, out.SessionToken)

		sessionToken := out.SessionToken

		// check profile after login must succeed
		t.Run(`profile`, func(t *testing.T) {
			in := &UserProfileIn{
				RequestCommon: RequestCommon{
					SessionToken: sessionToken,
				},
			}
			out := d.UserProfile(in)
			assert.Equal(t, out.Error, "")
			if out.User != nil {
				require.NotZero(t, out.User.Id)
			}
		})

		t.Run(`wrongUserAgent`, func(t *testing.T) {
			in := &UserProfileIn{
				RequestCommon: RequestCommon{
					SessionToken: sessionToken,
					UserAgent:    `otherUserAgent`,
				},
			}
			out := d.UserProfile(in)
			assert.Equal(t, out.Error, ErrSessionTokenInvalid)
		})

		t.Run(`logout`, func(t *testing.T) {
			in := &UserLogoutIn{
				RequestCommon: RequestCommon{
					SessionToken: sessionToken,
				},
			}
			out := d.UserLogout(in)
			assert.Equal(t, out.Error, "")
			require.NotZero(t, out.LogoutAt)
			previousLogoutAt := out.LogoutAt

			t.Run(`profilAfterLogout`, func(t *testing.T) {
				in := &UserProfileIn{
					RequestCommon: RequestCommon{
						SessionToken: sessionToken,
					},
				}
				out := d.UserProfile(in)
				assert.Equal(t, out.Error, ErrSessionTokenLoggedOut)
			})

			// idempotent
			t.Run(`logoutAgain`, func(t *testing.T) {
				out := d.UserLogout(in)
				assert.Equal(t, out.Error, ``)
				assert.Equal(t, out.LogoutAt, previousLogoutAt)
			})
		})
	})
}
