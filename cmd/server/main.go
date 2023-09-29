package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Pixelcutter/units_backend/cmd/server/controller"
	"github.com/Pixelcutter/units_backend/cmd/server/repository"
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

	dbPath := os.Getenv("DB_PATH")
	var (
		unitsRepo      repository.UnitsRepository = repository.NewPostgresRepo(dbPath)
		userService    service.UserService        = service.NewUserService(unitsRepo)
		userController controller.UserController  = controller.NewUserController(userService)
		itemService    service.ItemService        = service.NewItemService(unitsRepo)
		itemController controller.ItemController  = controller.NewItemController(itemService)
	)

	defer unitsRepo.CloseDB()
	server := gin.Default()

	server.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(
			http.StatusOK,
			userController.FetchAllUsers(),
		)
	})

	server.GET("/users/:name", func(ctx *gin.Context){
		ctx.String(
			http.StatusOK, 
			"hello, %s", 
			ctx.Param("name"),
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
			http.StatusOK,
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
