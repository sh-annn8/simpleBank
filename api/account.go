package api

import (
	"net/http"
	db "simpleBank/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createAccountParams struct {
	Owner    string `json:"owner" bindings:"required"`
	Currency string `json:"currency" bindings:"required",oneof=USD EUR`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountParams
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner: req.Owner,
		Currency: req.Currency,
		Balance: 0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}
