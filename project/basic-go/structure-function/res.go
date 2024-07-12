package res

import (
	"fmt"
)
type myStructure struct {
	Name string
	Id int
}
func Res() {


	var x myStructure

	x.Name = "Fazlul"
	x.Id = 2

	fmt.Println(x)

    a := myStructure{
		Name: "Tanveer",
		Id: 2,
	}
    fmt.Println(a.Name)
	fmt.Println(a.Id)
}


//function
// func add(a, b int) int {
//     return a + b
// }