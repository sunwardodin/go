package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	engine
	make  string
	model string
	doors int
	color string
}

func main() {

	v1 := vehicle{
		engine: engine{
			electric: false,
		},
		make:  "Toyota",
		model: "Corolla",
		doors: 4,
		color: "Grey",
	}

	v2 := vehicle{
		engine: engine{
			electric: true,
		},
		make:  "Toyota",
		model: "Camry",
		doors: 2,
		color: "Green",
	}

	fmt.Println(v1)
	fmt.Println(v1.engine)
	fmt.Println(v2)
	fmt.Println(v2.color)

}
