package service

import (
	"context"

	"fmt"

	"github.com/chelochambi/kinder-backend/internal/model"
	"github.com/chelochambi/kinder-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UsuarioService struct {
	Repo *repository.UsuarioRepository
}

func NewUsuarioService(repo *repository.UsuarioRepository) *UsuarioService {
	return &UsuarioService{Repo: repo}
}

func (s *UsuarioService) ListarUsuarios(ctx context.Context) ([]model.Usuario, error) {
	return s.Repo.GetAll(ctx)
}

func (s *UsuarioService) CrearUsuario(ctx context.Context, u *model.Usuario) error {
	// Hashear la contraseña si viene
	if u.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return fmt.Errorf("error al hashear la contraseña: %w", err)
		}
		u.PasswordHash = string(hash)
	} else {
		return fmt.Errorf("la contraseña es obligatoria")
	}

	// Llamar al repositorio
	return s.Repo.Create(ctx, u)
}

func (s *UsuarioService) ActualizarUsuario(ctx context.Context, u *model.Usuario) error {
	return s.Repo.Update(ctx, u)
}

func (s *UsuarioService) CambiarEstado(ctx context.Context, id int, estadoID int) error {
	return s.Repo.ChangeEstado(ctx, estadoID, id)
}
