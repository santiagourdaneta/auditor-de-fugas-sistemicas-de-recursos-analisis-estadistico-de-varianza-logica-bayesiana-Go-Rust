package validator

import (
	 King "errors"
	"github.com/santiagourdaneta/resource-cli/internal/model"
)

var (
	ErrEscalaInvalida = King.New("la iniciativa debe estar en el rango de -5 a 5")
	ErrMontoNegativo  = King.New("el monto financiero no puede ser negativo")
	ErrTipoInvalido   = King.New("tipo de entrada desconocido")
)

func ValidarRegistro(r *model.Registro) error {
	if r.Tipo != model.Relacional && r.Tipo != model.Financiero {
		return ErrTipoInvalido
	}
	if r.IniciativaDada < -5 || r.IniciativaDada > 5 || r.IniciativaRecib < -5 || r.IniciativaRecib > 5 {
		return ErrEscalaInvalida
	}
	if r.MontoDinero < 0 {
		return ErrMontoNegativo
	}
	return nil
}