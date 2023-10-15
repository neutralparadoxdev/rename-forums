package core

type Session struct {
	session string /* json webtoken */
}

func (session *Session) ToString() string {
	return session.session
}

func (session *Session) GetUserId() string {
	return ""
}

func (session *Session) GetUsername() string {
	return ""
}

func (session *Session) IsValid() bool {
	return false
}
