package user

import (
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, u *User) error
	FindById(ctx context.Context, id uuid.UUID) (*User, error)
	List(ctx context.Context, limit, offset int) ([]User, int64, error)
	Update(ctx context.Context, u *User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type gormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) Repository {
	return &gormRepository{db: db}
}

func (r *gormRepository) Create(ctx context.Context, u *User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

func (r *gormRepository) FindById(ctx context.Context, id uuid.UUID) (*User, error) {
	var u User
	if err := r.db.WithContext(ctx).First(&u, "id=?", id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *gormRepository) List(ctx context.Context, limit, offset int) ([]User, int64, error) {
	var users []User
	var total int64
	tx := r.db.WithContext(ctx).Model(&User{})
	if err := tx.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := tx.Limit(limit).Offset(offset).Order("created_at desc").Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil

}

func (r *gormRepository) Update(ctx context.Context, u *User) error {
	return r.db.WithContext(ctx).Save(u).Error
}

func (r *gormRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&User{}, "id=?", id).Error
}
