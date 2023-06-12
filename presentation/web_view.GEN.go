package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"
)


// Code generated by 1_codegen_test.go DO NOT EDIT.


var viewList = map[string]string{
	`Admin`: `../svelte/admin.html`, // ../svelte/admin.svelte
	`Buyer`: `../svelte/buyer.html`, // ../svelte/buyer.svelte
	`GuestOauthCallback`: `../svelte/guest/oauthCallback.html`, // ../svelte/guest/oauthCallback.svelte
	`GuestResetPassword`: `../svelte/guest/resetPassword.html`, // ../svelte/guest/resetPassword.svelte
	`GuestVerifyEmail`: `../svelte/guest/verifyEmail.html`, // ../svelte/guest/verifyEmail.svelte
	`Index`: `../svelte/index.html`, // ../svelte/index.svelte
	`Privacy`: `../svelte/privacy.html`, // ../svelte/privacy.svelte
	`Realtor`: `../svelte/realtor.html`, // ../svelte/realtor.svelte
	`Tos`: `../svelte/tos.html`, // ../svelte/tos.svelte
	`User`: `../svelte/user.html`, // ../svelte/user.svelte
}


func (v *Views) RenderAdmin(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(v.cache[`Admin`].Str(m))
}

func (v *Views) RenderBuyer(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(v.cache[`Buyer`].Str(m))
}

func (v *Views) RenderGuestOauthCallback(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(v.cache[`GuestOauthCallback`].Str(m))
}

func (v *Views) RenderGuestResetPassword(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(v.cache[`GuestResetPassword`].Str(m))
}

func (v *Views) RenderGuestVerifyEmail(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(v.cache[`GuestVerifyEmail`].Str(m))
}

func (v *Views) RenderIndex(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(v.cache[`Index`].Str(m))
}

func (v *Views) RenderPrivacy(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(v.cache[`Privacy`].Str(m))
}

func (v *Views) RenderRealtor(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(v.cache[`Realtor`].Str(m))
}

func (v *Views) RenderTos(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(v.cache[`Tos`].Str(m))
}

func (v *Views) RenderUser(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(v.cache[`User`].Str(m))
}

// Code generated by 1_codegen_test.go DO NOT EDIT.
