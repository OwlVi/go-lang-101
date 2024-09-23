package array_101

import "fmt"

func Array() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}

	for i := range array {
		fmt.Println(i)
	}

}
