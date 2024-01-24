package core

import (
	"errors"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
)

type Session struct {
	Session string /* json webtoken */
	Token   *jwt.Token
}

func NewSession(token string) Session {
	return Session{
		Session: token,
	}
}

func (session *Session) ToString() string {
	return session.Session
}

func (session *Session) GetUserId() (int64, error) {

	if session.Token == nil {
		return 0, errors.New("session::GetUserId: session.Token is nil")
	} 
	claims, ok := session.Token.Claims.(jwt.MapClaims) 

	if !ok {
		return 0, errors.New("session::GetUserId: Token::Claims not ok")
	}

	if !session.Token.Valid {
		return 0, errors.New("session::GetUserId: Not Valid")
	}


	if ok && session.Token.Valid {
		res, err := strconv.ParseInt(claims["userId"].(string), 10, 64)
		if err != nil {
			return 0, err
		}

		return res, nil
	} else {
		return 0, errors.New("get_user_id: Could not get userid")
	}
}

func (session *Session) GetUsername() (string, error) {
	if claims, ok := session.Token.Claims.(jwt.MapClaims); ok && session.Token.Valid {
		res, exists := claims["username"].(string)
		if !exists {
			return "", errors.New("get_user_name: username is not in token")
		}

		return res, nil
	} else {
		return "", errors.New("get_user_name: could not get username")
	}
}

func (session *Session) IsValid() bool {
	return false
}
