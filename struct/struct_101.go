package struct_101

// struct like c
type name_struct struct {
	x int
	y string
}

func Set_struct(x int, y string) name_struct {
	struct_ := name_struct{
		x: x,
		y: y,
	}
	return struct_
}

// func receiver
func (p name_struct) Get_x() int {
	return p.x
}

func (p name_struct) Get_y() string {
	return p.y
}
