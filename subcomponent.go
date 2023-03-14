package golevel7

// SubComponent is an HL7 subcomponent
type SubComponent struct {
	Value []rune
}

func (sc *SubComponent) GetValue() string {
	return string(sc.Value)
}
