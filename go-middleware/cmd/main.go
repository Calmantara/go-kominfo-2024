package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/handler"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/infrastructure"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/model"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/repository"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/router"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/internal/service"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/pkg"
	"github.com/Calmantara/go-kominfo-2024/go-middleware/pkg/helper"
	"github.com/gin-gonic/gin"

	_ "github.com/Calmantara/go-kominfo-2024/go-middleware/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			GO DTS USER API DUCUMENTATION
// @version		2.0
// @description	golong kominfo 006 api documentation
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	http://www.swagger.io/support
// @contact.email	support@swagger.io
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @host			localhost:3000
// @BasePath		/
// @schemes		http
func main() {
	// requirement technical:
	// [x] middleware untuk recover ketika panic
	// [x] mengecheck basic auth
	assignment3()
}

// Product:
// authorization menggunakan jwt
// authentication bisa dilakukan dengan login
// ketika user login, akan memunculkan JWT ketika success

func server() {
	g := gin.Default()
	g.Use(gin.Recovery())

	// /public => generate JWT public
	g.GET("/public", func(ctx *gin.Context) {
		now := time.Now()

		claim := model.StandardClaim{
			Jti: fmt.Sprintf("%v", time.Now().UnixNano()),
			Iss: "go-middleware",
			Aud: "golang-006",
			Sub: "public-token",
			Exp: uint64(now.Add(time.Hour).Unix()),
			Iat: uint64(now.Unix()),
			Nbf: uint64(now.Unix()),
		}
		token, err := helper.GenerateToken(claim)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, pkg.ErrorResponse{
				Message: "error generating public token",
				Errors:  []string{err.Error()},
			})
			return
		}
		ctx.JSON(http.StatusOK, map[string]any{"token": token})
	})

	usersGroup := g.Group("/users")
	// usersGroup.Use(middleware.CheckAuthBasic)
	// usersGroup.Use(middleware.CheckAuthBearer)

	// dependency injection
	// dig by uber
	// wire

	// https://s8sg.medium.com/solid-principle-in-go-e1a624290346
	gorm := infrastructure.NewGormPostgres()
	userRepo := repository.NewUserQuery(gorm)
	// userRepoMongo := repository.NewUserQueryMongo()
	userSvc := service.NewUserService(userRepo)
	userHdl := handler.NewUserHandler(userSvc)
	userRouter := router.NewUserRouter(usersGroup, userHdl)

	// mount
	userRouter.Mount()
	// swagger
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Run(":3000")
}

func assignment3() {
	g := gin.Default()

	data := map[string]any{}

	go func() {
		for {
			// read json file
			// update json file
			data["water"] = rand.Int31n(100)
			data["wind"] = rand.Int31n(100)
			data["fire"] = rand.Int31n(100)
			data["earth"] = rand.Int31n(100)
			time.Sleep(15 * time.Second)
		}
	}()

	g.GET("/data", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]any{
			"status": data,
		})
	})
	g.Run(":3030")
}
