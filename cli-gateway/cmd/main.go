package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/santiagourdaneta/resource-cli/internal/model"
	"github.com/santiagourdaneta/resource-cli/internal/validator"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("🛡️  AUDITOR DE FUGA DE RECURSOS V1.2 (2026) 🛡️")
	fmt.Println("==================================================")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n[1] Registrar Nuevo Evento\n[2] Salir del Sistema\n\nSelecciona una opción -> ")
		opcion, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		opcion = limpiarEntrada(opcion)

		if opcion == "2" {
			fmt.Println("\n🔒 Cerrando auditoría segura. Tus datos locales están a salvo.")
			break
		}

		if opcion == "1" {
			procesarFormulario(reader)
			continue
		}

		fmt.Println("❌ Opción de menú inválida. Elige 1 o 2.")
	}
}

func procesarFormulario(reader *bufio.Reader) {
	reg := &model.Registro{Timestamp: time.Now().Unix()}
	fmt.Println("\n--- NUEVO REGISTRO DE RECURSOS ---")

	reg.Tipo = capturarTipo(reader)
	reg.IniciativaDada = capturarRango(reader, "Tu inversión de energía/tiempo (-5 a 5): ")
	reg.IniciativaRecib = capturarRango(reader, "Retorno recibido de la otra parte (-5 a 5): ")
	reg.MontoDinero = capturarMonto(reader)

	if err := validator.ValidarRegistro(reg); err != nil {
		fmt.Printf("\n❌ Error crítico de consistencia: %v. Reiniciando formulario.\n", err)
		return
	}

	payload, _ := json.Marshal(reg)
	fmt.Println("\n--------------------------------------------------")
	fmt.Printf("🔒 Estructura Sanitizada: %s\n", string(payload))
	fmt.Println("✅ Datos validados en memoria local listos para analítica de Rust.")
	fmt.Println("--------------------------------------------------")
}

func capturarTipo(reader *bufio.Reader) model.TipoEntrada {
	for {
		fmt.Print("Selecciona Tipo -> [1] RELACIONAL | [2] FINANCIERO: ")
		tipoRaw, _ := reader.ReadString('\n')
		tipoOpt := limpiarEntrada(tipoRaw)

		if tipoOpt == "1" {
			return model.Relacional
		}
		if tipoOpt == "2" {
			return model.Financiero
		}
		fmt.Println("❌ Opción inválida. Presiona 1 o 2.")
	}
}

func capturarRango(reader *bufio.Reader, msg string) int {
	for {
		fmt.Print(msg)
		dadaRaw, _ := reader.ReadString('\n')
		val, err := strconv.Atoi(limpiarEntrada(dadaRaw))
		if err != nil {
			fmt.Println("❌ Error: Debes ingresar un número entero válido.")
			continue
		}
		if val < -5 || val > 5 {
			fmt.Println("❌ Error: El valor debe estar estrictamente entre -5 y 5.")
			continue
		}
		return val
	}
}

func capturarMonto(reader *bufio.Reader) float64 {
	for {
		fmt.Print("Monto económico gastado/prestado (0 si no aplica): ")
		montoRaw, _ := reader.ReadString('\n')
		monto, err := strconv.ParseFloat(limpiarEntrada(montoRaw), 64)
		if err != nil {
			fmt.Println("❌ Error: Debes ingresar un monto numérico válido.")
			continue
		}
		if monto < 0 {
			fmt.Println("❌ Error: El monto no puede ser negativo.")
			continue
		}
		return monto
	}
}

func limpiarEntrada(s string) string {
	s = strings.ReplaceAll(s, "\r", "")
	s = strings.ReplaceAll(s, "\n", "")
	return strings.TrimSpace(s)
}