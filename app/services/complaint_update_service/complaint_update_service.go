package complaint_update_service

import (
	"context"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/dto"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/http/requests"
)

type ComplaintUpdateService interface {
	Create(ctx context.Context, req requests.ComplaintUpdateCreateRequest) error
	Delete(ctx context.Context, id uint) error
	FindAll(ctx context.Context) ([]dto.ComplaintUpdateData, error)
	FindByID(ctx context.Context, id uint) (*dto.ComplaintUpdateData, error)
	Show(ctx context.Context, id uint) (*dto.ComplaintUpdateData, error)
	Update(ctx context.Context, id uint, req requests.ComplaintUpdateUpdateRequest) error
}
