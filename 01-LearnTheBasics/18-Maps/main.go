package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	if m == nil {
		fmt.Println("NIL")
	}

	m = make(map[string]Vertex)

	if m == nil {
		fmt.Println("NIL")
	}

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// initialization map
	var timeZone = map[string]int{
		"UTC": 0 * 60 * 60,
		"EST": -5 * 60 * 60,
		"CST": -6 * 60 * 60,
		"MST": -7 * 60 * 60,
		"PST": -8 * 60 * 60,
	}
	fmt.Println(timeZone)

	// Assigning and fetching map values
	offset := timeZone["EST"]
	fmt.Println(offset)

	offset2 := timeZone["EST2"]
	fmt.Println(offset2)

	_, present := timeZone["Asia/Ho_Chi_Minh"]
	fmt.Println(present)

	// Updating and adding map element
	timeZone["Asia"] = -7 * 60 * 60

	// Delete key
	delete(timeZone, "PDT") // Now on Standard Time
}
