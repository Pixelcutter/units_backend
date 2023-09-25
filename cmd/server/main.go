package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Pixelcutter/units_backend/cmd/server/controller"
	"github.com/Pixelcutter/units_backend/cmd/server/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Test struct {
	Something string `json:"blah"`
	SomeInt   int    `json:"blah_int"`
}

func main() {
	// loading .env vars
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env")
	}

	server := gin.Default()
	var (
		userService    service.UserService       = service.NewUserService()
		userController controller.UserController = controller.NewUserController(userService)
		itemService    service.ItemService       = service.NewItemService()
		itemController controller.ItemController = controller.NewItemController(itemService)
	)

	server.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusAccepted,
			userController.FetchAllUsers(),
		)
	})

	server.POST("/users", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusCreated,
			userController.SaveUser(ctx),
		)
	})

	server.GET("/items", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusAccepted,
			itemController.FetchAllItems(),
		)
	})

	server.POST("/items", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusCreated,
			itemController.SaveItem(ctx),
		)
	})

	server.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))

}
