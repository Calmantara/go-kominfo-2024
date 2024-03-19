package handler

import (
	"net/http"
	"strconv"

	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/middleware"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/model"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/service"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/pkg"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	// users
	GetUsers(ctx *gin.Context)
	GetUsersById(ctx *gin.Context)
	DeleteUsersById(ctx *gin.Context)

	// activity
	UserSignUp(ctx *gin.Context)
}

type userHandlerImpl struct {
	svc service.UserService
}

func NewUserHandler(svc service.UserService) UserHandler {
	return &userHandlerImpl{
		svc: svc,
	}
}

// ShowUsers godoc
//
//	@Summary		Show users list
//	@Description	will fetch 3rd party server to get users data
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	[]model.User
//	@Failure		400	{object}	pkg.ErrorResponse
//	@Failure		404	{object}	pkg.ErrorResponse
//	@Failure		500	{object}	pkg.ErrorResponse
//	@Router			/users [get]
func (u *userHandlerImpl) GetUsers(ctx *gin.Context) {
	users, err := u.svc.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

// ShowUsersById godoc
//
//	@Summary		Show users detail
//	@Description	will fetch 3rd party server to get users data to get detail user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	model.User
//	@Failure		400	{object}	pkg.ErrorResponse
//	@Failure		404	{object}	pkg.ErrorResponse
//	@Failure		500	{object}	pkg.ErrorResponse
//	@Router			/users/{id} [get]
func (u *userHandlerImpl) GetUsersById(ctx *gin.Context) {
	// get id user
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}
	user, err := u.svc.GetUsersById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (u *userHandlerImpl) UserSignUp(ctx *gin.Context) {
	// binding sign-up body
	userSignUp := model.UserSignUp{}
	if err := ctx.Bind(&userSignUp); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	if err := userSignUp.Validate(); err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := u.svc.SignUp(ctx, userSignUp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}

	token, err := u.svc.GenerateUserAccessToken(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
	}

	ctx.JSON(http.StatusOK, map[string]any{
		"token": token,
	})
}

// DeleteUsersById godoc
//
//		@Summary		Delete user by selected id
//		@Description	will delete user with given id from param
//		@Tags			users
//		@Accept			json
//		@Produce		json
//	 	@Param 			Authorization header string true "bearer token"
//		@Param			id	path		int	true	"User ID"
//		@Success		200	{object}	model.User
//		@Failure		400	{object}	pkg.ErrorResponse
//		@Failure		404	{object}	pkg.ErrorResponse
//		@Failure		500	{object}	pkg.ErrorResponse
//		@Router			/users/{id} [delete]
func (u *userHandlerImpl) DeleteUsersById(ctx *gin.Context) {
	// get id user
	id, err := strconv.Atoi(ctx.Param("id"))
	if id == 0 || err != nil {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid required param"})
		return
	}

	// check user id session from context
	userId, ok := ctx.Get(middleware.CLAIM_USER_ID)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "invalid user session"})
		return
	}
	userIdInt, ok := userId.(float64)
	if !ok {
		ctx.JSON(http.StatusBadRequest, pkg.ErrorResponse{Message: "invalid user id session"})
		return
	}
	if id != int(userIdInt) {
		ctx.JSON(http.StatusUnauthorized, pkg.ErrorResponse{Message: "invalid user request"})
		return
	}

	user, err := u.svc.DeleteUsersById(ctx, uint64(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{Message: err.Error()})
		return
	}
	if user.ID == 0 {
		ctx.JSON(http.StatusNotFound, pkg.ErrorResponse{Message: "user not found"})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
