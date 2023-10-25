package main

import (
	"fmt"
	"log"
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

	users := server.Group("users")
	{
		users.POST("", userController.SaveUser)
		users.GET("/:id", userController.FetchUser)
	}

	items := server.Group("/:id/items")
	{
		items.POST("", itemController.SaveItem)
		items.GET("", itemController.FetchAllItems)
	}

	server.Run(fmt.Sprintf(":%v", os.Getenv("PORT")))

}
