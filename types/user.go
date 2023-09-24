package types

import "golang.org/x/crypto/bcrypt"


type User struct {
	ID          string `bson:"_id,omitempty" json: "id"`
	Email         string `bson:"_id,omitempty" json: "id"`
	EncryptedPassword   string `bson:"_id,omitempty" json: "id"`
	IsAdmin             string `bson:"isAdmin" json: "isAdmin"`
	Token                 string `bson:"token" json: "token"`
}


// func NewAdminUser(email, password string) (*User, error ){
// 	user, err := NewUser(email, password)
// 	if err != nil {
// 		return nil, err
// 	}
// }

func  NewUser(email, password string) (*User, error) {
	epw,err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return  &User {
		Email: email,
		EncryptedPassword:  string(epw),
	}, nil
}

func (u *User) ValidatePassword(pw string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword),[]byte(pw))
	return err == nil
}