package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mickep76/hwinfo"
	//	"github.com/mickep76/hwinfo/cpu"
)

func main() {
	/*
		c := cpu.New()
		if err := c.Update(); err != nil {
			log.Fatal(err.Error())
		}

		data, err := json.MarshalIndent(c.GetData(), "", "    ")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(data))

		cache, err := json.MarshalIndent(c.GetCache(), "", "    ")
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(cache))

		if err := c.Update(); err != nil {
			log.Fatal(err.Error())
		}

		cache3, err4 := json.MarshalIndent(c.GetCache(), "", "    ")
		if err4 != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(string(cache3))

		c.SetTimeout(5)
		time.Sleep(10 * time.Second)

		if err := c.Update(); err != nil {
			log.Fatal(err.Error())
		}

		data2, err2 := json.MarshalIndent(c.GetData(), "", "    ")
		if err2 != nil {
			fmt.Println(err2.Error())
		}
		fmt.Println(string(data2))

		cache2, err3 := json.MarshalIndent(c.GetCache(), "", "    ")
		if err3 != nil {
			fmt.Println(err3.Error())
		}
		fmt.Println(string(cache2))
	*/

	i := hwinfo.New()

	// Update
	fmt.Println("---> Update 1")
	if err := i.Update(); err != nil {
		log.Fatal(err.Error())
	}

	// Data
	fmt.Println("---> Data 1")
	data, err := json.MarshalIndent(i.GetData(), "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))

	// Cache
	fmt.Println("---> Cache 1")
	cache, err := json.MarshalIndent(i.GetCache(), "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(cache))

	// Update
	fmt.Println("---> Update 2")
	if err := i.Update(); err != nil {
		log.Fatal(err.Error())
	}

	// Cache
	fmt.Println("---> Cache 2")
	cache3, err := json.MarshalIndent(i.GetCache(), "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(cache3))

	// Sleep
	fmt.Println("---> Sleep 10 sec")
	i.GetCPU().SetTimeout(5)
	time.Sleep(10 * time.Second)

	// Update
	fmt.Println("---> Update 3")
	if err := i.Update(); err != nil {
		log.Fatal(err.Error())
	}

	// Cache
	fmt.Println("---> Cache 3")
	cache2, err := json.MarshalIndent(i.GetCache(), "", "    ")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(cache2))
}
