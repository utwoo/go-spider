package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Process presents the spider process with tag
type Process struct {
	gorm.Model
	TagName    string `gorm:"type:varchar(64);not null;column:tag_name"`
	StartIndex int    `gorm:"type:integer;not null;column:start_index"`
	Finished   string `gorm:"type:varchar(1);not null;column:finished"`
}

// ProcessModel presents the process model and serializer
type ProcessModel struct {
	ID         uint       `json:"id"`
	CreatedAt  time.Time  `json:"create_at"`
	UpdateAt   time.Time  `json:"update_at"`
	DeleteAt   *time.Time `json:"delete_at"`
	TagName    string     `json:"tag_name"`
	StartIndex int        `json:"start_index"`
	Finished   string     `json:"finished"`
}

// Serializer transfers entity to model
func (m *Process) Serializer() ProcessModel {
	return ProcessModel{
		ID:         m.ID,
		CreatedAt:  m.CreatedAt,
		UpdateAt:   m.UpdatedAt,
		DeleteAt:   m.DeletedAt,
		TagName:    m.TagName,
		StartIndex: m.StartIndex,
		Finished:   m.Finished,
	}
}
