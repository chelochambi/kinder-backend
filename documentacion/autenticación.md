2. Endpoints necesarios
✅ Rutas mínimas:
POST /auth/login: Verifica credenciales y devuelve JWT.

GET /auth/me: Retorna información del usuario autenticado.

POST /auth/logout: (opcional si usás JWT sin sesión en servidor)

3. Autenticación con JWT
Usaremos JWT para mantener el sistema stateless.

Aseguramos con firmas HMAC (HS256).

Se almacena el token en localStorage (lado frontend) o en una cookie segura (HttpOnly si querés más seguridad en web).

4. Estructura del servicio auth
Podemos crear:

📁 internal/handler/auth_handler.go: para login y verificación.

📁 internal/security/jwt.go: funciones para generar y validar JWT.

📁 internal/utils/password.go: para encriptar y comparar contraseñas con bcrypt.

5. Lógica de login
✅ Flujo típico:
Recibe username y password.

Busca el usuario en la DB.

Compara password con PasswordHash usando bcrypt.

Si todo está OK, genera un token JWT con:

ID del usuario

Rol (si lo tenés)

Tiempo de expiración (ej. 15 min o 1h)

6. Seguridad recomendada
Usar bcrypt con un costo razonable (bcrypt.DefaultCost o 12).

Tokens JWT firmados con una clave secreta segura.

No exponer información sensible ni en logs ni en errores detallados.

Middleware para proteger rutas con JWT (ej: AuthMiddleware).

7. Bonus escalabilidad
En el futuro podés agregar:

Refresh tokens

Revocación de tokens (blacklist)

Auditoría de inicio de sesión

OAuth2 o inicio con Google/Microsoft si aplica


go get github.com/golang-jwt/jwt/v5
go get golang.org/x/crypto/bcrypt
go get github.com/jackc/pgx/v5/pgconn@v5.7.5
go get github.com/rs/cors


estructura json ejemplo:
´´´
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NDkyMjgwMjMsIm5hbWUiOiJtYXJjZWxvLmNoYW1iaSIsInN1YiI6MX0.3q25wGIctOrAgkQGvq_lmM04Ojt3_P7mDOIaNOJ9s0U",
    "usuario": {
        "id": 1,
        "username": "marcelo.chambi",
        "email": "admin@demo.com",
        "nombres": "Marcelo",
        "primer_apellido": "Chambi",
        "segundo_apellido": "Paredes",
        "telefono": "77777777",
        "foto_url": "",
        "roles": [
            "ADMIN"
        ],
        "permisos": null,
        "menus": [
            {
                "id": 1,
                "nombre": "Inicio",
                "icono": "fas fa-home",
                "ruta": "/",
                "padre_id": null,
                "tipo": "sidebar",
                "mostrar": true,
                "permisos": [
                    "LIS"
                ]
            },
            {
                "id": 2,
                "nombre": "Seguridad",
                "icono": "fas fa-lock",
                "ruta": "/seguridad",
                "padre_id": null,
                "tipo": "sidebar",
                "mostrar": true,
                "permisos": [
                    "ACT",
                    "CRE",
                    "ELI",
                    "LIS"
                ],
                "submenus": [
                    {
                        "id": 5,
                        "nombre": "Usuarios",
                        "icono": "fas fa-user",
                        "ruta": "/seguridad/usuarios",
                        "padre_id": null,
                        "tipo": "sidebar",
                        "mostrar": true,
                        "permisos": [
                            "ACT",
                            "CRE",
                            "ELI",
                            "LIS"
                        ]
                    },
                    {
                        "id": 6,
                        "nombre": "Roles",
                        "icono": "fas fa-user-tag",
                        "ruta": "/seguridad/roles",
                        "padre_id": null,
                        "tipo": "sidebar",
                        "mostrar": true,
                        "permisos": [
                            "ACT",
                            "CRE",
                            "ELI",
                            "LIS"
                        ]
                    },
                    {
                        "id": 7,
                        "nombre": "Permisos",
                        "icono": "fas fa-key",
                        "ruta": "/seguridad/permisos",
                        "padre_id": null,
                        "tipo": "sidebar",
                        "mostrar": true,
                        "permisos": [
                            "ACT",
                            "CRE",
                            "ELI",
                            "LIS"
                        ]
                    },
                    {
                        "id": 8,
                        "nombre": "Menús",
                        "icono": "fas fa-list",
                        "ruta": "/seguridad/menus",
                        "padre_id": null,
                        "tipo": "sidebar",
                        "mostrar": true,
                        "permisos": [
                            "ACT",
                            "CRE",
                            "ELI",
                            "LIS"
                        ]
                    }
                ]
            },
            {
                "id": 3,
                "nombre": "Sucursales",
                "icono": "fas fa-building",
                "ruta": "/sucursales",
                "padre_id": null,
                "tipo": "sidebar",
                "mostrar": true,
                "permisos": [
                    "LIS"
                ]
            },
            {
                "id": 4,
                "nombre": "Clientes",
                "icono": "fas fa-users",
                "ruta": "/clientes",
                "padre_id": null,
                "tipo": "sidebar",
                "mostrar": true,
                "permisos": [
                    "LIS"
                ]
            }
        ],
        "sucursales": [
            {
                "id": 1,
                "nombre": "Central"
            },
            {
                "id": 2,
                "nombre": "Norte"
            }
        ]
    }
}
´´´