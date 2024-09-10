package day3

type Schematic struct {
	MaxX       int
	MaxY       int
	References []SchematicReference
}

func NewSchematic(lines []string) Schematic {
	return Schematic{}
}
