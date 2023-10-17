package core

import (
	"errors"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type Session struct {
	session string /* json webtoken */
	Token   *jwt.Token
}

func NewSession(token string) Session {
	return Session{
		session: token,
	}
}

func (session *Session) ToString() string {
	return session.session
}

func (session *Session) GetUserId() (int64, error) {
	if claims, ok := session.Token.Claims.(jwt.MapClaims); ok && session.Token.Valid {
		res, err := strconv.ParseInt(claims["userId"].(string), 10, 64)
		if err != nil {
			return 0, err
		}

		return res, nil
	} else {
		return 0, errors.New("get_user_id: Could not get userid")
	}
}

func (session *Session) GetUsername() string {
	return ""
}

func (session *Session) IsValid() bool {
	return false
}
