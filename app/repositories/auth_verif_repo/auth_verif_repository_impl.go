package authverifrepo

import (
	"context"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/facades"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/models"

	"gorm.io/gorm"
)

type AuthVerifRepositoryImpl struct {
	orm *gorm.DB
}

func NewAuthVerifRepository() AuthVerifRepository {
	return &AuthVerifRepositoryImpl{orm: facades.Orm()}
}

// Delete implements AuthVerifRepository.
func (u *AuthVerifRepositoryImpl) Delete(ctx context.Context, id uint) error {
	err := u.orm.WithContext(ctx).Delete(&models.AuthVerif{}, id).Error
	return err
}

// Save implements AuthVerifRepository.
func (u *AuthVerifRepositoryImpl) Save(ctx context.Context, data *models.AuthVerif) error {
	panic("implement me")
}

// Update implements AuthVerifRepository.
func (u *AuthVerifRepositoryImpl) Update(ctx context.Context, data *models.AuthVerif) error {
	err := u.orm.WithContext(ctx).Model(&models.AuthVerif{}).Where("id = ?", data.ID).Updates(data).Error
	return err
}

// FindByCode implements AuthVerifRepository.
func (u *AuthVerifRepositoryImpl) FindByCode(ctx context.Context, code string) (models.AuthVerif, error) {
	var authVerif models.AuthVerif
	err := u.orm.WithContext(ctx).Model(&models.AuthVerif{}).Where("unique_code = ?", code).Where("deleted_at IS NULL").First(&authVerif).Error
	return authVerif, err
}
