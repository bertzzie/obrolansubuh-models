package models

import (
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"time"
)

type Post struct {
	ID        int64
	Title     string    `sql:"size:1024;not null"`
	Content   string    `sql:"type:text;not null"`
	Published bool      `sql:"default:false;not null"`
	CreatedAt time.Time `sql:"default:NOW();not null"`
	UpdatedAt time.Time `sql:"default:NOW();not null"`

	Author   *Contributor
	AuthorID int64 `sql:"not null"` // foreign key
}

func (p Post) Validate(v *revel.Validation) {
	v.Required(p.Title)
	v.Required(p.Content)
	v.Required(p.Author)
}
