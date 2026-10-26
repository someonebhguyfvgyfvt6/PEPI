# PEPI Backend

API de PEPI (portafolios para freelancers y empresas).
Stack: Go, Fiber v2, PostgreSQL (Supabase), Clean Architecture.

## Requisitos

- Go 1.22 o superior
- Una base PostgreSQL (por ejemplo, un proyecto de Supabase)

## Cómo correrlo

```bash
cd pepi-backend
cp .env.example .env        # luego completa DATABASE_URL con tu cadena de Supabase
go mod download
go run ./cmd/api
```

Comprobar que funciona:

```bash
curl http://localhost:8080/health
# {"status":"ok"}
```

Para detenerlo, pulsa `Ctrl+C`: el servidor termina las peticiones en curso y cierra la conexión a la base.

## Compilar

```bash
go build -o bin/api ./cmd/api
```

## Estructura

```
cmd/api/                 punto de entrada (main)
config/                  carga de variables de entorno
internal/domain/         entidades e interfaces de repositorio (no depende de nada)
internal/usecase/        lógica de negocio (solo depende de domain)
internal/repository/     implementaciones de repositorios con PostgreSQL
internal/delivery/http/  router, handlers y middlewares de Fiber
migrations/              migraciones SQL (golang-migrate)
pkg/                     utilidades reutilizables (conexión a la base)
```
