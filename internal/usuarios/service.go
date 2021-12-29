package usuarios

import (
	"github.com/curisantiago/proyectoWeb/internal/domain"
	"github.com/gin-gonic/gin"
)

type Service interface {
	GetAll() ([]domain.Usuario, error)
	Guardar(ctx *gin.Context, edad int, nombre, apellido, email, fechaDeCreacion string, altura float64, activo bool) (domain.Usuario, error)
	Update(ctx *gin.Context, id, edad int, nombre, apellido, email, fechaDeCreacion string, altura float64, activo bool) (domain.Usuario, error)
	UpdateName(id int, name string) (domain.Usuario, error)
	Delete(id int) error
}

type service struct {
	repositorio Repository
}

func NewService(r Repository) Service {
	return &service{repositorio: r}

}

func (s *service) GetAll() ([]domain.Usuario, error) {
	users, error := s.repositorio.GetAll()
	if error != nil {
		return nil, error
	}
	return users, nil
}

func (s *service) Guardar(ctx *gin.Context, edad int, nombre, apellido, email, fechaDeCreacion string, altura float64, activo bool) (domain.Usuario, error) {
	users, _ := s.repositorio.GetAll()
	var id int
	if len(users) == 0 {
		id = 1
	} else {
		id = users[len(users)-1].Id + 1
	}

	//ctx.JSON(200, req)
	u, err := s.repositorio.Guardar(id, edad, nombre, apellido, email, fechaDeCreacion, altura, activo)
	if err != nil {
		return domain.Usuario{}, err
	}

	return u, nil
}

func (s *service) Update(ctx *gin.Context, id, edad int, nombre, apellido, email, fechaDeCreacion string, altura float64, activo bool) (domain.Usuario, error) {
	return s.repositorio.Update(id, edad, nombre, apellido, email, fechaDeCreacion, altura, activo)
}

func (s *service) UpdateName(id int, name string) (domain.Usuario, error) {
	return s.repositorio.UpdateName(id, name)
}

func (s *service) Delete(id int) error {
	return s.repositorio.Delete(id)
}
