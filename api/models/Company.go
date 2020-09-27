package models

import (
	"strings"
)

// Company structure
type Company struct {
	ID           int      `gorm:"primary_key;auto_increment" json:"id"`
	Name         string   `gorm:"size:255;not null;unique" json:"name"`
	Size         int      `gorm:"not null" json:"size"`
	Industry     string   `gorm:"size:255; not null" json:"industry"`
	Headquarters string   `gorm:"size:255;not null" json:"headquarters"`
	SocialMedia  []string `json:"social_media"`
	Type         int      `gorm:"not null" json:"type"`
}

// Prepare func removes all white space before saving
func (c *Company) Prepare() {
	c.ID = 0
	c.Name = strings.TrimSpace(c.Name)
	c.Industry = strings.TrimSpace(c.Industry)
	c.Headquarters = strings.TrimSpace(c.Headquarters)
	socialMedia := []string{}
	for _, mediaLink := range c.SocialMedia {
		socialMedia = append(socialMedia, strings.TrimSpace(mediaLink))
	}
	c.SocialMedia = socialMedia
}
