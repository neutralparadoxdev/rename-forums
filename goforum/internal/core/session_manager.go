package core

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type SessionManager struct {
	sessionRepository SessionRepository
	userManager       *UserManager
	auth              *Authenticator
	secret            string
}

func NewSessionManager(
	secret string,
	userManager *UserManager,
	authenticator *Authenticator,
	sessionRepository SessionRepository) *SessionManager {
	return &SessionManager{
		secret:            secret,
		userManager:       userManager,
		auth:              authenticator,
		sessionRepository: sessionRepository,
	}
}

func (man *SessionManager) CreateSession(username string, password string) (Session, error) {
	if user, err := man.userManager.GetUserByName(username); err == nil {
		ok, err := man.auth.Check(password, user.Password)

		if err != nil {
			return Session{}, errors.New("server error: argon2 failed to check")
		}
		if ok {
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": user.Username,
				"userId":   user.UserId,
				"nbf":      time.Now().Unix(),
				"exp":      time.Now().Add(time.Hour * 24).Unix(),
			})

			tokenString, err := token.SignedString([]byte(man.secret))

			if err != nil {
				return Session{}, errors.New("create_session:SignedToken:Wrong key type")
			}

			session := Session{session: tokenString}

			man.sessionRepository.Save(session)

			return session, nil
		} else {
			return Session{}, errors.New("could Not Create Session")
		}
	}
	return Session{}, nil
}

func (man *SessionManager) DeleteSession(session Session) error {
	token, err := jwt.Parse(session.session, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(man.secret), nil
	})

	if err != nil {
		log.Printf("DeleteSession Error: %v", err)
		return err
	}

	if !token.Valid {
		return errors.New("not a valid token. potentially an attacker")
	}
	res, err := man.sessionRepository.DoesSessionExist(session)

	if err != nil {
		return err
	}

	if res {
		man.sessionRepository.Delete(session)
		return nil
	} else {
		return errors.New("session: cannot delete a session that doesn't exist")
	}
}
