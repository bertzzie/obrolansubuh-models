package models

import (
	"github.com/revel/revel"
	"time"
)

type Post struct {
	Id        int64        `db:"id"`
	Title     string       `db:"title"`
	Content   string       `db:"content"`
	Author    *Contributor `db:"author"`
	Published bool         `db:"published"`
	Created   time.Time    `db:"created"`
	Updated   time.Time    `db:"updated"`
}

func (p Post) Validate(v *revel.Validation) {
	v.Required(p.Title)
	v.Required(p.Content)
	v.Required(p.Author)
}
