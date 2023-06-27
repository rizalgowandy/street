package presentation

// Code generated by 1_codegen_test.go DO NOT EDIT.

import (
	"github.com/gofiber/fiber/v2"

	"street/domain"
)

func ApiRoutes(fw *fiber.App, d *domain.Domain) {

	// AdminDashboard
	fw.Post("/"+domain.AdminDashboardAction, func(c *fiber.Ctx) error {
		in := domain.AdminDashboardIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.AdminDashboardAction); err != nil {
			return err
		}
		out := d.AdminDashboard(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// AdminProperties
	fw.Post("/"+domain.AdminPropertiesAction, func(c *fiber.Ctx) error {
		in := domain.AdminPropertiesIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.AdminPropertiesAction); err != nil {
			return err
		}
		out := d.AdminProperties(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// AdminUsers
	fw.Post("/"+domain.AdminUsersAction, func(c *fiber.Ctx) error {
		in := domain.AdminUsersIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.AdminUsersAction); err != nil {
			return err
		}
		out := d.AdminUsers(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestDebug
	fw.Post("/"+domain.GuestDebugAction, func(c *fiber.Ctx) error {
		in := domain.GuestDebugIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestDebugAction); err != nil {
			return err
		}
		out := d.GuestDebug(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestExternalAuth
	fw.Post("/"+domain.GuestExternalAuthAction, func(c *fiber.Ctx) error {
		in := domain.GuestExternalAuthIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestExternalAuthAction); err != nil {
			return err
		}
		out := d.GuestExternalAuth(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestForgotPassword
	fw.Post("/"+domain.GuestForgotPasswordAction, func(c *fiber.Ctx) error {
		in := domain.GuestForgotPasswordIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestForgotPasswordAction); err != nil {
			return err
		}
		out := d.GuestForgotPassword(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestLogin
	fw.Post("/"+domain.GuestLoginAction, func(c *fiber.Ctx) error {
		in := domain.GuestLoginIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestLoginAction); err != nil {
			return err
		}
		out := d.GuestLogin(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestOauthCallback
	fw.Post("/"+domain.GuestOauthCallbackAction, func(c *fiber.Ctx) error {
		in := domain.GuestOauthCallbackIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestOauthCallbackAction); err != nil {
			return err
		}
		out := d.GuestOauthCallback(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestRegister
	fw.Post("/"+domain.GuestRegisterAction, func(c *fiber.Ctx) error {
		in := domain.GuestRegisterIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestRegisterAction); err != nil {
			return err
		}
		out := d.GuestRegister(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestResendVerificationEmail
	fw.Post("/"+domain.GuestResendVerificationEmailAction, func(c *fiber.Ctx) error {
		in := domain.GuestResendVerificationEmailIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestResendVerificationEmailAction); err != nil {
			return err
		}
		out := d.GuestResendVerificationEmail(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestResetPassword
	fw.Post("/"+domain.GuestResetPasswordAction, func(c *fiber.Ctx) error {
		in := domain.GuestResetPasswordIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestResetPasswordAction); err != nil {
			return err
		}
		out := d.GuestResetPassword(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestVerifyEmail
	fw.Post("/"+domain.GuestVerifyEmailAction, func(c *fiber.Ctx) error {
		in := domain.GuestVerifyEmailIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestVerifyEmailAction); err != nil {
			return err
		}
		out := d.GuestVerifyEmail(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserChangePassword
	fw.Post("/"+domain.UserChangePasswordAction, func(c *fiber.Ctx) error {
		in := domain.UserChangePasswordIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserChangePasswordAction); err != nil {
			return err
		}
		out := d.UserChangePassword(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserDeactivate
	fw.Post("/"+domain.UserDeactivateAction, func(c *fiber.Ctx) error {
		in := domain.UserDeactivateIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserDeactivateAction); err != nil {
			return err
		}
		out := d.UserDeactivate(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserLogout
	fw.Post("/"+domain.UserLogoutAction, func(c *fiber.Ctx) error {
		in := domain.UserLogoutIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserLogoutAction); err != nil {
			return err
		}
		out := d.UserLogout(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserProfile
	fw.Post("/"+domain.UserProfileAction, func(c *fiber.Ctx) error {
		in := domain.UserProfileIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserProfileAction); err != nil {
			return err
		}
		out := d.UserProfile(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserSearchProp
	fw.Post("/"+domain.UserSearchPropAction, func(c *fiber.Ctx) error {
		in := domain.UserSearchPropIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserSearchPropAction); err != nil {
			return err
		}
		out := d.UserSearchProp(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserUpdateProfile
	fw.Post("/"+domain.UserUpdateProfileAction, func(c *fiber.Ctx) error {
		in := domain.UserUpdateProfileIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserUpdateProfileAction); err != nil {
			return err
		}
		out := d.UserUpdateProfile(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

}

// Code generated by 1_codegen_test.go DO NOT EDIT.
