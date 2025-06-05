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
