package userSchema

import (
	"time"
)

type User struct {
	Name       string    `json:"name,omitempty" validate:"required"`
	Phone      string    `json:"phone,omitempty" validate:"required"`
	Birthday   string    `json:"birthday,omitempty" validate:"required"`
	Password   string    `json:"password,omitempty" validate:"required"`
	Created_At time.Time `json:"createdAt,omitempty" validate:"-"`
	Updated_At time.Time `json:"updatedAt,omitempty" validate:"-"`
	Is_Deleted bool      `json:"isDeleted,omitempty" validate:"-"`
}
