package repository

import (
	"context"
	"database/sql"
	"fmt"

	"errors"
	"strings"

	"log"

	"github.com/chelochambi/kinder-backend/internal/model"
)

type UsuarioRepository struct {
	DB *sql.DB
}

func NewUsuarioRepository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{DB: db}
}

func (r *UsuarioRepository) GetAll(ctx context.Context) ([]model.Usuario, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT u.id, u.username, u.email, u.nombres, u.primer_apellido, 
		       u.segundo_apellido, u.telefono, u.foto_url, 
		       e.id, e.nombre, e.codigo
		FROM usuarios u
		JOIN estados e ON u.estado_id = e.id
		JOIN tipo_estado te ON e.tipo_estado_id = te.id
		WHERE te.codigo = 'USR'
		order by u.id
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []model.Usuario
	for rows.Next() {
		var u model.Usuario
		err := rows.Scan(
			&u.ID, &u.Username, &u.Email, &u.Nombres,
			&u.PrimerApellido, &u.SegundoApellido, &u.Telefono, &u.FotoURL,
			&u.Estado.ID, &u.Estado.Nombre, &u.Estado.Codigo, // también podés mapear el código
		)
		if err != nil {
			return nil, err
		}
		usuarios = append(usuarios, u)
	}
	return usuarios, nil
}

// Crear un nuevo usuario
func (r *UsuarioRepository) Create(ctx context.Context, u *model.Usuario) error {

	_, err := r.DB.ExecContext(ctx, `
	INSERT INTO usuarios 
		(username, email, nombres, primer_apellido, segundo_apellido, telefono, foto_url, password_hash, estado_id, creado_en, creado_por)
	VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, CURRENT_TIMESTAMP, $10)`,
		u.Username, u.Email, u.Nombres, u.PrimerApellido, u.SegundoApellido, u.Telefono,
		u.FotoURL, u.PasswordHash, u.Estado.ID, u.CreadoPor,
	)
	if err != nil {
		// Verificamos si es un error de clave duplicada
		if strings.Contains(err.Error(), "usuarios_username_key") {
			return errors.New("el usuario ya existe")
		}
		if strings.Contains(err.Error(), "usuarios_email_key") {
			return errors.New("el correo electrónico ya está registrado")
		}
		return err
	}

	return nil
}

// Actualizar usuarios
func (r *UsuarioRepository) Update(ctx context.Context, u *model.Usuario) error {
	result, err := r.DB.ExecContext(ctx, `
		UPDATE usuarios SET
			email = $1, 
			nombres = $2, 
			primer_apellido = $3,
			segundo_apellido = $4, 
			telefono = $5, 
			foto_url = $6,
			actualizado_en = CURRENT_TIMESTAMP, 
			actualizado_por = $7
		WHERE id = $8
	`, u.Email, u.Nombres, u.PrimerApellido, u.SegundoApellido, u.Telefono, u.FotoURL,
		u.ActualizadoPor, u.ID)

	log.Println("Usuario actualizado:", u.ID)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *UsuarioRepository) ChangeEstado(ctx context.Context, id int, estadoID int) error {
	log.Printf("ChangeEstado llamado con estadoID=%d para usuarioID=%d", estadoID, id)

	var existe bool
	query := `
        SELECT EXISTS (
            SELECT 1 FROM estados e
            JOIN tipo_estado te ON e.tipo_estado_id = te.id
            WHERE e.id = $1 AND e.activo = true AND te.codigo = 'USR'
        )
    `
	err := r.DB.QueryRowContext(ctx, query, estadoID).Scan(&existe)
	if err != nil {
		return err
	}

	log.Printf("Existe estado válido? %v", existe)
	if !existe {
		return fmt.Errorf("estado no válido o inactivo")
	}

	_, err = r.DB.ExecContext(ctx, `
        UPDATE usuarios SET estado_id = $1 WHERE id = $2
    `, estadoID, id)

	return err
}
