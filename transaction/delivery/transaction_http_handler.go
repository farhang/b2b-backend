package delivery

import (
	"backend-core/common"
	"backend-core/domain"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type TransactionHttpHandler struct {
	tu domain.TransactionUseCase
}

// Store godoc
// @Summary  Create a transaction for a user
// @Tags     transaction,user
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
// @Param    id    path    string true "id"
// @Param    message  body      domain.StoreTransactionRequestDTO true  "Create transaction"
// @Success  200  {object} common.ResponseDTO
// @Router    /users/{id}/transactions [post]
func (t *TransactionHttpHandler) Store(ctx echo.Context) error {
	p := domain.StoreTransactionRequestDTO{}
	userId, _ := strconv.Atoi(ctx.Param("id"))
	c := ctx.Request().Context()
	if err := ctx.Bind(&p); err != nil {
		return err
	}

	err := t.tu.Store(c, uint(userId), p)

	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, common.ResponseDTO{
		Message: http.StatusText(http.StatusOK),
	})
}

// Fetch godoc
// @Summary  Get authenticated user transactions
// @Tags     transaction,user
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
// @Success  200  {object} common.ResponseDTO
// @Router    /users/transactions [get]
func (t *TransactionHttpHandler) Fetch(ctx echo.Context) error {
	c := ctx.Request().Context()

	transactions, err := t.tu.Fetch(c)

	transactionsResponse := make([]domain.TransactionResponseDTO, len(transactions))
	for i := range transactions {
		transactionsResponse[i] = domain.TransactionResponseDTO{
			CreatedAt:       transactions[i].CreatedAt,
			Amount:          transactions[i].Amount,
			Email:           transactions[i].User.Email,
			TransactionType: transactions[i].TransactionType.Name,
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
// @Summary  Get authenticated user transactions
// @Tags     transaction,user
// @Accept   json
// @Produce  json
// @Security  ApiKeyAuth
// @Success  200  {object} common.ResponseDTO
// @Router    /user/transactions [get]
func (t *TransactionHttpHandler) MyTransactions(ctx echo.Context) error {
	c := ctx.Request().Context()
	uid := ctx.Get("userID").(int)
	transactions, err := t.tu.FetchByUserId(c, uid)

	transactionsResponse := make([]domain.TransactionResponseDTO, len(transactions))
	for i := range transactions {
		transactionsResponse[i] = domain.TransactionResponseDTO{
			CreatedAt: transactions[i].CreatedAt,
			Amount:    transactions[i].Amount,
		}
	}

	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, common.ResponseDTO{
		Data: transactionsResponse,
	})
}

// FetchTransactionsByUserId godoc
// @Summary  Get a user's transactions
// @Tags     transaction
// @Accept   json
// @Produce  json
// @Param    id    path    string                   true  "User id"
// @Security  ApiKeyAuth
// @Success  200  {object} common.ResponseDTO
// @Router    /users/{id}/transactions/ [get]
func (t *TransactionHttpHandler) FetchTransactionsByUserId(ctx echo.Context) error {
	c := ctx.Request().Context()

	userId, err := strconv.Atoi(ctx.Param("id"))
	transactions, err := t.tu.FetchByUserId(c, userId)

	transactionsResponse := make([]domain.TransactionResponseDTO, len(transactions))
	for i := range transactions {
		transactionsResponse[i] = domain.TransactionResponseDTO{
			CreatedAt: transactions[i].CreatedAt,
			Amount:    transactions[i].Amount,
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
	ug.GET("/:id/transactions", handler.FetchTransactionsByUserId)
	ug.POST("/:id/transactions", handler.Store)
	e.GET("/user/transactions", handler.MyTransactions, common.AuthMiddleWare())
	e.GET("/users/transactions", handler.Fetch)

	return handler
}
