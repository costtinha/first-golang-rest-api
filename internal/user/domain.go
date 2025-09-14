package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"type:varchar(120);not null"`
	Email     string    `gorm:"type:varchar(180);uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) BeforeCreated(_ any) error {
	if u.ID == uuid.Nil {
		u.ID = uuid.New()
	}
	return nil
}

type CreateUserInput struct {
	Name string `json:"name" validate:"required,min=2,max=120"`
	Emal string `json:"email" validate:"required,email,max=180"`
}

type UpdateUserInput struct {
	Name  *string `json:"name" validate:"omitempty,min=2,max=120"`
	Email *string `json:"email" validate:"omitempty,email,max=180"`
}
