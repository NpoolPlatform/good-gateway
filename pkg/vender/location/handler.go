package location

import (
	"context"
	"fmt"

	constant "github.com/NpoolPlatform/good-gateway/pkg/const"

	"github.com/google/uuid"
)

type Handler struct {
	ID       *string
	Country  *string
	Province *string
	City     *string
	Address  *string
	BrandID  *string
	Offset   int32
	Limit    int32
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

func WithCountry(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid country")
			}
			return nil
		}
		if len(*s) < leastStrLen {
			return fmt.Errorf("invalid country")
		}
		h.Country = s
		return nil
	}
}

func WithProvince(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid province")
			}
			return nil
		}
		const leastStrLen = 3
		if len(*s) < leastStrLen {
			return fmt.Errorf("invalid province")
		}
		h.Province = s
		return nil
	}
}

func WithCity(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid city")
			}
			return nil
		}
		const leastStrLen = 3
		if len(*s) < leastStrLen {
			return fmt.Errorf("invalid city")
		}
		h.City = s
		return nil
	}
}

func WithAddress(s *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if s == nil {
			if must {
				return fmt.Errorf("invalid address")
			}
			return nil
		}
		const leastStrLen = 3
		if len(*s) < leastStrLen {
			return fmt.Errorf("invalid address")
		}
		h.Address = s
		return nil
	}
}

func WithBrandID(id *string, must bool) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if id == nil {
			if must {
				return fmt.Errorf("invalid brandid")
			}
			return nil
		}
		if _, err := uuid.Parse(*id); err != nil {
			return err
		}
		h.BrandID = id
		return nil
	}
}

func WithOffset(n int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		h.Offset = n
		return nil
	}
}

func WithLimit(n int32) func(context.Context, *Handler) error {
	return func(ctx context.Context, h *Handler) error {
		if n == 0 {
			n = constant.DefaultRowLimit
		}
		h.Limit = n
		return nil
	}
}
