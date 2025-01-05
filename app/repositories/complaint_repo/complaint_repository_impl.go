package complaint_repo

import (
	"context"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/facades"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/models"

	"gorm.io/gorm"
)

type ComplaintRepositoryImpl struct {
	orm *gorm.DB
}

// FindAllByIDUser implements ComplaintRepository.
func (c *ComplaintRepositoryImpl) FindAll(ctx context.Context) (result []models.Complaint, err error) {
	err = c.orm.Find(&result).Error
	return
}

// Delete implements ComplaintRepository.
func (c *ComplaintRepositoryImpl) Delete(ctx context.Context, id uint) error {
	return c.orm.WithContext(ctx).Where("id = ?", id).Delete(&models.Complaint{}).Error
}

// FindByIDUser implements ComplaintRepository.
func (c *ComplaintRepositoryImpl) FindByID(ctx context.Context, id uint) (result models.Complaint, err error) {
	err = c.orm.Where("id = ?", id).First(&result).Error
	return
}

// Save implements ComplaintRepository.
func (c *ComplaintRepositoryImpl) Save(ctx context.Context, data *models.Complaint) error {
	return c.orm.WithContext(ctx).Create(data).Error
}

// Update implements ComplaintRepository.
func (c *ComplaintRepositoryImpl) Update(ctx context.Context, data *models.Complaint) error {
	return c.orm.WithContext(ctx).Where("id = ?", data.ID).Updates(data).Error
}

func NewComplaintRepository() ComplaintRepository {
	return &ComplaintRepositoryImpl{orm: facades.Orm()}
}
