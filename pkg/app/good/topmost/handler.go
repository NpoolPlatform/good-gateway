package topmost

import (
	"context"
	"fmt"

	appmwcli "github.com/NpoolPlatform/appuser-middleware/pkg/client/app"
	constant "github.com/NpoolPlatform/good-gateway/pkg/const"
	types "github.com/NpoolPlatform/message/npool/basetypes/good/v1"

	"github.com/google/uuid"
)

type Handler struct {
	ID                     *string
	AppID                  *string
	TopMostType            *types.GoodTopMostType
	Title                  *string
	Message                *string
	Posters                []string
	StartAt                *uint32
	EndAt                  *uint32
	ThresholdCredits       *string
	RegisterElapsedSeconds *uint32
	ThresholdPurchases     *uint32
	ThresholdPaymentAmount *string
	KycMust                *bool
	Offset                 int32
	Limit                  int32
}

func NewHandler(ctx context.Context, options ...func(context.Context, *Handler) error) (*Handler, error) {
	handler := &Handler{}
	for _, opt := range options {
		if err := opt(ctx, handler); err != nil {
			return nil, err
		}
	}
	return handler, nil
}

func WithID(id *string, must bool) func(context.Context, *Handler) error {
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
		h.ID = id
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

func WithTopMostType(e *types.GoodTopMostType, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if e == nil {
			if must {
				return fmt.Errorf("invalid topmosttype")
			}
			return nil
		}
		switch *e {
		case types.GoodTopMostType_TopMostPromotion:
		case types.GoodTopMostType_TopMostNoviceExclusive:
		case types.GoodTopMostType_TopMostBestOffer:
		case types.GoodTopMostType_TopMostInnovationStarter:
		case types.GoodTopMostType_TopMostLoyaltyExclusive:
		default:
			return fmt.Errorf("invalid topmosttype")
		}
		h.TopMostType = e
		return nil
	}
}

func WithTitle(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Title = s
		return nil
	}
}

func WithMessage(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Message = s
		return nil
	}
}

func WithPosters(ss []string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Posters = ss
		return nil
	}
}

func WithStartAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.StartAt = n
		return nil
	}
}

func WithEndAt(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.EndAt = n
		return nil
	}
}

func WithThresholdCredits(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid thresholdcredits")
			}
			return nil
		}
		h.ThresholdCredits = s
		return nil
	}
}

func WithRegisterElapsedSeconds(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.RegisterElapsedSeconds = n
		return nil
	}
}

func WithThresholdPurchases(n *uint32, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.ThresholdPurchases = n
		return nil
	}
}

func WithThresholdPaymentAmount(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid thresholdpaymentamount")
			}
			return nil
		}
		h.ThresholdPaymentAmount = s
		return nil
	}
}

func WithKycMust(b *bool, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.KycMust = b
		return nil
	}
}

func WithOffset(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = value
		return nil
	}
}

func WithLimit(value int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if value == 0 {
			value = constant.DefaultRowLimit
		}
		h.Limit = value
		return nil
	}
}
