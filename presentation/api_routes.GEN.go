package presentation

// Code generated by 1_codegen_test.go DO NOT EDIT.

import (
	"github.com/gofiber/fiber/v2"

	"street/domain"
)

func ApiRoutes(fw *fiber.App, d *domain.Domain) {

	// AdminAccessLogs
	fw.Post("/"+domain.AdminAccessLogsAction, func(c *fiber.Ctx) error {
		in := domain.AdminAccessLogsIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.AdminAccessLogsAction); err != nil {
			return nil
		}
		out := d.AdminAccessLogs(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// AdminDashboard
	fw.Post("/"+domain.AdminDashboardAction, func(c *fiber.Ctx) error {
		in := domain.AdminDashboardIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.AdminDashboardAction); err != nil {
			return nil
		}
		out := d.AdminDashboard(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// AdminFiles
	fw.Post("/"+domain.AdminFilesAction, func(c *fiber.Ctx) error {
		in := domain.AdminFilesIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.AdminFilesAction); err != nil {
			return nil
		}
		out := d.AdminFiles(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// AdminPropHistories
	fw.Post("/"+domain.AdminPropHistoriesAction, func(c *fiber.Ctx) error {
		in := domain.AdminPropHistoriesIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.AdminPropHistoriesAction); err != nil {
			return nil
		}
		out := d.AdminPropHistories(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// AdminProperties
	fw.Post("/"+domain.AdminPropertiesAction, func(c *fiber.Ctx) error {
		in := domain.AdminPropertiesIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.AdminPropertiesAction); err != nil {
			return nil
		}
		out := d.AdminProperties(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// AdminPropertiesUS
	fw.Post("/"+domain.AdminPropertiesUSAction, func(c *fiber.Ctx) error {
		in := domain.AdminPropertiesUSIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.AdminPropertiesUSAction); err != nil {
			return nil
		}
		out := d.AdminPropertiesUS(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// AdminUsers
	fw.Post("/"+domain.AdminUsersAction, func(c *fiber.Ctx) error {
		in := domain.AdminUsersIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.AdminUsersAction); err != nil {
			return nil
		}
		out := d.AdminUsers(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestAutoLogin
	fw.Post("/"+domain.GuestAutoLoginAction, func(c *fiber.Ctx) error {
		in := domain.GuestAutoLoginIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestAutoLoginAction); err != nil {
			return nil
		}
		out := d.GuestAutoLogin(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestDebug
	fw.Post("/"+domain.GuestDebugAction, func(c *fiber.Ctx) error {
		in := domain.GuestDebugIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestDebugAction); err != nil {
			return nil
		}
		out := d.GuestDebug(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestExternalAuth
	fw.Post("/"+domain.GuestExternalAuthAction, func(c *fiber.Ctx) error {
		in := domain.GuestExternalAuthIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestExternalAuthAction); err != nil {
			return nil
		}
		out := d.GuestExternalAuth(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestFiles
	fw.Post("/"+domain.GuestFilesAction, func(c *fiber.Ctx) error {
		in := domain.GuestFilesIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestFilesAction); err != nil {
			return nil
		}
		out := d.GuestFiles(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestForgotPassword
	fw.Post("/"+domain.GuestForgotPasswordAction, func(c *fiber.Ctx) error {
		in := domain.GuestForgotPasswordIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestForgotPasswordAction); err != nil {
			return nil
		}
		out := d.GuestForgotPassword(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestLogin
	fw.Post("/"+domain.GuestLoginAction, func(c *fiber.Ctx) error {
		in := domain.GuestLoginIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestLoginAction); err != nil {
			return nil
		}
		out := d.GuestLogin(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestOauthCallback
	fw.Post("/"+domain.GuestOauthCallbackAction, func(c *fiber.Ctx) error {
		in := domain.GuestOauthCallbackIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestOauthCallbackAction); err != nil {
			return nil
		}
		out := d.GuestOauthCallback(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestProperty
	fw.Post("/"+domain.GuestPropertyAction, func(c *fiber.Ctx) error {
		in := domain.GuestPropertyIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestPropertyAction); err != nil {
			return nil
		}
		out := d.GuestProperty(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestRegister
	fw.Post("/"+domain.GuestRegisterAction, func(c *fiber.Ctx) error {
		in := domain.GuestRegisterIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestRegisterAction); err != nil {
			return nil
		}
		out := d.GuestRegister(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestResendVerificationEmail
	fw.Post("/"+domain.GuestResendVerificationEmailAction, func(c *fiber.Ctx) error {
		in := domain.GuestResendVerificationEmailIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestResendVerificationEmailAction); err != nil {
			return nil
		}
		out := d.GuestResendVerificationEmail(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestResetPassword
	fw.Post("/"+domain.GuestResetPasswordAction, func(c *fiber.Ctx) error {
		in := domain.GuestResetPasswordIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestResetPasswordAction); err != nil {
			return nil
		}
		out := d.GuestResetPassword(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestVerifyEmail
	fw.Post("/"+domain.GuestVerifyEmailAction, func(c *fiber.Ctx) error {
		in := domain.GuestVerifyEmailIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestVerifyEmailAction); err != nil {
			return nil
		}
		out := d.GuestVerifyEmail(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// RealtorOwnedProperties
	fw.Post("/"+domain.RealtorOwnedPropertiesAction, func(c *fiber.Ctx) error {
		in := domain.RealtorOwnedPropertiesIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.RealtorOwnedPropertiesAction); err != nil {
			return nil
		}
		out := d.RealtorOwnedProperties(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// RealtorProperty
	fw.Post("/"+domain.RealtorPropertyAction, func(c *fiber.Ctx) error {
		in := domain.RealtorPropertyIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.RealtorPropertyAction); err != nil {
			return nil
		}
		out := d.RealtorProperty(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// RealtorUpsertProperty
	fw.Post("/"+domain.RealtorUpsertPropertyAction, func(c *fiber.Ctx) error {
		in := domain.RealtorUpsertPropertyIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.RealtorUpsertPropertyAction); err != nil {
			return nil
		}
		out := d.RealtorUpsertProperty(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserAutoLoginLink
	fw.Post("/"+domain.UserAutoLoginLinkAction, func(c *fiber.Ctx) error {
		in := domain.UserAutoLoginLinkIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserAutoLoginLinkAction); err != nil {
			return nil
		}
		out := d.UserAutoLoginLink(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserChangePassword
	fw.Post("/"+domain.UserChangePasswordAction, func(c *fiber.Ctx) error {
		in := domain.UserChangePasswordIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserChangePasswordAction); err != nil {
			return nil
		}
		out := d.UserChangePassword(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserDeactivate
	fw.Post("/"+domain.UserDeactivateAction, func(c *fiber.Ctx) error {
		in := domain.UserDeactivateIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserDeactivateAction); err != nil {
			return nil
		}
		out := d.UserDeactivate(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserGpsCountry
	fw.Post("/"+domain.UserGpsCountryAction, func(c *fiber.Ctx) error {
		in := domain.UserGpsCountryIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserGpsCountryAction); err != nil {
			return nil
		}
		out := d.UserGpsCountry(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserLikeProp
	fw.Post("/"+domain.UserLikePropAction, func(c *fiber.Ctx) error {
		in := domain.UserLikePropIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserLikePropAction); err != nil {
			return nil
		}
		out := d.UserLikeProp(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserLogout
	fw.Post("/"+domain.UserLogoutAction, func(c *fiber.Ctx) error {
		in := domain.UserLogoutIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserLogoutAction); err != nil {
			return nil
		}
		out := d.UserLogout(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserNearbyFacilities
	fw.Post("/"+domain.UserNearbyFacilitiesAction, func(c *fiber.Ctx) error {
		in := domain.UserNearbyFacilitiesIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserNearbyFacilitiesAction); err != nil {
			return nil
		}
		out := d.UserNearbyFacilities(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserProfile
	fw.Post("/"+domain.UserProfileAction, func(c *fiber.Ctx) error {
		in := domain.UserProfileIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserProfileAction); err != nil {
			return nil
		}
		out := d.UserProfile(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserPropHistory
	fw.Post("/"+domain.UserPropHistoryAction, func(c *fiber.Ctx) error {
		in := domain.UserPropHistoryIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserPropHistoryAction); err != nil {
			return nil
		}
		out := d.UserPropHistory(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserProperty
	fw.Post("/"+domain.UserPropertyAction, func(c *fiber.Ctx) error {
		in := domain.UserPropertyIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserPropertyAction); err != nil {
			return nil
		}
		out := d.UserProperty(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserSearchProp
	fw.Post("/"+domain.UserSearchPropAction, func(c *fiber.Ctx) error {
		in := domain.UserSearchPropIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserSearchPropAction); err != nil {
			return nil
		}
		out := d.UserSearchProp(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserSearchPropUS
	fw.Post("/"+domain.UserSearchPropUSAction, func(c *fiber.Ctx) error {
		in := domain.UserSearchPropUSIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserSearchPropUSAction); err != nil {
			return nil
		}
		out := d.UserSearchPropUS(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserSessionKill
	fw.Post("/"+domain.UserSessionKillAction, func(c *fiber.Ctx) error {
		in := domain.UserSessionKillIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserSessionKillAction); err != nil {
			return nil
		}
		out := d.UserSessionKill(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserSessionsActive
	fw.Post("/"+domain.UserSessionsActiveAction, func(c *fiber.Ctx) error {
		in := domain.UserSessionsActiveIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserSessionsActiveAction); err != nil {
			return nil
		}
		out := d.UserSessionsActive(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserUpdateProfile
	fw.Post("/"+domain.UserUpdateProfileAction, func(c *fiber.Ctx) error {
		in := domain.UserUpdateProfileIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserUpdateProfileAction); err != nil {
			return nil
		}
		out := d.UserUpdateProfile(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserUploadFile
	fw.Post("/"+domain.UserUploadFileAction, func(c *fiber.Ctx) error {
		in := domain.UserUploadFileIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserUploadFileAction); err != nil {
			return nil
		}
		out := d.UserUploadFile(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

}

// Code generated by 1_codegen_test.go DO NOT EDIT.
