package main

import "fmt"

func main() {

	/*var names map[string]int

	names = make(map[string]int, 0)

	names["mustafa"] = 112
	names["murat"] = 1

	fmt.Println(names)
	*/

	names := map[string]int{

		"mustafa": 1,
		"Ahmed ":  2,
		"coşkun":  3,
	}
	delete(names, "mustafa")
	fmt.Println(names)

}
