package delivery

import (
	"backend-core/common"
	"backend-core/domain"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
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
func (j *TransactionHttpHandler) Fetch(ctx echo.Context) error {
	c := ctx.Request().Context()

	transactions, err := j.tu.Fetch(c)

	transactionsResponse := make([]domain.TransactionResponseDTO, len(transactions))
	for i := range transactions {
		transactionsResponse[i] = domain.TransactionResponseDTO{
			CreatedAt:       transactions[i].CreatedAt,
			TransactionType: transactions[i].TransactionType,
			Amount:          transactions[i].Amount,
			Email:           transactions[i].User.Email,
		}
	}

	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data: transactionsResponse,
	})
}

// MyTransactions godoc
// @Summary   Get user information
// @Tags     transaction
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
// @Success  200  {object} common.ResponseDTO
// @Router    /transactions/me [get]
func (j *TransactionHttpHandler) MyTransactions(ctx echo.Context) error {
	c := ctx.Request().Context()
	uid := ctx.Get("userID").(int)
	transactions, err := j.tu.FetchByUserId(c, uid)

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
// @Security  ApiKeyAuth
// @Success  200      {object}  common.ResponseDTO
// @Router   /transactions/deposit [post]
func (j *TransactionHttpHandler) Deposit(ctx echo.Context) error {
	c := ctx.Request().Context()
	var p domain.DepositRequestDTO

	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}
	err := j.tu.Deposit(c, p)

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
// @Security  ApiKeyAuth
// @Success  200      {object}  common.ResponseDTO
// @Router   /transactions/withdraw [post]
func (j *TransactionHttpHandler) WithDraw(ctx echo.Context) error {
	c := ctx.Request().Context()
	var p domain.WithDrawRequestDTO

	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}

	err := j.tu.WithDraw(c, p)

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

func (j *TransactionHttpHandler) Profit(ctx echo.Context) error {
	c := ctx.Request().Context()
	var p domain.ProfitRequestDTO

	if err := ctx.Bind(&p); err != nil {
		return common.ErrHttpBadRequest(err)
	}
	err := j.tu.Profit(c, p)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, common.ResponseDTO{
		Message: "the transaction is created successfully",
	})
}

// FetchTransactionsByUserId godoc
// @Summary   Get user information
// @Tags     user
// @Accept   json
// @Produce  json
// @Param    id    path    string                   true  "User id"
// @Security  ApiKeyAuth
// @Success  200  {object} common.ResponseDTO
// @Router    /users/{id}/transactions/ [get]
func (j *TransactionHttpHandler) FetchTransactionsByUserId(ctx echo.Context) error {
	c := ctx.Request().Context()

	userId, err := strconv.Atoi(ctx.Param("id"))
	transactions, err := j.tu.FetchByUserId(c, userId)

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

func NewTransactionHttpHandler(e *echo.Echo, tu domain.TransactionUseCase) domain.TransactionHttpHandler {
	handler := &TransactionHttpHandler{tu}

	ug := e.Group("users", common.AuthMiddleWare(), common.CASBINMiddleWare())
	ug.GET("/:id:/transactions/", handler.FetchTransactionsByUserId)
	tg := e.Group("transactions", common.AuthMiddleWare(), common.CASBINMiddleWare())
	tg.GET("/me", handler.MyTransactions)
	tg.GET("/", handler.Fetch)
	tg.POST("/deposit", handler.Deposit)
	tg.POST("/withdraw", handler.WithDraw)
	tg.POST("/profit", handler.Profit)
	return handler
}
