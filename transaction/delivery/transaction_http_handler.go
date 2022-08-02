package delivery

import (
	"backend-core/common"
	"backend-core/domain"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type TransactionHttpHandler struct {
	tu domain.TransactionUseCase
}

// Fetch godoc
// @Summary   Get user information
// @Tags     transaction
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
// @Success  200  {object} common.ResponseDTO
// @Router    /transactions/ [get]
func (t *TransactionHttpHandler) Fetch(ctx echo.Context) error {
	c := ctx.Request().Context()
	uid := ctx.Get("userID").(int)
	transactions, err := t.tu.FetchByUserId(c, uid)

	transactionsResponse := make([]domain.TransactionResponseDTO, len(transactions))
	for i := range transactions {
		transactionsResponse[i] = domain.TransactionResponseDTO{
			CreatedAt:       transactions[i].CreatedAt,
			TransactionType: transactions[i].TransactionType,
			Amount:          transactions[i].Amount,
		}
	}

	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data: transactionsResponse,
	})
}

// Deposit  godoc
// @Summary  create deposit transaction
// @Tags     transaction
// @Accept   json
// @Produce  json
// @Param    message  body      domain.DepositRequestDTO true  "Registration data"
// @Success  200      {object}  common.ResponseDTO
// @Router   /transactions/deposit [post]
func (t *TransactionHttpHandler) Deposit(ctx echo.Context) error {
	c := ctx.Request().Context()
	var p domain.DepositRequestDTO

	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}
	err := t.tu.Deposit(c, p)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, common.ResponseDTO{
		Message: "the transaction is created successfully",
	})
}

// WithDraw  godoc
// @Summary  create withdraw transaction
// @Tags     transaction
// @Accept   json
// @Produce  json
// @Param    message  body      domain.WithDrawRequestDTO true  "Withdraw data"
// @Success  200      {object}  common.ResponseDTO
// @Router   /transactions/withdraw [post]
func (t *TransactionHttpHandler) WithDraw(ctx echo.Context) error {
	c := ctx.Request().Context()
	var p domain.WithDrawRequestDTO

	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}

	err := t.tu.WithDraw(c, p)

	if errors.Is(domain.ErrNotEnoughAssetAmount, err) {
		return common.ErrHttpBadRequest(err)
	}

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, common.ResponseDTO{
		Message: "the transaction is created successfully",
	})
}

// Profit godoc
// @Summary  create withdraw transaction
// @Tags     transaction
// @Accept   json
// @Produce  json
// @Param    message  body      domain.ProfitRequestDTO true  "Withdraw data"
// @Success  200      {object}  common.ResponseDTO
// @Router   /transactions/profit [post]
func (t *TransactionHttpHandler) Profit(ctx echo.Context) error {
	c := ctx.Request().Context()
	var p domain.ProfitRequestDTO

	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}
	err := t.tu.Profit(c, p)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, common.ResponseDTO{
		Message: "the transaction is created successfully",
	})
}

func NewTransactionHttpHandler(e *echo.Echo, tu domain.TransactionUseCase) domain.TransactionHttpHandler {
	handler := &TransactionHttpHandler{tu}
	tg := e.Group("transactions")
	tg.GET("/", handler.Fetch, common.AuthMiddleWare())
	tg.POST("/deposit", handler.Deposit)
	tg.POST("/withdraw", handler.WithDraw)
	tg.POST("/profit", handler.Profit)
	return handler
}
