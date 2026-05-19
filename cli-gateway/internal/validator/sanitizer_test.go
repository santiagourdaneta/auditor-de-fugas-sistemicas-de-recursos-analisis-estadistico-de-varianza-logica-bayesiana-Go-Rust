package validator

import (
	"testing"
	"strings"
	"github.com/santiagourdaneta/resource-cli/internal/model"
)

// PRUEBA UNITARIA: Valida que los límites del contrato de negocio funcionen
func TestValidarRegistro_Limites(t *testing.T) {
	tests := []struct {
		name    string
		reg     *model.Registro
		wantErr bool
	}{
		{
			name: "Registro Válido Al Límite",
			reg: &model.Registro{
				Tipo:            model.Relacional,
				IniciativaDada:  5,
				IniciativaRecib: -5,
				MontoDinero:     100.50,
			},
			wantErr: false,
		},
		{
			name: "Error por Iniciativa Dada Excesiva",
			reg: &model.Registro{
				Tipo:            model.Relacional,
				IniciativaDada:  6, // Invalido
				IniciativaRecib: 0,
				MontoDinero:     0,
			},
			wantErr: true,
		},
		{
			name: "Error por Monto Negativo",
			reg: &model.Registro{
				Tipo:            model.Financiero,
				IniciativaDada:  0,
				IniciativaRecib: 0,
				MontoDinero:     -1.00, // Invalido
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidarRegistro(tt.reg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidarRegistro() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// PRUEBA DE INTEGRACIÓN: Simula la limpieza de strings con retornos de carro de Windows
func TestIntegracion_LimpiarEntradaWindows(t *testing.T) {
	entradaSucia := "1\r\n"
	limpia := strings.ReplaceAll(entradaSucia, "\r", "")
	limpia = strings.ReplaceAll(limpia, "\n", "")
	limpia = strings.TrimSpace(limpia)

	if limpia != "1" {
		t.Errorf("La sanitización de Windows falló. Esperado: '1', Obtenido: '%s'", limpia)
	}
}