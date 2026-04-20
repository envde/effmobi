package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/envde/effmobi/interanl/service"
	"github.com/envde/effmobi/interanl/transport/dto"
	"github.com/envde/effmobi/interanl/transport/response"
)

// убеждаемся что db используется в аннотациях
var _ = dto.SubscriptionResponse{}

type SubscriptionHandler struct {
	svc *service.SubscriptionService
}

func NewSubscriptionHandler(svc *service.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{svc: svc}
}

// Create godoc
// @Summary      Создать подписку
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        body body dto.CreateSubscriptionRequest true "Данные подписки"
// @Success      201 {object} dto.SubscriptionResponse
// @Failure      400 {object} map[string]string
// @Failure      422 {object} map[string]string
// @Router       /subscriptions [post]
func (h *SubscriptionHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	sub, err := h.svc.Create(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, sub)
}

// Get godoc
// @Summary      Получить подписку по ID
// @Tags         subscriptions
// @Produce      json
// @Param        id path int true "ID подписки"
// @Success      200 {object} dto.SubscriptionResponse
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Router       /subscriptions/{id} [get]
func (h *SubscriptionHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	sub, err := h.svc.Get(r.Context(), id)
	if err != nil {
		response.Error(w, http.StatusNotFound, "subscription not found")
		return
	}

	response.JSON(w, http.StatusOK, sub)
}

// List godoc
// @Summary      Список всех подписок
// @Tags         subscriptions
// @Produce      json
// @Success      200 {array}  dto.SubscriptionResponse
// @Failure      500 {object} map[string]string
// @Router       /subscriptions [get]
func (h *SubscriptionHandler) List(w http.ResponseWriter, r *http.Request) {
	subs, err := h.svc.List(r.Context())
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, subs)
}

// Update godoc
// @Summary      Обновить подписку
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        id   path int true "ID подписки"
// @Param        body body dto.UpdateSubscriptionRequest true "Новые данные подписки"
// @Success      200 {object} dto.SubscriptionResponse
// @Failure      400 {object} map[string]string
// @Failure      422 {object} map[string]string
// @Router       /subscriptions/{id} [put]
func (h *SubscriptionHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	var req dto.UpdateSubscriptionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	sub, err := h.svc.Update(r.Context(), id, req)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, sub)
}

// Delete godoc
// @Summary      Удалить подписку
// @Tags         subscriptions
// @Param        id path int true "ID подписки"
// @Success      204
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /subscriptions/{id} [delete]
func (h *SubscriptionHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	if err := h.svc.Delete(r.Context(), id); err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Sum godoc
// @Summary      Подсчёт суммарной стоимости подписок за период
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        body body dto.SumRequest true "Фильтры и период"
// @Success      200 {object} dto.SumResponse
// @Failure      400 {object} map[string]string
// @Failure      422 {object} map[string]string
// @Router       /subscriptions/sum [post]
func (h *SubscriptionHandler) Sum(w http.ResponseWriter, r *http.Request) {
	var req dto.SumRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	total, err := h.svc.Sum(r.Context(), req)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	response.JSON(w, http.StatusOK, dto.SumResponse{Total: total})
}
