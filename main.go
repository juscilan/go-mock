package main

import (
	"fmt"
	"log"

	"github.com/juscilan/godeep/deep"
)

func main() {
	d := new(deep.Deep)
	r := CallDeep(d)
	fmt.Println(r)
}

func CallDeep(dep deep.DeepInterface) string {
	d, error := dep.Exec()
	if error != nil {
		log.Fatal(error)
	}
	return d
}
