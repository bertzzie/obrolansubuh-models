package models

import (
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

	Author     *Contributor
	AuthorID   int64      `sql:"not null;REFERENCES contributors(id)"` // foreign key
	Categories []Category `gorm:"many2many:post_categories"`           // many to many to prepare for future use
}

func (p Post) Validate(v *revel.Validation) {
	v.Required(p.Title)
	v.Required(p.Content)
	v.Required(p.Author)
}
