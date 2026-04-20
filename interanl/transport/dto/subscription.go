package dto

// CreateSubscriptionRequest — тело запроса на создание подписки
type CreateSubscriptionRequest struct {
	ServiceName string  `json:"service_name" example:"Yandex Plus"`
	Price       int32   `json:"price"        example:"400"`
	UserID      string  `json:"user_id"      example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string  `json:"start_date"   example:"07-2025"`
	EndDate     *string `json:"end_date"     example:"12-2025"`
}

// UpdateSubscriptionRequest — тело запроса на обновление подписки
type UpdateSubscriptionRequest struct {
	ServiceName string  `json:"service_name" example:"Yandex Plus"`
	Price       int32   `json:"price"        example:"400"`
	StartDate   string  `json:"start_date"   example:"07-2025"`
	EndDate     *string `json:"end_date"     example:"12-2025"`
}

// SumRequest — тело запроса для подсчёта суммы подписок за период
type SumRequest struct {
	UserID      *string `json:"user_id"      example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	ServiceName *string `json:"service_name" example:"Yandex Plus"`
	From        string  `json:"from"         example:"01-2025"`
	To          string  `json:"to"           example:"12-2025"`
}

// SumResponse — ответ с суммарной стоимостью подписок
type SumResponse struct {
	Total int32 `json:"total" example:"1200"`
}

// SubscriptionResponse — подписка в ответе API
type SubscriptionResponse struct {
	ID          int64   `json:"id"           example:"1"`
	ServiceName string  `json:"service_name" example:"Yandex Plus"`
	Price       int32   `json:"price"        example:"400"`
	UserID      string  `json:"user_id"      example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
	StartDate   string  `json:"start_date"   example:"07-2025"`
	EndDate     *string `json:"end_date"     example:"12-2025"`
	CreatedAt   string  `json:"created_at"   example:"2025-07-01T00:00:00Z"`
	UpdatedAt   string  `json:"updated_at"   example:"2025-07-01T00:00:00Z"`
}
