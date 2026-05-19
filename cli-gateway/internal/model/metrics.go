package model

type TipoEntrada string

const (
	Relacional TipoEntrada = "RELACIONAL"
	Financiero TipoEntrada = "FINANCIERO"
)

type Registro struct {
	Tipo            TipoEntrada `json:"tipo"`
	IniciativaDada  int         `json:"iniciativa_dada"`  // Escala -5 a 5
	IniciativaRecib int         `json:"iniciativa_recib"` // Escala -5 a 5
	MontoDinero     float64     `json:"monto_dinero"`     // Precisión fija para dinero
	Timestamp       int64       `json:"timestamp"`
}