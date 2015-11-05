package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
)

type HotelUser struct {
	UserId             int
	Name               string
	Username, Password string
	HashedPassword     []byte
}

func (u *HotelUser) String() string {
	return fmt.Sprintf("HotelUser(%s)", u.Username)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (user *HotelUser) Validate(v *revel.Validation) {
	v.Check(user.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	ValidatePassword(v, user.Password).
		Key("user.Password")

	v.Check(user.Name,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}
