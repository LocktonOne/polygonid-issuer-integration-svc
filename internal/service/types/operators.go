package types

type Operators int64

const (
	NOOP Operators = iota // No operation, skip query verification in circuit
	EQ                    // equal
	LT                    // less than
	GT                    // greater than
	IN                    // in
	NIN                   // not in
	NE                    // not equal

)

var operators = map[string]Operators{
	"NOOP": NOOP, // No operation, skip query verification in circuit
	"EQ":   EQ,   // equal
	"LT":   LT,   // less than
	"GT":   GT,   // greater than
	"IN":   IN,   // in
	"NIN":  NIN,  // not in
	"NE":   NE,   // not equal
}

func StringOperator(s string) Operators {
	return operators[s]
}
