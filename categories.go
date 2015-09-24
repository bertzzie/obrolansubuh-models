package models

import (
	"github.com/revel/revel"
	"regexp"
	"strings"
)

type Category struct {
	ID          int64
	Heading     string `sql:"type:text;not null"`
	Description string `sql:"type:text;"`
	Image       string `sql:"type:text;"`

	Posts []*Post `gorm:"many2many:post_categories"` // many to many to prepare for future use
}

func (c Category) Validate(v *revel.Validation) {
	v.Required(c.Heading)
}

func (c *Category) Slug() string {
	reg, err := regexp.Compile("[^A-Za-z0-9- ]+")
	if err != nil {
		revel.ERROR.Println("[LGERR] Failed to compile regex in Models.Category.TitleSlug")
		return ""
	}

	safe := reg.ReplaceAllString(c.Heading, "")
	safe = strings.Replace(safe, " ", "-", -1)
	safe = strings.Trim(safe, " ")

	return safe
}
