package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"
)


// Code generated by 1_codegen_test.go DO NOT EDIT.


var viewList = map[string]string{
	`Admin`: `../svelte/admin.html`, // ../svelte/admin.svelte
	`AdminAccessLog`: `../svelte/admin/accessLog.html`, // ../svelte/admin/accessLog.svelte
	`AdminFiles`: `../svelte/admin/files.html`, // ../svelte/admin/files.svelte
	`AdminPropHistories`: `../svelte/admin/propHistories.html`, // ../svelte/admin/propHistories.svelte
	`AdminProperties`: `../svelte/admin/properties.html`, // ../svelte/admin/properties.svelte
	`AdminPropertiesUS`: `../svelte/admin/propertiesUS.html`, // ../svelte/admin/propertiesUS.svelte
	`AdminProperty`: `../svelte/admin/property.html`, // ../svelte/admin/property.svelte
	`AdminSessions`: `../svelte/admin/sessions.html`, // ../svelte/admin/sessions.svelte
	`AdminUsers`: `../svelte/admin/users.html`, // ../svelte/admin/users.svelte
	`Buyer`: `../svelte/buyer.html`, // ../svelte/buyer.svelte
	`Debug`: `../svelte/debug.html`, // ../svelte/debug.svelte
	`Error`: `../svelte/error.html`, // ../svelte/error.svelte
	`GuestOauthCallback`: `../svelte/guest/oauthCallback.html`, // ../svelte/guest/oauthCallback.svelte
	`GuestPropertyPublic`: `../svelte/guest/property/public.html`, // ../svelte/guest/property/public.svelte
	`GuestResetPassword`: `../svelte/guest/resetPassword.html`, // ../svelte/guest/resetPassword.svelte
	`GuestVerifyEmail`: `../svelte/guest/verifyEmail.html`, // ../svelte/guest/verifyEmail.svelte
	`Index`: `../svelte/index.html`, // ../svelte/index.svelte
	`Privacy`: `../svelte/privacy.html`, // ../svelte/privacy.svelte
	`Realtor`: `../svelte/realtor.html`, // ../svelte/realtor.svelte
	`RealtorOwnedProperty`: `../svelte/realtor/ownedProperty.html`, // ../svelte/realtor/ownedProperty.svelte
	`RealtorProperty`: `../svelte/realtor/property.html`, // ../svelte/realtor/property.svelte
	`Tos`: `../svelte/tos.html`, // ../svelte/tos.svelte
	`User`: `../svelte/user.html`, // ../svelte/user.svelte
	`UserPropertyIndex`: `../svelte/user/property/index.html`, // ../svelte/user/property/index.svelte
}


func (v *Views) RenderAdmin(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Admin`].Str(m))
}

func (v *Views) RenderAdminAccessLog(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`AdminAccessLog`].Str(m))
}

func (v *Views) RenderAdminFiles(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`AdminFiles`].Str(m))
}

func (v *Views) RenderAdminPropHistories(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`AdminPropHistories`].Str(m))
}

func (v *Views) RenderAdminProperties(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`AdminProperties`].Str(m))
}

func (v *Views) RenderAdminPropertiesUS(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`AdminPropertiesUS`].Str(m))
}

func (v *Views) RenderAdminProperty(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`AdminProperty`].Str(m))
}

func (v *Views) RenderAdminSessions(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`AdminSessions`].Str(m))
}

func (v *Views) RenderAdminUsers(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`AdminUsers`].Str(m))
}

func (v *Views) RenderBuyer(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Buyer`].Str(m))
}

func (v *Views) RenderDebug(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Debug`].Str(m))
}

func (v *Views) RenderError(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Error`].Str(m))
}

func (v *Views) RenderGuestOauthCallback(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`GuestOauthCallback`].Str(m))
}

func (v *Views) RenderGuestPropertyPublic(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`GuestPropertyPublic`].Str(m))
}

func (v *Views) RenderGuestResetPassword(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`GuestResetPassword`].Str(m))
}

func (v *Views) RenderGuestVerifyEmail(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`GuestVerifyEmail`].Str(m))
}

func (v *Views) RenderIndex(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Index`].Str(m))
}

func (v *Views) RenderPrivacy(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Privacy`].Str(m))
}

func (v *Views) RenderRealtor(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Realtor`].Str(m))
}

func (v *Views) RenderRealtorOwnedProperty(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`RealtorOwnedProperty`].Str(m))
}

func (v *Views) RenderRealtorProperty(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`RealtorProperty`].Str(m))
}

func (v *Views) RenderTos(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Tos`].Str(m))
}

func (v *Views) RenderUser(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`User`].Str(m))
}

func (v *Views) RenderUserPropertyIndex(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`UserPropertyIndex`].Str(m))
}

// Code generated by 1_codegen_test.go DO NOT EDIT.
