package core

import (
	"errors"
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
				"nbf":      time.Now(),
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
