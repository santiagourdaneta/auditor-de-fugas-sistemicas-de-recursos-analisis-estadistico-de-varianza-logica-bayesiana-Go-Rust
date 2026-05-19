.PHONY: all test lint build run

all: lint test build

build:
	@echo "🦀 Compilando motor matemático en Rust..."
	@cd math-engine && cargo build --release
	@echo "🦫 Compilando CLI de Go..."
	@cd cli-gateway && go build -ldflags="-s -w" -o ../auditor cmd/main.go

test:
	@echo "🦫 Probando componentes Go..."
	@cd cli-gateway && go test -v -race ./...
	@echo "🦀 Probando componentes Rust y Estrés..."
	@cd math-engine && cargo test -- --nocapture
	@echo "🏁 Ejecutando prueba de integración de caja negra (E2E)..."
	@./test-e2e.sh

lint:
	@echo "🛡️  Limpiando caché de validación..."
	@golangci-lint cache clean
	@echo "🛡️  Validando calidad de código en el módulo de Go..."
	@cd cli-gateway && golangci-lint run --timeout 5m ./...