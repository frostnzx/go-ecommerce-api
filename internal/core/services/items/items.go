package items

import (
	"context"

	"github.com/frostnzx/go-ecommerce-api/internal/ports"
)

type API interface {
	AddItem(ctx context.Context , AddItemReq) (*AddItemResp , error)
	GetItem(ctx context.Context , GetItemReq) (*GetItemResp , error)
	DeleteItem(ctx context.Context , DeleteItemReq) (*DeleteItemResp , error)
	ListItems(ctx context.Context , ListItemReq) (*ListItemResp , error)
}

type Service struct {
	itemsRepo ports.ItemsRepo
}

func NewService(ir ports.ItemsRepo) *Service {
	return &Service{
		itemsRepo: ir,
	}
}
