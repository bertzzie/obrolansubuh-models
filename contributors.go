package models

import (
	"github.com/revel/revel"
	"time"
)

type Contributor struct {
	Id        int64     `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	About     string    `db:"about"`
	JoinSince time.Time `db:"join_since"`
	Photo     string    `db:"photo"`
}

func (c Contributor) Validate(v *revel.Validation) {
	v.Required(c.Name)
	v.Required(c.Email)
	v.Required(c.Password)
}
