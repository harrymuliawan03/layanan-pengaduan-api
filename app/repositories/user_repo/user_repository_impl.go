package userrepo

import (
	"context"
	"database/sql"
	"time"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/facades"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/models"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	orm *gorm.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{orm: facades.Orm()}
}

// Delete implements UserRepository.
func (u *UserRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return u.orm.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Where("deleted_at IS NULL").Update("deleted_at", sql.NullTime{Valid: true, Time: time.Now()}).Error
}

// FindByEmail implements UserRepository.
func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := u.orm.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).Where("deleted_at IS NULL").First(&user).Error
	return user, err
}

// FindByID implements UserRepository.
func (u *UserRepositoryImpl) FindByID(ctx context.Context, id uint) (models.User, error) {
	var user models.User
	err := u.orm.WithContext(ctx).Model(&models.User{}).Where("id = ?", id).Where("deleted_at IS NULL").First(&user).Error
	return user, err
}

// Save implements UserRepository.
func (u *UserRepositoryImpl) Save(ctx context.Context, user *models.User) error {
	res := u.orm.WithContext(ctx).Create(&user)
	return res.Error
}

// Update implements UserRepository.
func (u *UserRepositoryImpl) Update(ctx context.Context, user models.User) error {
	err := u.orm.WithContext(ctx).Model(&models.User{}).Where("id = ?", user.ID).Updates(user).Error
	return err
}
