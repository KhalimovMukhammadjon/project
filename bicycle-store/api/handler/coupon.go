package handler

import (
	"app/api/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCoupon(c *gin.Context) {

	var createCoupon models.CreateCoupon

	err := c.ShouldBindJSON(&createCoupon) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create coupon", http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.storages.Coupon().CreateCoupon(context.Background(), &createCoupon)
	if err != nil {
		h.handlerResponse(c, "storage.coupon.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.Coupon().GetByID(context.Background(), &models.CouponPrimaryKey{CouponID: id})
	if err != nil {
		h.handlerResponse(c, "storage.coupon.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create coupon", http.StatusCreated, resp)
}

func (h *Handler) GetByIdCoupon(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.coupon.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	resp, err := h.storages.Coupon().GetByID(context.Background(), &models.CouponPrimaryKey{CouponID: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.coupon.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get coupon by id", http.StatusCreated, resp)
}

func (h *Handler) GetListCoupon(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list coupon", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list coupon", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Coupon().GetList(context.Background(), models.GetListCouponRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})

	if err != nil {
		h.handlerResponse(c, "storage.coupon.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list coupon response", http.StatusOK, resp)
}

func (h *Handler) DeleteCoupon(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.coupon.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	rowsAffected, err := h.storages.Coupon().Delete(context.Background(), &models.CouponPrimaryKey{CouponID: idInt})
	if err != nil {
		h.handlerResponse(c, "storage.coupon.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.coupon.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete coupon", http.StatusNoContent, nil)
}
