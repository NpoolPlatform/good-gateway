package good

import (
	"context"
	"fmt"

	coinmwcli "github.com/NpoolPlatform/chain-middleware/pkg/client/coin"
	constant "github.com/NpoolPlatform/good-gateway/pkg/const"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Handler struct {
	ID                    *uint32
	EntID                 *string
	DeviceInfoID          *string
	DurationDays          *int32
	CoinTypeID            *string
	VendorLocationID      *string
	UnitPrice             *string
	BenefitType           *types.BenefitType
	GoodType              *types.GoodType
	Title                 *string
	QuantityUnit          *string
	QuantityUnitAmount    *string
	DeliveryAt            *uint32
	StartAt               *uint32
	StartMode             *types.GoodStartMode
	TestOnly              *bool
	Total                 *string
	Posters               []string
	Labels                []types.GoodLabel
	BenefitIntervalHours  *uint32
	UnitLockDeposit       *string
	UnitType              *types.GoodUnitType
	QuantityCalculateType *types.GoodUnitCalculateType
	DurationType          *types.GoodDurationType
	DurationCalculateType *types.GoodUnitCalculateType
	SettlementType        *types.GoodSettlementType
	Offset                int32
	Limit                 int32
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

func WithDeviceInfoID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid deviceinfoid")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.DeviceInfoID = id
		return nil
	}
}

func WithDurationDays(n *int32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid durationdays")
			}
			return nil
		}
		h.DurationDays = n
		return nil
	}
}

func WithCoinTypeID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid cointypeid")
			}
			return nil
		}
		exist, err := coinmwcli.ExistCoin(ctx, *id)
		if err != nil {
			return err
		}
		if !exist {
			return fmt.Errorf("invalid coin")
		}
		h.CoinTypeID = id
		return nil
	}
}

func WithVendorLocationID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid vendorlocationid")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.VendorLocationID = id
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
		if _, err := decimal.NewFromString(*s); err != nil {
			return err
		}
		h.UnitPrice = s
		return nil
	}
}

func WithBenefitType(e *types.BenefitType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid benefittype")
			}
			return nil
		}
		switch *e {
		case types.BenefitType_BenefitTypePlatform:
		case types.BenefitType_BenefitTypePool:
		default:
			return fmt.Errorf("invalid benefittype")
		}
		h.BenefitType = e
		return nil
	}
}

func WithGoodType(e *types.GoodType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid goodtype")
			}
			return nil
		}
		switch *e {
		case types.GoodType_PowerRenting:
		case types.GoodType_MachineRenting:
			fallthrough //nolint
		case types.GoodType_MachineHosting:
			fallthrough //nolint
		case types.GoodType_TechniqueServiceFee:
			fallthrough //nolint
		case types.GoodType_ElectricityFee:
			return fmt.Errorf("not implemented")
		default:
			return fmt.Errorf("invalid goodtype")
		}
		h.GoodType = e
		return nil
	}
}

func WithTitle(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid title")
			}
			return nil
		}
		if len(*s) < leastStrLen {
			return fmt.Errorf("invalid title")
		}
		h.Title = s
		return nil
	}
}

func WithQuantityUnit(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid quantityunit")
			}
			return nil
		}
		const leastUnitLen = 2
		if len(*s) < leastUnitLen {
			return fmt.Errorf("invalid unit")
		}
		h.QuantityUnit = s
		return nil
	}
}

func WithQuantityUnitAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid quantityunitamount")
			}
			return nil
		}
		_, err := decimal.NewFromString(*s)
		if err != nil {
			return err
		}
		h.QuantityUnitAmount = s
		return nil
	}
}

func WithDeliveryAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid deliveryat")
			}
			return nil
		}
		h.DeliveryAt = n
		return nil
	}
}

func WithStartAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid startat")
			}
			return nil
		}
		h.StartAt = n
		return nil
	}
}

func WithStartMode(e *types.GoodStartMode, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid goodstartmode")
			}
			return nil
		}
		switch *e {
		case types.GoodStartMode_GoodStartModeTBD:
		case types.GoodStartMode_GoodStartModeConfirmed:
		case types.GoodStartMode_GoodStartModeNextDay:
		case types.GoodStartMode_GoodStartModeInstantly:
		case types.GoodStartMode_GoodStartModePreset:
		default:
			return fmt.Errorf("invalid goodstartmode")
		}
		h.StartMode = e
		return nil
	}
}

func WithTestOnly(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.TestOnly = b
		return nil
	}
}

func WithTotal(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid total")
			}
			return nil
		}
		if _, err := decimal.NewFromString(*s); err != nil {
			return err
		}
		h.Total = s
		return nil
	}
}

func WithPosters(ss []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, s := range ss {
			if len(s) < leastStrLen {
				return fmt.Errorf("invalid poster")
			}
		}
		h.Posters = ss
		return nil
	}
}

func WithLabels(es []types.GoodLabel, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		for _, e := range es {
			switch e {
			case types.GoodLabel_GoodLabelPromotion:
			case types.GoodLabel_GoodLabelNoviceExclusive:
			case types.GoodLabel_GoodLabelInnovationStarter:
			case types.GoodLabel_GoodLabelLoyaltyExclusive:
			default:
				return fmt.Errorf("invalid label")
			}
		}
		h.Labels = es
		return nil
	}
}

func WithBenefitIntervalHours(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == nil {
			if must {
				return fmt.Errorf("invalid benefitintervalhours")
			}
			return nil
		}
		h.BenefitIntervalHours = n
		return nil
	}
}

func WithUnitLockDeposit(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid unitlockdeposit")
			}
			return nil
		}
		if _, err := decimal.NewFromString(*s); err != nil {
			return err
		}
		h.UnitLockDeposit = s
		return nil
	}
}

func WithUnitType(e *types.GoodUnitType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid unittype")
			}
			return nil
		}
		switch *e {
		case types.GoodUnitType_GoodUnitByDuration:
		case types.GoodUnitType_GoodUnitByQuantity:
		case types.GoodUnitType_GoodUnitByDurationAndQuantity:
		default:
			return fmt.Errorf("invalid unittype")
		}
		h.UnitType = e
		return nil
	}
}

func WithQuantityCalculateType(e *types.GoodUnitCalculateType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid quantitycalculatetype")
			}
			return nil
		}
		switch *e {
		case types.GoodUnitCalculateType_GoodUnitCalculateBySelf:
		case types.GoodUnitCalculateType_GoodUnitCalculateByParent:
		default:
			return fmt.Errorf("invalid quantitycalculatetype")
		}
		h.QuantityCalculateType = e
		return nil
	}
}

func WithDurationType(e *types.GoodDurationType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid durationtype")
			}
			return nil
		}
		switch *e {
		case types.GoodDurationType_GoodDurationByHour:
		case types.GoodDurationType_GoodDurationByDay:
		case types.GoodDurationType_GoodDurationByMonth:
		case types.GoodDurationType_GoodDurationByYear:
		default:
			return fmt.Errorf("invalid durationtype")
		}
		h.DurationType = e
		return nil
	}
}

func WithSettlementType(e *types.GoodSettlementType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid settlementype")
			}
			return nil
		}
		switch *e {
		case types.GoodSettlementType_GoodSettledByCash:
		case types.GoodSettlementType_GoodSettledByProfit:
		default:
			return fmt.Errorf("invalid settlementtype")
		}
		h.SettlementType = e
		return nil
	}
}

func WithDurationCalculateType(e *types.GoodUnitCalculateType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid durationcalculatetype")
			}
			return nil
		}
		switch *e {
		case types.GoodUnitCalculateType_GoodUnitCalculateBySelf:
		case types.GoodUnitCalculateType_GoodUnitCalculateByParent:
		default:
			return fmt.Errorf("invalid durationcalculatetype")
		}
		h.DurationCalculateType = e
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
