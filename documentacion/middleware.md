Middleware de autenticación
Con esto vas a poder proteger cualquier ruta del backend y permitir el acceso solo a usuarios autenticados.

graph TD
    subgraph Internal/Handler
        LoginHandler["auth_handler.go\n(LoginHandler)"]
        AuthMeHandler["auth_me_handler.go\n(AuthMeHandler)"]
        UsuarioHandler["usuario_handler.go\n(ListarUsuariosHandler)"]
    end

    subgraph Internal/Middleware
        AuthMiddleware["auth_middleware.go\n(AuthMiddleware)"]
    end

    subgraph Internal/Util
        JWT["jwt.go\n(CrearToken / VerificarToken)"]
    end

    Router["router.go"] --> LoginHandler
    Router --> AuthMeHandler
    Router --> UsuarioHandler

    LoginHandler --> JWT
    AuthMiddleware --> JWT
    Router --> AuthMiddleware



router.go
│
├──→ auth_handler.go       (LoginHandler)
│       └──→ jwt.go        (crear/verificar tokens)
│
├──→ auth_me_handler.go    (AuthMeHandler)
│       └──→ auth_middleware.go  (requiere JWT válido)
│
├──→ usuario_handler.go    (ListarUsuariosHandler)
│       └──→ auth_middleware.go  (protegido con token)
│
└──→ tipo_estado_handler.go (público o protegido según config)
