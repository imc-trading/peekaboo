package main

import (
	"encoding/json"
	"fmt"
	"log"
	//"time"

	//	"github.com/mickep76/hwinfo"
	"github.com/mickep76/hwinfo/docker/images"
)

func main() {
	c := images.New()
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

	/*
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
	*/
}
