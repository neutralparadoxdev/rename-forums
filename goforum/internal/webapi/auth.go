package webapi

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/neutralparadoxdev/rename-forums/goforum/internal/core"
)

var (
	WebApiErrorTokenInvalid          = NewWebApiError(400, "token is invalid")
	WebApiErrorUserAuthRequired      = NewWebApiError(403, "user authentication level required")
	WebApiErrorServerError           = NewWebApiError(404, "server error")
	WebApiErrorTokenValidationFailed = NewWebApiError(405, "token validation failed")
)

func CheckForSession(c *fiber.Ctx, sessionManager *core.SessionManager) (*core.Session, *WebApiError) {
	headers := c.GetReqHeaders()
	tokenString, exists := headers["Bearer-Token"]

	if !exists {
		return nil, nil
	}

	return TokenToSession(tokenString, sessionManager)
}

func TokenToSession(token string, sessionManager *core.SessionManager) (*core.Session, *WebApiError) {
	session := core.NewSession(token)
	ok, err := sessionManager.VerifySession(&session)

	if err != nil {
		if err.Code == core.CoreErrorAuthenticatorFailedToVerify.Code ||
			err.Code == core.CoreErrorSessionManagerUnexpectedSigningMethod.Code ||
			err.Code == core.CoreErrorSessionManagerParsingFailed.Code ||
			err.Code == core.CoreErrorSessionManagerInvalidToken.Code {
			log.Printf("check_session: %s", err)
			return nil, &WebApiErrorTokenValidationFailed

		} else {
			return nil, &WebApiErrorServerError
		}
	}

	if !ok {
		return nil, &WebApiErrorTokenInvalid
	}

	return &session, nil
}
