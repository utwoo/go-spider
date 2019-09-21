package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// TagProcess presents the spider process with tag
type TagProcess struct {
	gorm.Model
	TagName    string `gorm:"type:varchar(64);not null;column:tag_name"`
	TagURL     string `gorm:"type:varchar(128);not null;column:tag_url"`
	StartIndex int    `gorm:"type:integer;not null;column:start_index"`
	Finished   string `gorm:"type:varchar(1);not null;column:finished"`
}

// TagProcessModel presents the process model and serializer
type TagProcessModel struct {
	ID         uint       `json:"id"`
	CreatedAt  time.Time  `json:"create_at"`
	UpdateAt   time.Time  `json:"update_at"`
	DeleteAt   *time.Time `json:"delete_at"`
	TagName    string     `json:"tag_name"`
	TagURL     string     `json:"tag_url"`
	StartIndex int        `json:"start_index"`
	Finished   string     `json:"finished"`
}

// TableName indicate the table name
func (m *TagProcess) TableName() string {
	return "process"
}

// Serializer transfers entity to model
func (m *TagProcess) Serializer() TagProcessModel {
	return TagProcessModel{
		ID:         m.ID,
		CreatedAt:  m.CreatedAt,
		UpdateAt:   m.UpdatedAt,
		DeleteAt:   m.DeletedAt,
		TagURL:     m.TagURL,
		TagName:    m.TagName,
		StartIndex: m.StartIndex,
		Finished:   m.Finished,
	}
}
