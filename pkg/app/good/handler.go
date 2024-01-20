//nolint:dupl
package good

import (
	"context"
	"fmt"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	constant "github.com/NpoolPlatform/good-gateway/pkg/const"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID                     *uint32
	EntID                  *string
	AppID                  *string
	GoodID                 *string
	Online                 *bool
	Visible                *bool
	GoodName               *string
	UnitPrice              *string
	PackagePrice           *string
	DisplayIndex           *int32
	PurchaseLimit          *int32
	SaleStartAt            *uint32
	SaleEndAt              *uint32
	ServiceStartAt         *uint32
	TechniqueFeeRatio      *string
	ElectricityFeeRatio    *string
	Descriptions           []string
	GoodBanner             *string
	DisplayNames           []string
	EnablePurchase         *bool
	EnableProductPage      *bool
	CancelMode             *types.CancelMode
	UserPurchaseLimit      *string
	DisplayColors          []string
	CancellableBeforeStart *uint32
	ProductPage            *string
	EnableSetCommission    *bool
	Posters                []string
	MinOrderAmount         *string
	MaxOrderAmount         *string
	MaxUserAmount          *string
	MinOrderDuration       *uint32
	MaxOrderDuration       *uint32
	PackageWithRequireds   *bool
	Offset                 int32
	Limit                  int32
}

const leastStrLen = 3

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		h.ID = id
		return nil
	}
}

func WithEntID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid id")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.EntID = id
		return nil
	}
}

func WithAppID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid appid")
			}
			return nil
		}
		exist, err := appmwcli.ExistApp(ctx, *id)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid app")
		}
		h.AppID = id
		return nil
	}
}

func WithGoodID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid goodid")
			}
			return nil
		}
		h.GoodID = id
		return nil
	}
}

func WithOnline(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Online = b
		return nil
	}
}

func WithVisible(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Visible = b
		return nil
	}
}

func WithGoodName(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid goodname")
			}
			return nil
		}
		if len(*s) < leastStrLen {
			return fmt.Errorf("invalid goodname")
		}
		h.GoodName = s
		return nil
	}
}

func WithUnitPrice(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid unitprice")
			}
			return nil
		}
		price, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		if price.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid unitprice")
		}
		h.UnitPrice = s
		return nil
	}
}

func WithPackagePrice(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid packageprice")
			}
			return nil
		}
		price, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		if price.Cmp(decimal.NewFromInt(0)) <= 0 {
			return fmt.Errorf("invalid packageprice")
		}
		h.PackagePrice = s
		return nil
	}
}

func WithDisplayIndex(n *int32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.DisplayIndex = n
		return nil
	}
}

func WithPurchaseLimit(n *int32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.PurchaseLimit = n
		return nil
	}
}

func WithSaleStartAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SaleStartAt = n
		return nil
	}
}

func WithSaleEndAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.SaleEndAt = n
		return nil
	}
}

func WithServiceStartAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ServiceStartAt = n
		return nil
	}
}

func WithTechniqueFeeRatio(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid techniquefeeratio")
			}
			return nil
		}
		if _, err := decimal.NewFromString(*s); err != nil {
			return err
		}
		h.TechniqueFeeRatio = s
		return nil
	}
}

func WithElectricityFeeRatio(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid electricityfeeratio")
			}
			return nil
		}
		if _, err := decimal.NewFromString(*s); err != nil {
			return err
		}
		h.ElectricityFeeRatio = s
		return nil
	}
}

func WithDescriptions(ss []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Descriptions = ss
		return nil
	}
}

func WithGoodBanner(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid goodbanner")
			}
			return nil
		}
		if len(*s) < leastStrLen {
			return fmt.Errorf("invalid goodbanner")
		}
		h.GoodBanner = s
		return nil
	}
}

func WithDisplayNames(ss []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.DisplayNames = ss
		return nil
	}
}

func WithEnablePurchase(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.EnablePurchase = b
		return nil
	}
}

func WithEnableProductPage(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.EnableProductPage = b
		return nil
	}
}

func WithCancelMode(e *types.CancelMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid cancelmode")
			}
			return nil
		}
		switch *e {
		case types.CancelMode_CancellableBeforeStart:
		case types.CancelMode_CancellableBeforeBenefit:
		case types.CancelMode_Uncancellable:
		default:
			return fmt.Errorf("invalid cancelmode")
		}
		h.CancelMode = e
		return nil
	}
}

func WithUserPurchaseLimit(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid userpurchaseamount")
			}
			return nil
		}
		if _, err := decimal.NewFromString(*s); err != nil {
			return err
		}
		h.UserPurchaseLimit = s
		return nil
	}
}

func WithDisplayColors(ss []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.DisplayColors = ss
		return nil
	}
}

func WithPosters(ss []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Posters = ss
		return nil
	}
}

func WithCancellableBeforeStart(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.CancellableBeforeStart = n
		return nil
	}
}

func WithProductPage(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ProductPage = s
		return nil
	}
}

func WithEnableSetCommission(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.EnableSetCommission = b
		return nil
	}
}

func WithMinOrderAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid minorderamount")
			}
			return nil
		}
		_, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.MinOrderAmount = s
		return nil
	}
}

func WithMaxOrderAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid maxorderamount")
			}
			return nil
		}
		_, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.MaxOrderAmount = s
		return nil
	}
}

func WithMaxUserAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid price")
			}
			return nil
		}
		_, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.MaxUserAmount = s
		return nil
	}
}

func WithMinOrderDuration(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.MinOrderDuration = n
		return nil
	}
}

func WithMaxOrderDuration(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.MaxOrderDuration = n
		return nil
	}
}

func WithPackageWithRequireds(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.PackageWithRequireds = b
		return nil
	}
}

func WithOffset(offset int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = offset
		return nil
	}
}

func WithLimit(limit int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if limit == 0 {
			limit = constant.DefaultRowLimit
		}
		h.Limit = limit
		return nil
	}
}
