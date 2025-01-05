package complaint_service

import (
	"context"

	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/dto"
	"github.com/harrymuliawan03/layanan-pengaduan-api.git/app/http/requests"
)

type ComplaintService interface {
	Create(ctx context.Context, req requests.ComplaintCreateRequest) error
	Delete(ctx context.Context, id uint) error
	FindAll(ctx context.Context) ([]dto.ComplaintData, error)
	FindByID(ctx context.Context, id uint) (*dto.ComplaintData, error)
	Show(ctx context.Context, id uint) (*dto.ComplaintData, error)
	Update(ctx context.Context, id uint, req requests.ComplaintUpdateRequest) error
}
