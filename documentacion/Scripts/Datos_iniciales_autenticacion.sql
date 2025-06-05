
-- =======================
-- Insertar datos en tipo_estado
-- =======================
INSERT INTO tipo_estado (id, nombre, descripcion, codigo, creado_por, creado_en)
VALUES 
  (1, 'General', 'Estados de uso general del sistema', 'GEN', 1, CURRENT_TIMESTAMP),
  (2, 'Seguridad', 'Estados relacionados a usuarios, roles, etc.', 'SEG', 1, CURRENT_TIMESTAMP);

-- =======================
-- Insertar datos en estados
-- =======================
INSERT INTO estados (id, tipo_estado_id, nombre, codigo, descripcion, activo, orden, creado_por, creado_en)
VALUES 
  (1, 1, 'Activo', 'ACT', 'Elemento en estado activo', TRUE, 1, 1, CURRENT_TIMESTAMP),
  (2, 1, 'Inactivo', 'INA', 'Elemento en estado inactivo', TRUE, 2, 1, CURRENT_TIMESTAMP),
  (3, 2, 'Habilitado', 'HAB', 'Elemento de seguridad habilitado', TRUE, 1, 1, CURRENT_TIMESTAMP),
  (4, 2, 'Deshabilitado', 'DES', 'Elemento de seguridad deshabilitado', TRUE, 2, 1, CURRENT_TIMESTAMP),
  ;

-- =======================
-- Insertar un usuario admin
-- contraseña: admin123
-- =======================
INSERT INTO usuarios (
  id, username, email, password_hash, nombres, primer_apellido, segundo_apellido, estado_id, telefono, foto_url, creado_por, creado_en
) VALUES (
  1, 'marcelo.chambi', 'admin@demo.com', '$2a$10$ShA2eCyjautV9HjkTx2.auUimgetTZw56PtUkolwOvJEnaYqgELiC', 'Marcelo', 'Chambi', 'Paredes',
  3, '77777777', NULL, 1, CURRENT_TIMESTAMP
);

-- =======================
-- Insertar roles
-- =======================
INSERT INTO roles (id, nombre, codigo, descripcion, estado_id, creado_por, creado_en)
VALUES 
  (1, 'Administrador', 'ADMIN', 'Rol con todos los permisos', 3, 1, CURRENT_TIMESTAMP);

-- =======================
-- Insertar usuario_rol
-- =======================
INSERT INTO usuario_rol (usuario_id, rol_id, estado_id, vigente_desde, creado_por, creado_en)
VALUES 
  (1, 1, 3, CURRENT_DATE, 1, CURRENT_TIMESTAMP);

-- =======================
-- Insertar menús (2 niveles)
-- =======================
INSERT INTO menus (id, nombre, icono, ruta, orden, tipo, mostrar, padre_id, estado_id, creado_por, creado_en)
VALUES
  (1, 'Inicio', 'fas fa-home', '/inicio', 1, 'L', TRUE, NULL, 0, 1, CURRENT_TIMESTAMP),
  (2, 'Seguridad', 'fas fa-lock', NULL, 2, 'L', TRUE, NULL, 0, 1, CURRENT_TIMESTAMP),
  (3, 'Sucursales', 'fas fa-building', '/sucursales', 3, 'L', TRUE, NULL, 0, 1, CURRENT_TIMESTAMP),
  (4, 'Clientes', 'fas fa-users', '/clientes', 4, 'L', TRUE, NULL, 0, 1, CURRENT_TIMESTAMP),
  (5, 'Usuarios', 'fas fa-user', '/seguridad/usuarios', 1, 'L', TRUE, 2, 2, 1, CURRENT_TIMESTAMP),
  (6, 'Roles', 'fas fa-user-tag', '/seguridad/roles', 2, 'L', TRUE, 2, 2, 1, CURRENT_TIMESTAMP),
  (7, 'Permisos', 'fas fa-key', '/seguridad/permisos', 3, 'L', TRUE, 2, 2, 1, CURRENT_TIMESTAMP),
  (8, 'Menús', 'fas fa-list', '/seguridad/menus', 4, 'L', TRUE, 2, 2, 1, CURRENT_TIMESTAMP);

-- =======================
-- Insertar permisos
-- =======================
INSERT INTO permisos (id, nombre, codigo, descripcion, menu_id, accion, estado_id, creado_por, creado_en)
VALUES 
  (1, 'Listar', 'USUARIOS_VIEW', 'Permite ver usuarios', 5, 'view', 3, 1, CURRENT_TIMESTAMP),
  (2, 'Crear', 'USUARIOS_CREATE', 'Permite crear usuarios', 5, 'create', 3, 1, CURRENT_TIMESTAMP),
  (3, 'Modificar', 'ROLES_VIEW', 'Permite ver roles', 6, 'view', 3, 1, CURRENT_TIMESTAMP),
  (4, 'Eliminar', 'PERMISOS_VIEW', 'Permite ver permisos', 7, 'view', 3, 1, CURRENT_TIMESTAMP);

-- =======================
-- Insertar rol_permiso
-- =======================
INSERT INTO rol_permiso (rol_id, permiso_id, estado_id, vigente_desde, creado_por, creado_en)
SELECT 1, id, 3, CURRENT_DATE, 1, CURRENT_TIMESTAMP FROM permisos;

-- =======================
-- Insertar sucursales
-- =======================
INSERT INTO sucursales (id, nombre, codigo, descripcion, direccion, telefono, email, estado_id, creado_por, creado_en)
VALUES 
  (1, 'Central', 'SUC-CEN', 'Sucursal principal', 'Av. Principal #123', '71234567', 'central@demo.com', 1, 1, CURRENT_TIMESTAMP),
  (2, 'Norte', 'SUC-NOR', 'Sucursal zona norte', 'Av. Norte #45', '73456789', 'norte@demo.com', 1, 1, CURRENT_TIMESTAMP);

-- =======================
-- Asociar usuario a sucursales
-- =======================
INSERT INTO usuario_sucursal (usuario_id, sucursal_id, estado_id, creado_por, creado_en)
VALUES 
  (1, 1, 1, 1, CURRENT_TIMESTAMP),
  (1, 2, 1, 1, CURRENT_TIMESTAMP);
