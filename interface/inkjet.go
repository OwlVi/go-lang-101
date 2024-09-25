package interface_101

import "fmt"

type InkjetPrinter struct {
	Name string
}

// class implements format
func NewInkjet(name string) Printer {
	return &InkjetPrinter{
		Name: name,
	}
}

func (i *InkjetPrinter) Print() {
	// process
	fmt.Println("Prepare for Inkjet")

	// result
	fmt.Println("PRINT")
}
