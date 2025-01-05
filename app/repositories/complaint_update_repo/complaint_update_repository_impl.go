package complaint_update_repo

import (
	"context"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/facades"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/models"

	"gorm.io/gorm"
)

type ComplaintUpdateRepositoryImpl struct {
	orm *gorm.DB
}

// FindAllByIDUser implements ComplaintUpdateRepository.
func (c *ComplaintUpdateRepositoryImpl) FindAll(ctx context.Context) (result []models.ComplaintUpdate, err error) {
	err = c.orm.Find(&result).Error
	return
}

// Delete implements ComplaintUpdateRepository.
func (c *ComplaintUpdateRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return c.orm.WithContext(ctx).Where("id = ?", id).Delete(&models.ComplaintUpdate{}).Error
}

// FindByIDUser implements ComplaintUpdateRepository.
func (c *ComplaintUpdateRepositoryImpl) FindByID(ctx context.Context, id uint) (result models.ComplaintUpdate, err error) {
	err = c.orm.Where("id = ?", id).First(&result).Error
	return
}

// Save implements ComplaintUpdateRepository.
func (c *ComplaintUpdateRepositoryImpl) Save(ctx context.Context, data *models.ComplaintUpdate) error {
	return c.orm.WithContext(ctx).Create(data).Error
}

// Update implements ComplaintUpdateRepository.
func (c *ComplaintUpdateRepositoryImpl) Update(ctx context.Context, data *models.ComplaintUpdate) error {
	return c.orm.WithContext(ctx).Where("id = ?", data.ID).Updates(data).Error
}

func NewComplaintUpdateRepository() ComplaintUpdateRepository {
	return &ComplaintUpdateRepositoryImpl{orm: facades.Orm()}
}
