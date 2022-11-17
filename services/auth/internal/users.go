package internal

type User struct {
	EMail        string
	PasswordHash string
}

var users []User

func GetUser(email string) (User, bool) {
	for _, user := range users {
		if user.EMail == email {
			return user, true
		}
	}
	return User{}, false
}

func (u *User) ValidatePasswordHash(pswdhash string) bool {
	return u.PasswordHash == pswdhash
}
