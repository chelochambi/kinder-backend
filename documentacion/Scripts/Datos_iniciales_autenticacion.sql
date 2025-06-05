-- =====================================
-- 🧩 Datos de inicio para tipo_estado
-- =====================================
INSERT INTO tipo_estado (id, nombre, descripcion, codigo) VALUES
  (1, 'General', 'Estados generales para uso común', 'GEN'),
  (2, 'Usuario', 'Estados de los usuarios', 'USR'),
  (3, 'Rol', 'Estados para roles', 'ROL'),
  (4, 'Permiso', 'Estados para permisos', 'PER'),
  (5, 'Menú', 'Estados para los menús', 'MEN'),
  (6, 'Sucursal', 'Estados de las sucursales', 'SUC');

-- ===============================
-- 🧩 Datos de inicio para estados
-- ===============================
INSERT INTO estados (tipo_estado_id, nombre, codigo, descripcion, orden) VALUES
  (1, 'Activo', 'ACT', 'Elemento activo', 1),
  (2, 'Inactivo', 'INA', 'Elemento inactivo', 2),
  (3, 'Vigente', 'VIG', 'Usuario vigente', 1),
  (4, 'Bloqueado', 'BLQ', 'Usuario bloqueado', 2),
  (5, 'Bloqueado Aut.', 'BLQA', 'Usuario bloqueado de forma automática', 3);

-- ================================
-- 🏢 Datos de inicio para sucursales
-- ================================
INSERT INTO sucursales (nombre, codigo, descripcion, direccion, telefono, email, estado_id) VALUES
  ('Sucursal Central', 'CENTRAL', 'Sede principal del centro educativo', 'Av. Principal 123', '12345678', 'central@sistema.edu', 1),
  ('Sucursal Norte', 'NORTE', 'Ubicada en la zona norte', 'Calle 10 y Av. Norte', '87654321', 'norte@sistema.edu', 1);

-- ============================
-- 🔐 Datos de inicio para roles
-- ============================
INSERT INTO roles (nombre, codigo, descripcion, estado_id) VALUES
  ('Administrador', 'ADMIN', 'Rol con acceso total al sistema', 1),
  ('Director', 'DIR', 'Rol con permisos de gestión escolar', 1),
  ('Docente', 'DOC', 'Rol asignado a maestros', 1);

-- ==============================
-- 📋 Datos de inicio para menús
-- ==============================
INSERT INTO menus (nombre, icono, ruta, orden, tipo, mostrar, estado_id) VALUES
  ('Panel', 'fas fa-tachometer-alt', '/dashboard', 1, 'link', TRUE, 1),
  ('Gestión Usuarios', 'fas fa-users-cog', '/usuarios', 2, 'link', TRUE, 1),
  ('Sucursales', 'fas fa-building', '/sucursales', 3, 'link', TRUE, 1);

-- ================================
-- 🔑 Datos de inicio para permisos
-- ================================
INSERT INTO permisos (nombre, codigo, descripcion, menu_id, accion, estado_id) VALUES
  ('Ver panel', 'PANEL_VIEW', 'Acceso al panel principal', 1, 'ver', 1),
  ('Gestionar usuarios', 'USUARIOS_MANAGE', 'Crear/editar/eliminar usuarios', 2, 'gestion', 1),
  ('Ver sucursales', 'SUCURSAL_VIEW', 'Ver listado de sucursales', 3, 'ver', 1),
  ('Gestionar sucursales', 'SUCURSAL_MANAGE', 'Agregar/modificar sucursales', 3, 'gestion', 1);

-- ===============================
-- 👤 Datos de inicio para usuarios
-- ===============================
-- ⚠️ La contraseña es 'admin123' hasheada (ejemplo: usando bcrypt)
INSERT INTO usuarios (username, email, password_hash, nombres, primer_apellido, segundo_apellido, estado_id, telefono) VALUES
  ('admin', 'admin@sistema.edu', '$2a$10$ShA2eCyjautV9HjkTx2.auUimgetTZw56PtUkolwOvJEnaYqgELiC', 'Super', 'Admin', 'Sistema', 3, '70000000');

-- ====================================
-- 📎 Datos de inicio para usuario_rol
-- ====================================
INSERT INTO usuario_rol (usuario_id, rol_id, estado_id, vigente_desde) VALUES
  (1, 1, 1, CURRENT_DATE);

-- =====================================
-- 🧩 Datos de inicio para rol_permiso
-- =====================================
INSERT INTO rol_permiso (rol_id, permiso_id, estado_id, vigente_desde) VALUES
  (1, 1, 1, CURRENT_DATE),
  (1, 2, 1, CURRENT_DATE),
  (1, 3, 1, CURRENT_DATE),
  (1, 4, 1, CURRENT_DATE);

-- ===========================================
-- 🧩 Datos de inicio para usuario_sucursal
-- ===========================================
INSERT INTO usuario_sucursal (usuario_id, sucursal_id, estado_id) VALUES
  (1, 1, 1),
  (1, 2, 1);
