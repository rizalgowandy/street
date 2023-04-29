package presentation

import (
	"fmt"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"

	"street/domain"
)

func cmdRun(b *domain.Domain, action string, payload []byte) {
	switch action {

	case domain.GuestLoginAction:
		in := domain.GuestLoginIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.GuestLogin(&in)
		fmt.Println(X.ToJsonPretty(out))


	case domain.GuestRegisterAction:
		in := domain.GuestRegisterIn{}
		if L.IsError(json.Unmarshal(payload, &in), "json.Unmarshal") {
			return
		}
		out := b.GuestRegister(&in)
		fmt.Println(X.ToJsonPretty(out))

	}
}
