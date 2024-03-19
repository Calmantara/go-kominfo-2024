package router

import (
	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/handler"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter interface {
	Mount()
}

type userRouterImpl struct {
	v       *gin.RouterGroup
	handler handler.UserHandler
}

func NewUserRouter(v *gin.RouterGroup, handler handler.UserHandler) UserRouter {
	return &userRouterImpl{v: v, handler: handler}
}

func (u *userRouterImpl) Mount() {
	// activity
	// /users/sign-up
	u.v.POST("/sign-up", u.handler.UserSignUp)

	// users
	u.v.Use(middleware.CheckAuthBearer)
	// /users
	u.v.GET("", u.handler.GetUsers)
	// /users/:id
	u.v.GET("/:id", u.handler.GetUsersById)
	u.v.DELETE("/:id", u.handler.DeleteUsersById)
}
