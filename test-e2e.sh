#!/bin/bash
echo "🏁 Iniciando Prueba End-to-End (E2E)..."

# 1. Entrar al módulo de Go para que resuelva los paquetes internos con su go.mod
cd cli-gateway
go build -o ../auditor.exe ./cmd/main.go
cd ..

# 2. Verificar que el ejecutable realmente se creó antes de continuar
if [ ! -f ./auditor.exe ]; then
    echo "❌ FAILED: Error crítico de compilación en el módulo de Go."
    exit 1
fi

# 3. Inyectar datos simulados asegurando saltos de línea estándar (\n)
# Opción 1 -> Tipo 1 -> Dada 5 -> Recibida -5 -> Monto 0 -> Opción 2 (Salir)
printf "1\n1\n5\n-5\n0\n2\n" | ./auditor.exe > e2e_output.txt 2>&1

# 4. Validar el token de éxito ignorando diferencias de codificación
if grep -a -q "Estructura Sanitizada:" e2e_output.txt; then
    echo "✅ PASSED: El flujo E2E completó la sanitización y validación de forma nativa."
    rm -f e2e_output.txt auditor.exe
    exit 0
else
    echo "❌ FAILED: El validador de caja negra no encontró el payload sanitizado."
    cat e2e_output.txt
    rm -f e2e_output.txt auditor.exe
    exit 1
fi