package main

import (
	"fmt"
	"log"

	"github.com/mickep76/dquery"
)

func main() {
	d := map[string]interface{}{
		"Cats": []string{
			"Leopard",
			"Lion",
			"Puma",
		},
		"Creatures": map[string]interface{}{
			"Evil": []string{
				"Dracula",
				"Dragon",
			},
			"Good": []string{
				"Unicorns",
				"Cats",
			},
		},
	}

	/*
		j, err := dquery.FilterJSON(".Cats", d)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(j))
	*/
	r2, err2 := dquery.Filter(".", d)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(r2)
	/*
		r4, err4 := dquery.FilterJSON(".", d)
		if err4 != nil {
			log.Fatal(err4)
		}
		fmt.Println(string(r4))

		r3, err3 := dquery.FilterJSON(".Cats.[5]", d)
		if err3 != nil {
			log.Fatal(err3)
		}
		fmt.Println(string(r3))
	*/
}
