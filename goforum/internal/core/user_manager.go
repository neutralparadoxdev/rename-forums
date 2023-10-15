package core

type UserManager struct {
	db Database
}

func NewUserManager(db Database) *UserManager {
	return &UserManager{
		db: db,
	}
}

func (man *UserManager) GetUserByName(username string) (*User, error) {
	rep := man.db.GetUserRepository()
	user, err := rep.GetByName(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}
