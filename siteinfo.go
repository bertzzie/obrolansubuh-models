package models

import (
	"github.com/revel/revel"
)

type SiteInfo struct {
	ID             int64
	AboutUsTitle   string `sql:"size:1024;not null"`
	AboutUsContent string `sql:"type:text;not null"`
	TwitterURL     string `sql:"type:text;not null"`
	FacebookURL    string `sql:"type:text;not null"`
}

func (si SiteInfo) Validate(v *revel.Validation) {
	v.MaxSize(si.AboutUsTitle, 1024)

	v.Required(si.AboutUsTitle)
	v.Required(si.AboutUsContent)
	v.Required(si.TwitterURL)
	v.Required(si.FacebookURL)
}
