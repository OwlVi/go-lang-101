// package
package main

// import
import (
	printer "github.com/OwlVi/go-lang-101/interface"
)

func Print(pt printer.Printer) {
	pt.Print()
}

func main() {
	inkjet := printer.NewInkjet("Inkjet")
	laser := printer.NewLaser("Laser")

	Print(inkjet)
	Print(laser)
}
