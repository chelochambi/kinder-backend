-- =======================
-- 🧩 Tabla: tipo_estado
-- =======================
CREATE TABLE tipo_estado (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL UNIQUE,
    descripcion TEXT,
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    actualizado_en TIMESTAMP,
    codigo VARCHAR(50) UNIQUE,
    creado_por INTEGER,
    actualizado_por INTEGER
);

-- =======================
-- 🧩 Tabla: estados
-- =======================
CREATE TABLE estados (
    id SERIAL PRIMARY KEY,
    tipo_estado_id INTEGER REFERENCES tipo_estado(id),
    nombre VARCHAR(100) NOT NULL,
    codigo VARCHAR(50),
    descripcion TEXT,
    activo BOOLEAN DEFAULT TRUE,
    orden INTEGER DEFAULT 0,
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    actualizado_en TIMESTAMP,
    creado_por INTEGER,
    actualizado_por INTEGER,
    UNIQUE (tipo_estado_id, nombre)
);

-- =======================
-- 👤 Tabla: usuarios
-- =======================
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL DEFAULT '12345678',
    nombres VARCHAR(100),
    primer_apellido VARCHAR(100),
    segundo_apellido VARCHAR(100),
    estado_id INTEGER NOT NULL DEFAULT 1 REFERENCES estados(id),
    ultimo_login TIMESTAMP,
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    telefono VARCHAR(20),
    foto_url TEXT DEFAULT '/src/assets/fotos_usuarios/sin_foto.png',
    actualizado_en TIMESTAMP,
    creado_por INTEGER,
    actualizado_por INTEGER
);

-- =======================
-- 🔐 Tabla: roles
-- =======================
CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) UNIQUE NOT NULL,
    codigo VARCHAR(10) UNIQUE NOT NULL,
    descripcion TEXT,
    estado_id INTEGER REFERENCES estados(id),
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    actualizado_en TIMESTAMP,
    creado_por INTEGER,
    actualizado_por INTEGER
);

-- =======================
-- 📋 Tabla: menus
-- =======================
CREATE TABLE menus (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    icono VARCHAR(50),
    ruta VARCHAR(255),
    orden INTEGER DEFAULT 0,
    tipo VARCHAR(50),
    mostrar BOOLEAN DEFAULT TRUE,
    padre_id INTEGER REFERENCES menus(id),
    estado_id INTEGER REFERENCES estados(id),
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    actualizado_en TIMESTAMP,
    creado_por INTEGER,
    actualizado_por INTEGER
);

-- =======================
-- 🔑 Tabla: permisos
-- LIS-listar
-- CRE-crear
-- ACT-actualizar
-- ELI-eliminar
-- =======================
CREATE TABLE permisos (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    codigo VARCHAR(100) UNIQUE,
    descripcion TEXT,
    estado_id INTEGER REFERENCES estados(id),
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    actualizado_en TIMESTAMP,
    creado_por INTEGER,
    actualizado_por INTEGER
);

-- =======================
-- 🔑 Tabla: rol_menu
-- =======================
CREATE TABLE rol_menu_permiso (
    id SERIAL PRIMARY KEY,
    rol_id INTEGER REFERENCES roles(id),
    menu_id INTEGER REFERENCES menus(id),
    permiso_id INTEGER REFERENCES permisos(id),
    estado_id INTEGER REFERENCES estados(id),
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    actualizado_en TIMESTAMP,
    creado_por INTEGER,
    actualizado_por INTEGER,
    UNIQUE(rol_id, menu_id, permiso_id)
);
-- =======================
-- 📎 Tabla: usuario_rol
-- =======================
CREATE TABLE usuario_rol (
    id SERIAL PRIMARY KEY,
    usuario_id INTEGER REFERENCES usuarios(id),
    rol_id INTEGER REFERENCES roles(id),
    estado_id INTEGER REFERENCES estados(id),
    vigente_desde DATE,
    vigente_hasta DATE,
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    creado_por INTEGER,
    actualizado_en TIMESTAMP,
    actualizado_por INTEGER
);

-- =======================
-- 🏢 Tabla: sucursales
-- =======================
CREATE TABLE sucursales (
    id SERIAL PRIMARY KEY,
    nombre VARCHAR(100) NOT NULL,
    codigo VARCHAR(50) UNIQUE,
    descripcion TEXT,
    direccion TEXT,
    telefono VARCHAR(20),
    email VARCHAR(100),
    estado_id INTEGER REFERENCES estados(id),
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    actualizado_en TIMESTAMP,
    creado_por INTEGER,
    actualizado_por INTEGER
);

-- =======================
-- 🏢 Tabla: usuario_sucursal
-- =======================
CREATE TABLE usuario_sucursal (
    id SERIAL PRIMARY KEY,
    usuario_id INTEGER REFERENCES usuarios(id),
    sucursal_id INTEGER REFERENCES sucursales(id),
    estado_id INTEGER REFERENCES estados(id),
    creado_en TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    actualizado_en TIMESTAMP,
    creado_por INTEGER,
    actualizado_por INTEGER,
    UNIQUE(usuario_id, sucursal_id)
);

