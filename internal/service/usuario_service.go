package service

import (
	"context"

	"github.com/chelochambi/kinder-backend/internal/model"
	"github.com/chelochambi/kinder-backend/internal/repository"
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
	return s.Repo.Create(ctx, u)
}

func (s *UsuarioService) ActualizarUsuario(ctx context.Context, u *model.Usuario) error {
	return s.Repo.Update(ctx, u)
}

func (s *UsuarioService) CambiarEstado(ctx context.Context, id int, estadoID int) error {
	return s.Repo.ChangeEstado(ctx, estadoID, id)
}
