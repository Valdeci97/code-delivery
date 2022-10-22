package main

import (
	"fmt"

	"github.com/Valdeci97/full-cycle-simulator/application/route"
)

func main() {
	r := route.Route{
		ID:       "1",
		ClientID: "1",
	}
	r.LoadPositions()
	stringjson, _ := r.ExportJsonPositions()
	fmt.Println(stringjson[0])
}
