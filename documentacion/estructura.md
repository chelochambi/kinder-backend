/src
│
├── components/              # Componentes reutilizables (Botones, Modales, Tabs, Cards, etc.)
├── context/                 # AuthContext, UserContext, ThemeContext...
├── hooks/                   # Hooks personalizados (useFetch, useForm, usePermisos)
├── layouts/                 # Layouts principales (DashboardLayout, AuthLayout)
├── pages/                  
│   ├── auth/                # Login, Registro (si aplica)
│   ├── dashboard/           # Vista general (widgets)
│   ├── usuarios/            # Lista, creación, edición de usuarios
│   ├── estudiantes/         # Gestión de estudiantes
│   ├── docentes/            # Gestión de docentes
│   ├── grados/              # Cursos, paralelos, niveles
│   ├── materias/            # Materias disponibles
│   ├── asistencia/          # Registro de asistencia
│   ├── evaluaciones/        # Registro de notas
│   ├── pagos/               # Finanzas, facturación
│   └── reportes/            # Reportes dinámicos o generados
│
├── routes/                  # Definición de rutas protegidas y públicas
├── services/                # Llamadas a la API (usuariosService, authService, etc.)
├── utils/                   # Funciones utilitarias
├── App.jsx
└── main.jsx
