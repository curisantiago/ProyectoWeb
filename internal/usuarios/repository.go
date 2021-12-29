package usuarios

import (
	"fmt"

	"github.com/curisantiago/proyectoWeb/internal/domain"
	"github.com/curisantiago/proyectoWeb/pkg/store"
)

//var Usuarios []domain.Usuario

type Repository interface {
	GetAll() ([]domain.Usuario, error)
	Guardar(id, edad int, nombre, apellido, email, fechaDeCreacion string, altura float64, activo bool) (domain.Usuario, error)
	Update(id, edad int, nombre, apellido, email, fechaDeCreacion string, altura float64, activo bool) (domain.Usuario, error)
	UpdateName(id int, name string) (domain.Usuario, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll() ([]domain.Usuario, error) {
	var Usuarios []domain.Usuario
	r.db.Read(&Usuarios)
	return Usuarios, nil
}

func (r *repository) Guardar(id, edad int, nombre, apellido, email, fechaDeCreacion string, altura float64, activo bool) (domain.Usuario, error) {
	var Usuarios []domain.Usuario
	r.db.Read(&Usuarios)
	//newUser := domain.Usuario{id, nombre, apellido, email, edad, altura, activo, fechaDeCreacion}
	newUser := domain.Usuario{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaDeCreacion: fechaDeCreacion}
	Usuarios = append(Usuarios, newUser)
	if err := r.db.Write(Usuarios); err != nil {
		return domain.Usuario{}, err
	}
	return newUser, nil
}

func (r *repository) Update(id, edad int, nombre, apellido, email, fechaDeCreacion string, altura float64, activo bool) (domain.Usuario, error) {
	var Usuarios []domain.Usuario
	r.db.Read(&Usuarios)
	updateUser := domain.Usuario{Id: id, Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaDeCreacion: fechaDeCreacion}
	updated := false
	for i := range Usuarios {
		if Usuarios[i].Id == id {
			Usuarios[i] = updateUser
			updated = true
		}
	}
	if !updated {
		return domain.Usuario{}, fmt.Errorf("usuario %d no encontrado", id)
	}
	return updateUser, nil
}

func (r *repository) UpdateName(id int, name string) (domain.Usuario, error) {
	var Usuarios []domain.Usuario
	r.db.Read(&Usuarios)
	for i := range Usuarios {
		if Usuarios[i].Id == id {
			Usuarios[i].Nombre = name
			if err := r.db.Write(Usuarios); err != nil {
				return domain.Usuario{}, err
			}
			return Usuarios[i], nil
		}
	}
	return domain.Usuario{}, fmt.Errorf("usuario %d no encontrado", id)
}

func (r *repository) Delete(id int) error {
	var Usuarios []domain.Usuario
	r.db.Read(&Usuarios)
	deleted := false
	var index int
	for i := range Usuarios {
		if Usuarios[i].Id == id {
			index = i
			deleted = true
		}
	}
	if !deleted {
		return fmt.Errorf("usuario %d no encontrado", id)
	}
	Usuarios = append(Usuarios[:index], Usuarios[index+1:]...)
	if err := r.db.Write(Usuarios); err != nil {
		return err
	}
	return nil
}
