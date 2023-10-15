package core

type User struct {
	Username     string
	Password     string
	UserId       string
	Email        string
	CreatedAt    string
	LastModified string
	LastLogin    string
	Locked       bool
}

func DefaultUser() User {
	return User{
		Username:     "Something",
		Password:     "Else",
		Email:        "Somethingelse@example.com",
		CreatedAt:    "Not Today",
		LastModified: "Yesterday",
		LastLogin:    "Today",
		UserId:       "4324i23ou4o23iu4oi",
		Locked:       false,
	}
}
