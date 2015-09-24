package models

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/revel/revel"
	"regexp"
	"strings"
	"time"
)

type Post struct {
	ID         int64
	Title      string    `sql:"size:1024;not null"`
	Content    string    `sql:"type:text;not null"`
	Published  bool      `sql:"default:false;not null"`
	CoverImage string    `sql:"type:text;not null"`
	CreatedAt  time.Time `sql:"default:NOW();not null"`
	UpdatedAt  time.Time `sql:"default:NOW();not null"`

	Author     *Contributor
	AuthorID   int64       `sql:"not null;REFERENCES contributors(id)"` // foreign key
	Categories []*Category `gorm:"many2many:post_categories"`           // many to many to prepare for future use
}

func (p *Post) TitleSlug() string {
	reg, err := regexp.Compile("[^A-Za-z0-9- ]+")
	if err != nil {
		revel.ERROR.Println("[LGERR] Failed to compile regex in Models.Post.TitleSlug")
		return ""
	}

	safe := reg.ReplaceAllString(p.Title, "")
	safe = strings.Replace(safe, " ", "-", -1)
	safe = strings.ToLower(strings.Trim(safe, " "))

	return safe
}

func (p *Post) FirstParagraph() string {
	doc, rerr := goquery.NewDocumentFromReader(strings.NewReader(p.Content))
	if rerr != nil {
		revel.ERROR.Println("[LGERR] Failed to read HTML from Post.Content")
		return ""
	}

	return doc.Find("p").First().Text()
}

func (p Post) Validate(v *revel.Validation) {
	v.Required(p.Title)
	v.Required(p.Content)
	v.Required(p.Author)
}
