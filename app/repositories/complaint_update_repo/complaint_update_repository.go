package complaint_update_repo

import (
	"context"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/models"
)

type ComplaintUpdateRepository interface {
	Save(ctx context.Context, data *models.ComplaintUpdate) error
	FindByID(ctx context.Context, id uint) (models.ComplaintUpdate, error)
	FindAll(ctx context.Context) ([]models.ComplaintUpdate, error)
	Update(ctx context.Context, data *models.ComplaintUpdate) error
	Delete(ctx context.Context, id uint) error
}
