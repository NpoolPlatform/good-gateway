//nolint:dupl
package topmost

import (
	"context"

	topmost1 "github.com/NpoolPlatform/good-gateway/pkg/app/good/topmost"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	npool "github.com/NpoolPlatform/message/npool/good/gw/v1/app/good/topmost"
)

func (s *Server) UpdateTopMost(ctx context.Context, in *npool.UpdateTopMostRequest) (*npool.UpdateTopMostResponse, error) {
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithID(&in.ID, true),
		topmost1.WithEntID(&in.EntID, true),
		topmost1.WithAppID(&in.AppID, true),
		topmost1.WithTitle(in.Title, false),
		topmost1.WithMessage(in.Message, false),
		topmost1.WithPosters(in.Posters, false),
		topmost1.WithStartAt(in.StartAt, false),
		topmost1.WithEndAt(in.EndAt, false),
		topmost1.WithThresholdCredits(in.ThresholdCredits, false),
		topmost1.WithRegisterElapsedSeconds(in.RegisterElapsedSeconds, false),
		topmost1.WithThresholdPurchases(in.ThresholdPurchases, false),
		topmost1.WithThresholdPaymentAmount(in.ThresholdPaymentAmount, false),
		topmost1.WithKycMust(in.KycMust, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateTopMost(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateTopMostResponse{
		Info: info,
	}, nil
}

func (s *Server) UpdateNTopMost(ctx context.Context, in *npool.UpdateNTopMostRequest) (*npool.UpdateNTopMostResponse, error) {
	handler, err := topmost1.NewHandler(
		ctx,
		topmost1.WithID(&in.ID, true),
		topmost1.WithEntID(&in.EntID, true),
		topmost1.WithAppID(&in.TargetAppID, true),
		topmost1.WithTitle(in.Title, false),
		topmost1.WithMessage(in.Message, false),
		topmost1.WithPosters(in.Posters, false),
		topmost1.WithStartAt(in.StartAt, false),
		topmost1.WithEndAt(in.EndAt, false),
		topmost1.WithThresholdCredits(in.ThresholdCredits, false),
		topmost1.WithRegisterElapsedSeconds(in.RegisterElapsedSeconds, false),
		topmost1.WithThresholdPurchases(in.ThresholdPurchases, false),
		topmost1.WithThresholdPaymentAmount(in.ThresholdPaymentAmount, false),
		topmost1.WithKycMust(in.KycMust, false),
	)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	info, err := handler.UpdateTopMost(ctx)
	if err != nil {
		logger.Sugar().Errorw(
			"UpdateNTopMost",
			"In", in,
			"Error", err,
		)
		return &npool.UpdateNTopMostResponse{}, status.Error(codes.Aborted, err.Error())
	}

	return &npool.UpdateNTopMostResponse{
		Info: info,
	}, nil
}
