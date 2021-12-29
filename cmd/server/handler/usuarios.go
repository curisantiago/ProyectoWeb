package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/curisantiago/proyectoWeb/internal/usuarios"
	"github.com/curisantiago/proyectoWeb/pkg/web"
	"github.com/gin-gonic/gin"
)

type Request struct {
	//ID              int
	Nombre          string  `json:"nombre" binding:"required"`
	Apellido        string  `json:"apellido" binding:"required"`
	Email           string  `json:"email" binding:"required"`
	Edad            int     `json:"edad" binding:"required"`
	Altura          float64 `json:"altura" binding:"required"`
	Activo          bool    `json:"activo" binding:"required"`
	FechaDeCreacion string  `json:"fechaDeCreacion" binding:"required"`
}

type Usuario struct {
	service usuarios.Service
}

func NewUser(u usuarios.Service) *Usuario {
	return &Usuario{
		service: u,
	}
}

func validarToken(ctx *gin.Context) bool {
	// verificacion de token
	token := ctx.GetHeader("token")
	fmt.Println(os.Getenv("MY_TOKEN"))
	if token != os.Getenv("MY_TOKEN") {
		ctx.JSON(401, gin.H{"error": "token invalido"})
		return false
	}
	return true

}

// ListUsers godoc
// @Sumary list all users
// @Tags Users
// @Description get users
// @Accept json
// @Return json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /usuarios [get]
func (c *Usuario) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validar token
		valido := validarToken(ctx)
		if !valido {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}

		u, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(500, web.NewResponse(500, nil, err.Error()))
			return
		}
		if len(u) == 0 {
			ctx.JSON(404, web.NewResponse(404, nil, "no hay usuarios"))
			return
		}
		ctx.JSON(200, web.NewResponse(200, u, ""))
	}
}

// StoreUsers godoc
// @Sumary store users
// @Tags Users
// @Description store users
// @Accept json
// @Return json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /usuarios [post]
func (us *Usuario) Guardar() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//Validar token
		valido := validarToken(ctx)
		if !valido {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}

		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(401, nil, fmt.Sprintf("el campo %s es requerido", err.Error())))
			return
		}
		usuarioGuardado, _ := us.service.Guardar(ctx, req.Edad, req.Nombre, req.Apellido, req.Email, req.FechaDeCreacion, req.Altura, req.Activo)
		ctx.JSON(200, usuarioGuardado)

	}
}

// UpdateUsers godoc
// @Sumary update users
// @Tags Users
// @Description update users
// @Accept json
// @Return json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /usuarios [put]
func (c *Usuario) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validar token
		valido := validarToken(ctx)
		if !valido {
			return
		}

		//Parsear id string a int
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "ID invalido"))
			return
		}

		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, web.NewResponse(401, nil, fmt.Sprintf("el campo %s es requerido", err.Error())))
			return
		}

		//Validaciones
		if req.Nombre == "" {
			ctx.JSON(401, web.NewResponse(401, nil, "El nombre  es requerido"))
			return
		}
		if req.Apellido == "" {
			ctx.JSON(401, web.NewResponse(401, nil, "El apellido  es requerido"))
			return
		}
		if req.Email == "" {
			ctx.JSON(401, web.NewResponse(401, nil, "El email  es requerido"))
			return
		}
		if req.FechaDeCreacion == "" {
			ctx.JSON(401, web.NewResponse(401, nil, "la fecha  es requerida"))
			return
		}
		if req.Altura == 0 {
			ctx.JSON(401, web.NewResponse(401, nil, "Altura requerida"))
			return
		}
		if req.Edad == 0 {
			ctx.JSON(401, web.NewResponse(401, nil, "Edad requerida"))
			return
		}

		//Hago el Update
		p, err := c.service.Update(ctx, int(id), req.Edad, req.Nombre, req.Apellido, req.Email, req.FechaDeCreacion, req.Altura, req.Activo)
		if err != nil {
			ctx.JSON(500, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, p)
	}
}

// UpdateUsers godoc
// @Sumary update name user
// @Tags Users
// @Description update name users
// @Accept json
// @Return json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /usuarios [patch]
func (c *Usuario) UpdateName() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validar token
		valido := validarToken(ctx)
		if !valido {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}

		//Parsear id string a int
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "invalid ID"})
			return
		}

		var req Request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		//Validaciones
		if req.Nombre == "" {
			ctx.JSON(400, gin.H{"error": "El nombre del producto es requerido"})
			return
		}

		//Hago el patch del producto
		p, err := c.service.UpdateName(int(id), req.Nombre)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

// DeleteUsers godoc
// @Sumary delete users
// @Tags Users
// @Description delete users from the data base
// @Param id path int true "Product ID"
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 401
// @Router /usuarios [delete]
func (c *Usuario) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//Validar token
		valido := validarToken(ctx)
		if !valido {
			ctx.JSON(401, web.NewResponse(401, nil, "token invalido"))
			return
		}

		//Parsear id string a int
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			ctx.JSON(401, web.NewResponse(401, nil, "ID invalido"))
			return
		}

		//Hago la baja fisica del producto
		err = c.service.Delete(int(id))
		if err != nil {
			ctx.JSON(404, web.NewResponse(404, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200, nil, fmt.Sprintf("El producto %d ha sido eliminado", id)))
	}
}
