package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgtype"

	db "github.com/envde/effmobi/interanl/pkg/postgres"
	"github.com/envde/effmobi/interanl/transport/dto"
)

type SubscriptionService struct {
	q   *db.Queries
	log *slog.Logger
}

func NewSubscriptionService(q *db.Queries, log *slog.Logger) *SubscriptionService {
	return &SubscriptionService{q: q, log: log}
}

// parseMonthYear парсит строку вида "07-2025" в time.Time (первый день месяца)
func parseMonthYear(s string) (time.Time, error) {
	t, err := time.Parse("01-2006", s)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date format, expected MM-YYYY: %w", err)
	}
	return t, nil
}

func toPgtypeDate(t time.Time) pgtype.Date {
	return pgtype.Date{
		Time:  t,
		Valid: true,
	}
}

func (s *SubscriptionService) Create(ctx context.Context, r dto.CreateSubscriptionRequest) (db.Subscription, error) {
	s.log.Debug("Create subscription", "service_name", r.ServiceName, "user_id", r.UserID)

	startDate, err := parseMonthYear(r.StartDate)
	if err != nil {
		return db.Subscription{}, err
	}

	var userID pgtype.UUID
	if err := userID.Scan(r.UserID); err != nil {
		return db.Subscription{}, fmt.Errorf("invalid user_id: %w", err)
	}

	params := db.CreateSubscriptionParams{
		ServiceName: r.ServiceName,
		Price:       r.Price,
		UserID:      userID,
		StartDate:   toPgtypeDate(startDate),
	}

	if r.EndDate != nil {
		endDate, err := parseMonthYear(*r.EndDate)
		if err != nil {
			return db.Subscription{}, err
		}
		params.EndDate = toPgtypeDate(endDate)
	}

	return s.q.CreateSubscription(ctx, params)
}

func (s *SubscriptionService) Get(ctx context.Context, id int64) (db.Subscription, error) {
	return s.q.GetSubscription(ctx, id)
}

func (s *SubscriptionService) List(ctx context.Context) ([]db.Subscription, error) {
	return s.q.ListSubscriptions(ctx)
}

func (s *SubscriptionService) Update(ctx context.Context, id int64, r dto.UpdateSubscriptionRequest) (db.Subscription, error) {
	startDate, err := parseMonthYear(r.StartDate)
	if err != nil {
		return db.Subscription{}, err
	}

	params := db.UpdateSubscriptionParams{
		ID:          id,
		ServiceName: r.ServiceName,
		Price:       r.Price,
		StartDate:   toPgtypeDate(startDate),
	}

	if r.EndDate != nil {
		endDate, err := parseMonthYear(*r.EndDate)
		if err != nil {
			return db.Subscription{}, err
		}
		params.EndDate = toPgtypeDate(endDate)
	}

	return s.q.UpdateSubscription(ctx, params)
}

func (s *SubscriptionService) Delete(ctx context.Context, id int64) error {
	return s.q.DeleteSubscription(ctx, id)
}

func (s *SubscriptionService) Sum(ctx context.Context, r dto.SumRequest) (int32, error) {
	from, err := parseMonthYear(r.From)
	if err != nil {
		return 0, fmt.Errorf("invalid 'from': %w", err)
	}
	to, err := parseMonthYear(r.To)
	if err != nil {
		return 0, fmt.Errorf("invalid 'to': %w", err)
	}

	var userID pgtype.UUID
	if r.UserID != nil {
		if err := userID.Scan(*r.UserID); err != nil {
			return 0, fmt.Errorf("invalid user_id: %w", err)
		}
	}

	total, err := s.q.SumSubscriptionCost(ctx, db.SumSubscriptionCostParams{
		UserID:      userID,
		ServiceName: r.ServiceName,
		FromDate:    toPgtypeDate(from),
		ToDate:      toPgtypeDate(to),
	})
	if err != nil {
		return 0, err
	}
	return total, nil
}
