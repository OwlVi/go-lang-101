package interface_101

import "fmt"

type LaserPrinter struct {
	Name string
}

func NewLaser(name string) Printer {
	return &LaserPrinter{
		Name: name,
	}
}

func (i *LaserPrinter) Print() {
	// process
	fmt.Println("Prepare for Laser")

	// result
	fmt.Println("PRINT")
}
