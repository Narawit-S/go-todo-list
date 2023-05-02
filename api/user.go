package api

import (
	"net/http"

	db "github.com/Narawit-S/go-todo-list/db/sqlc"
	"github.com/Narawit-S/go-todo-list/utils"
	"github.com/gin-gonic/gin"
)

type SignUpRequest struct {
	Email			string `json:"email" binding:"required"`
	Password 	string `json:"password" binding:"required"`
}

func (server *Server) SignUp(ctx *gin.Context)  {
	var reqBody SignUpRequest
	// validate req body
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		return
	}

	hash_password, err := utils.HashPassword(reqBody.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		return
	}

	arg := db.CreateUserParams{
		Email: reqBody.Email,
		EncryptedPassword: hash_password,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

type SignInRequest struct {
	Email			string	`json:"email" binding:"required"`
	Password	string	`json:"password" binding:"required"`
}

func (server *Server) SignIn(ctx *gin.Context)  {
	var reqBody SignInRequest
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err.Error()))
		return
	}

	user, err := server.store.GetUser(ctx, reqBody.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse("incorrect email"))
		return
	}

	if err := utils.CheckPassword(reqBody.Password, user.EncryptedPassword); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse("incorrect password"))
		return
	}

	ctx.JSON(http.StatusOK, user)
}
