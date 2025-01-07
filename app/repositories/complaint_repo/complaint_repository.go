package complaint_repo

import (
	"context"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/models"
)

type ComplaintRepository interface {
	Save(ctx context.Context, data *models.Complaint) error
	FindByID(ctx context.Context, id uint) (models.Complaint, error)
	FindAll(ctx context.Context) ([]models.Complaint, error)
	FindAllByUserID(ctx context.Context, id uint) ([]models.Complaint, error)
	Update(ctx context.Context, data *models.Complaint) error
	Delete(ctx context.Context, id uint) error
}
	