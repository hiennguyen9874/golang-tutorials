package main

import (
	"encoding/json"
	"fmt"
)

type Bird struct {
	Species     string
	Description string
}

type Dimensions struct {
	Height int
	Width  int
}

type Bird2 struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

type Bird3 struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

// The same json tags will be used to encode data into JSON
type Bird4 struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

type Bird5 struct {
	Species string `json:"birdType"`
	// we can set the "omitempty" property as part of the JSON tag
	Description string `json:"what it does,omitempty"`
}

func main() {
	// 1. Structured Data
	birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
	var bird Bird
	json.Unmarshal([]byte(birdJson), &bird)
	fmt.Printf("Species: %s, Description: %s\n", bird.Species, bird.Description)

	// 1.1 Arrays
	birdJson2 := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`
	var birds []Bird
	json.Unmarshal([]byte(birdJson2), &birds)
	fmt.Printf("Birds : %+v\n", birds)

	// 1.2 Nested Objects
	birdJson3 := `{"species":"pigeon","description":"likes to perch on rocks", "dimensions":{"height":24,"width":10}}`
	var birds2 Bird2
	json.Unmarshal([]byte(birdJson3), &birds2)
	fmt.Println(birds2)

	numberJson := "3"
	floatJson := "3.1412"
	stringJson := `"bird"`

	var n int
	var pi float64
	var str string

	json.Unmarshal([]byte(numberJson), &n)
	fmt.Println(n)

	json.Unmarshal([]byte(floatJson), &pi)
	fmt.Println(pi)

	json.Unmarshal([]byte(stringJson), &str)
	fmt.Println(str)

	// 1.3 Custom Field Names
	birdJson4 := `{"birdType": "pigeon","what it does": "likes to perch on rocks"}`
	var bird2 Bird3
	json.Unmarshal([]byte(birdJson4), &bird2)
	fmt.Println(bird2)

	// 2. Unstructured Data
	birdJson5 := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"},"animals":"none"}`
	var result map[string]any
	json.Unmarshal([]byte(birdJson5), &result)

	// The object stored in the "birds" key is also stored as
	// a map[string]any type, and its type is asserted from
	// the `any` type
	birds3 := result["birds"].(map[string]any)

	for key, value := range birds3 {
		// Each value is an `any` type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}

	// 2.1 Validating JSON Data
	birdJson6 := `{"birds":{"pigeon":"likes to perch on rocks","eagle":"bird of prey"`
	if !json.Valid([]byte(birdJson6)) {
		// handle the error here
		fmt.Println("invalid JSON string:", birdJson6)
		// return
	}

	pigeon := &Bird4{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	// we can use the json.Marhal function to
	// encode the pigeon variable to a JSON string
	data, _ := json.Marshal(pigeon)
	// data is the JSON string represented as bytes
	// the second parameter here is the error, which we
	// are ignoring for now, but which you should ideally handle
	// in production grade code

	// to print the data, we can typecast it to a string
	fmt.Println(string(data))

	// 2.2 Ignoring Empty Fields
	pigeon2 := &Bird{
		Species: "Pigeon",
	}

	data2, _ := json.Marshal(pigeon2)

	fmt.Println(string(data2))

	// 2.3 Marshaling Slices
	pigeon3 := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	// Now we pass a slice of two pigeons
	data3, _ := json.Marshal([]*Bird{pigeon3, pigeon3})
	fmt.Println(string(data3))

	// 2.4 Marshaling Maps
	// The keys need to be strings, the values can be
	// any serializable value
	birdData4 := map[string]any{
		"birdSounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squak",
		},
		"total birds": 2,
	}

	// JSON encoding is done the same way as before
	data4, _ := json.Marshal(birdData4)
	fmt.Println(string(data4))

	fmt.Println("Done!")
}
