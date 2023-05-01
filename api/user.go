package api

import (
	"net/http"

	db "github.com/Narawit-S/go-todo-list/db/sqlc"
	"github.com/Narawit-S/go-todo-list/utils"
	"github.com/gin-gonic/gin"
)

type createUserRequest struct {
	Email			string `json:"email" binding:"required"`
	Password 	string `json:"password" binding:"required"`
}

func (server *Server) createUser(ctx *gin.Context)  {
	var reqBody createUserRequest
	// validate req body
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	hash_password, err := utils.HashPassword(reqBody.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email: reqBody.Email,
		EncryptedPassword: hash_password,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
