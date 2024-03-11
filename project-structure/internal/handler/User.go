package handler

import (
	"net/http"
	"strconv"

	"github.com/Calmantara/go-kominfo-2024/project-structure/internal/service"
	"github.com/Calmantara/go-kominfo-2024/project-structure/pkg"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUsers(ctx *gin.Context)
	GetUsersById(ctx *gin.Context)
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
	ctx.JSON(http.StatusOK, user)
}
