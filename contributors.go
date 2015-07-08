package models

import (
	"github.com/revel/revel"
	"time"
)

type Contributor struct {
	ID        int64
	Name      string    `sql:"size:255;not null"`
	Email     string    `sql:"size:255;not null"`
	Password  string    `sql:"type:text;not null"`
	About     string    `sql:"type:text;not null"`
	JoinSince time.Time `sql:"default:NOW();not null"`
	Photo     string    `sql:"type:text;not null"`

	Posts []Post
}

func (c Contributor) Validate(v *revel.Validation) {
	v.Required(c.Name)
	v.Required(c.Email)
	v.Required(c.Password)
}
