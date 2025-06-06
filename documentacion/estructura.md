D:.
│   .env
│   go.mod
│   go.sum
│
├───.vscode
│       launch.json
│
├───cmd
│   └───server
│           main.go
│
├───documentacion
│   │   autenticación.md
│   │   estructura.md
│   │   middleware.md
│   │
│   └───Scripts
│           Datos_iniciales_autenticacion.sql
│           Estructura_inicio_autenticacion.sql
│
└───internal
    ├───db
    │       db.go
    │
    ├───handler
    │       auth_handler.go
    │       auth_me_handler.go
    │       tipo_estado_handler.go
    │       usuario_handler.go
    │
    ├───middleware
    │       auth_middleware.go
    │
    ├───model
    │       estado.go
    │       menu.go
    │       permiso.go
    │       rol.go
    │       sucursal.go
    │       tipo_estado.go
    │       usuario.go
    │       usuarioInfo.go
    │
    ├───repository
    ├───router
    │       router.go
    │
    ├───security
    │       jwt.go
    │
    ├───service
    └───utils
            password.go