package repo

import (
	"context"
	"finnomena/internal/controller/job/model"
	"time"
)

type HTTPFundService interface {
	GetFundRanking(ctx context.Context, startDate time.Time, endDate time.Time) ([]model.Fund, error)
}
