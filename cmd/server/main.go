package main

import (
	"os"

	"github.com/curisantiago/proyectoWeb/cmd/server/handler"
	"github.com/curisantiago/proyectoWeb/internal/usuarios"
	"github.com/curisantiago/proyectoWeb/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/curisantiago/proyectoWeb/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title MELI BOOTCAMP ProyectoWeb API
// @version 1.0
// @description This API handles users and
// @author Santiago Curi
//@contact.email santiago.curi@mercadolibre.com
func main() {
	_ = godotenv.Load()

	db := store.NewStore(store.FileType, "usuario.json")
	repo := usuarios.NewRepository(db)
	service := usuarios.NewService(repo)
	u := handler.NewUser(service)

	r := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ur := r.Group("/usuarios")
	ur.GET("/", u.GetAll())
	ur.POST("/", u.Guardar())
	ur.PUT("/:id", u.Update())
	ur.PATCH("/:id", u.UpdateName())
	ur.DELETE("/:id", u.Delete())

	r.Run()
}
