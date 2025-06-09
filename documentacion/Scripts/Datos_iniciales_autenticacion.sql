
-- =======================
-- Insertar datos en tipo_estado
-- =======================
INSERT INTO tipo_estado (id, nombre, descripcion, codigo, creado_por, creado_en)
VALUES 
  (1, 'General', 'Estados de uso general del sistema', 'GEN', 1, CURRENT_TIMESTAMP),
  (2, 'Seguridad', 'Estados relacionados a usuarios, roles, etc.', 'SEG', 1, CURRENT_TIMESTAMP),
  (3, 'Usuario', 'Estados específicos para usuarios', 'USR', 1, CURRENT_TIMESTAMP);

-- =======================
-- Insertar datos en estados
-- =======================
INSERT INTO estados (id, tipo_estado_id, nombre, codigo, descripcion, activo, orden, creado_por, creado_en)
VALUES 
  (1, 1, 'Activo', 'ACT', 'Elemento en estado activo', TRUE, 1, 1, CURRENT_TIMESTAMP),
  (2, 1, 'Inactivo', 'INA', 'Elemento en estado inactivo', TRUE, 2, 1, CURRENT_TIMESTAMP),
  (3, 2, 'Habilitado', 'HAB', 'Elemento de seguridad habilitado', TRUE, 1, 1, CURRENT_TIMESTAMP),
  (4, 2, 'Deshabilitado', 'DES', 'Elemento de seguridad deshabilitado', TRUE, 2, 1, CURRENT_TIMESTAMP),
  (5, 3, 'Activo', 'UACT', 'Usuario activo', TRUE, 1, 1, CURRENT_TIMESTAMP),
  (6, 3, 'Bloqueado', 'UBLQ', 'Usuario bloqueado', TRUE, 2, 1, CURRENT_TIMESTAMP),
  (7, 3, 'Bloqueado aut', 'UBLQ', 'Usuario bloqueado autimático', TRUE, 3, 1, CURRENT_TIMESTAMP);

-- =======================
-- Insertar un usuario admin
-- contraseña: admin123
-- =======================
INSERT INTO usuarios (
  id, username, email, password_hash, nombres, primer_apellido, segundo_apellido, estado_id, telefono, foto_url, creado_por, creado_en
) VALUES (
  1, 'marcelo.chambi', 'admin@demo.com', '$2a$10$ShA2eCyjautV9HjkTx2.auUimgetTZw56PtUkolwOvJEnaYqgELiC', 'Marcelo', 'Chambi', 'Paredes',
  5, '77777777', NULL, 1, CURRENT_TIMESTAMP
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
  (1, 'Inicio', 'fas fa-home', '/', 1, 'sidebar', TRUE, NULL, 3, 1, CURRENT_TIMESTAMP),
  (2, 'Seguridad', 'fas fa-lock', '/seguridad', 2, 'sidebar', TRUE, NULL, 3, 1, CURRENT_TIMESTAMP),
  (3, 'Sucursales', 'fas fa-building', '/sucursales', 3, 'sidebar', TRUE, NULL, 3, 1, CURRENT_TIMESTAMP),
  (4, 'Clientes', 'fas fa-users', '/clientes', 4, 'sidebar', TRUE, NULL, 3, 1, CURRENT_TIMESTAMP),
  (5, 'Usuarios', 'fas fa-user', '/seguridad/usuarios', 1, 'sidebar', TRUE, 2, 3, 1, CURRENT_TIMESTAMP),
  (6, 'Roles', 'fas fa-user-tag', '/seguridad/roles', 2, 'sidebar', TRUE, 2, 3, 1, CURRENT_TIMESTAMP),
  (7, 'Permisos', 'fas fa-key', '/seguridad/permisos', 3, 'sidebar', TRUE, 2, 3, 1, CURRENT_TIMESTAMP),
  (8, 'Menús', 'fas fa-list', '/seguridad/menus', 4, 'sidebar', TRUE, 2, 3, 1, CURRENT_TIMESTAMP);

-- =======================
-- Insertar permisos
-- =======================
INSERT INTO permisos (id, nombre, codigo, descripcion,  estado_id, creado_por, creado_en)
VALUES 
  (1, 'Listar', 'LIS', 'Permite Listar registros',  3, 1, CURRENT_TIMESTAMP),
  (2, 'Crear', 'CRE', 'Permite crear registros', 3, 1, CURRENT_TIMESTAMP),
  (3, 'Actualizar', 'ACT', 'Permite actualizar registros', 3, 1, CURRENT_TIMESTAMP),
  (4, 'Eliminar', 'ELI', 'Permite cambiar de estado  deshabilitado a registros', 3, 1, CURRENT_TIMESTAMP);

-- =======================
-- Insertar rol menus permisos
-- =======================
INSERT INTO rol_menu_permiso (id , rol_id, menu_id, permiso_id, estado_id, creado_por, creado_en)
VALUES
  (1, 1, 1, 1, 3, 1, CURRENT_TIMESTAMP), -- Inicio - Listar
  (2, 1, 2, 1, 3, 1, CURRENT_TIMESTAMP), -- Seguridad - Listar
  (3, 1, 2, 2, 3, 1, CURRENT_TIMESTAMP), -- Seguridad - Crear
  (4, 1, 2, 3, 3, 1, CURRENT_TIMESTAMP), -- Seguridad - Actualizar
  (5, 1, 2, 4, 3, 1, CURRENT_TIMESTAMP), -- Seguridad - Eliminar
  (6, 1, 3, 1, 3, 1, CURRENT_TIMESTAMP), -- Sucursales - Listar
  (7, 1, 4, 1, 3, 1, CURRENT_TIMESTAMP), -- Clientes - Listar
  (8, 1, 5, 1, 3, 1, CURRENT_TIMESTAMP), -- Usuarios - Listar
  (9, 1, 5, 2 ,3 ,1 ,CURRENT_TIMESTAMP), -- Usuarios - Crear
  (10, 1, 5, 3, 3, 1, CURRENT_TIMESTAMP), -- Usuarios - Actualizar
  (11, 1, 5, 4, 3, 1, CURRENT_TIMESTAMP), -- Usuarios - Eliminar
  (12, 1, 6, 1, 3, 1, CURRENT_TIMESTAMP), -- Roles - Listar
  (13, 1, 6, 2 ,3 ,1 ,CURRENT_TIMESTAMP), -- Roles - Crear
  (14, 1, 6, 3 ,3 ,1 ,CURRENT_TIMESTAMP), -- Roles - Actualizar
  (15, 1, 6, 4 ,3 ,1 ,CURRENT_TIMESTAMP), -- Roles - Eliminar
  (16, 1, 7, 1 ,3 ,1 ,CURRENT_TIMESTAMP), -- Permisos - Listar
  (17, 1, 7, 2 ,3 ,1 ,CURRENT_TIMESTAMP), -- Permisos - Crear
  (18, 1, 7, 3 ,3 ,1 ,CURRENT_TIMESTAMP), -- Permisos - Actualizar
  (19, 1, 7, 4 ,3 ,1 ,CURRENT_TIMESTAMP), -- Permisos - Eliminar
  (20, 1, 8, 1 ,3 ,1 ,CURRENT_TIMESTAMP), -- Menús - Listar
  (21, 1, 8, 2 ,3 ,1 ,CURRENT_TIMESTAMP), -- Menús - Crear
  (22, 1, 8, 3 ,3 ,1 ,CURRENT_TIMESTAMP), -- Menús - Actualizar
  (23, 1, 8, 4 ,3 ,1 ,CURRENT_TIMESTAMP); -- Menús - Eliminar

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
