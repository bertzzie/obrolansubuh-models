package models

import (
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Contributor struct {
	ID        int64
	Name      string    `sql:"size:255;not null;index:name"`
	Email     string    `sql:"size:255;not null;unique_index"`
	Password  string    `sql:"type:text;not null"`
	About     string    `sql:"type:text;not null"`
	JoinSince time.Time `sql:"default:NOW();not null"`
	Photo     string    `sql:"type:text;not null"`

	Type   *ContributorType
	TypeID int64 `sql:"not null"` // foreign key

	Posts []Post
}

func (c Contributor) Validate(v *revel.Validation) {
	v.Required(c.Name)
	v.Required(c.Email)
	v.Required(c.Password)
}

func (c Contributor) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(password))

	return err == nil
}

func (c *Contributor) SetPassword(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), 15)

	if err != nil {
		return "", err
	}

	pass := string(result[:])

	c.Password = pass
	return pass, nil
}
