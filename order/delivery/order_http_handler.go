package delivery

import (
	"backend-core/common"
	"backend-core/domain"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type OrderHttpHandler struct {
	e  *echo.Echo
	ou domain.OrderUseCase
}

func (o *OrderHttpHandler) Store(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

// Update godoc
// @Summary  update order
// @Tags     order
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Param    id    path    string                   true  "Order id"
// @Param    message  body      domain.UpdateOrderDTO  true  "Update Order"
// @Router   /orders/{id} [patch]
func (o *OrderHttpHandler) Update(ctx echo.Context) error {
	c := ctx.Request().Context()
	var p domain.UpdateOrderDTO
	orderId, err := strconv.Atoi(ctx.Param("id"))
	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}
	err = o.ou.Update(c, uint(orderId), p)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, "ok")
}

// GetByUser godoc
// @Summary  Get authenticated user orders
// @Tags     order,user
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /user/orders/ [get]
func (o *OrderHttpHandler) GetByUser(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

// Fetch godoc
// @Summary  Get orders
// @Tags     order
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /orders [get]
func (o *OrderHttpHandler) Fetch(ctx echo.Context) error {
	//TODO implement me
	c := ctx.Request().Context()
	orders, err := o.ou.Fetch(c)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, orders)
}

// FetchByUser godoc
// @Summary  Get a user's order
// @Tags     order,user
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Router   /users/:id/orders/ [get]
func (o *OrderHttpHandler) FetchByUser(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}

// StoreForAuthenticateUser godoc
// @Summary  Create an order for authenticated user
// @Tags     order
// @Security  ApiKeyAuth
// @Accept   json
// @Produce  json
// @Success   200  {string}  string  "ok"
// @Param    message  body      domain.StoreForAuthenticateUserDTO  true  "Store Order"
// @Router   /user/orders [post]
func (o *OrderHttpHandler) StoreForAuthenticateUser(ctx echo.Context) error {
	c := ctx.Request().Context()
	var p domain.StoreForAuthenticateUserDTO
	id := ctx.Get("userID").(int)

	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}

	storeOrderDTO := domain.StoreOrderDTO{PlanID: p.PlanID, UserID: uint(id)}

	err := o.ou.Store(c, storeOrderDTO)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, "OK")
}

func NewOrderHttpHandler(e *echo.Echo, ou domain.OrderUseCase) domain.OrderDelivery {
	handler := &OrderHttpHandler{e, ou}
	e.GET("/orders", handler.Fetch, common.AuthMiddleWare())
	e.PATCH("/orders/:id", handler.Update, common.AuthMiddleWare())
	e.POST("/user/orders", handler.StoreForAuthenticateUser, common.AuthMiddleWare())
	return handler
}
