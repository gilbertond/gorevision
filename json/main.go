package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//Use case: Read json from stream, delete a field, and encode back to json
//can also be used in cases such as reading and writing to HTTP connections, WebSockets, or files
func decodeAndEncodeStreamsOfJsonData() {
	dec:= json.NewDecoder(os.Stdin)
	enc:= json.NewEncoder(os.Stdin)

	for{
		var v map[string]interface{}
		if err := dec.Decode(&v); err != nil {
			log.Println(err)
			return
		}
		for k := range v {
			if k != "Name" {
				delete(v, k)
			}
		}
		if err := enc.Encode(&v); err != nil {
			log.Println(err)
		}
	}
}

func interfaceStuff() {
	var i interface{}
	i = 2
	i = 3.4
	i = "a string"
	i = 12.11

	r := i.(float64)
	fmt.Printf("square: %f\n", r)

	switch v := i.(type) {
	case string:
		fmt.Printf("a string: %v", v)
	case float64:
		fmt.Printf("a float64: %v", v)
	case float32:
		fmt.Printf("a float32: %v", v)
	case int:
		fmt.Printf("an int: %v", v)
	default:
		fmt.Printf("default: %v", v)
	}
}

func unMarshallWithUnknownType() {

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`) //type of json used to marshal is unknown!!

	var f interface{}
	err := json.Unmarshal(b, &f)
	if err != nil {
		fmt.Errorf("Error marshalling, %value", err)
	}

	//	f will now be a map whose keys are strings and values as empty interface-values
	//Like below------->>
	//f = map[string]interface{}{
	//	"Name": "Wednesday",
	//	"Age":  6,
	//	"Parents": []interface{}{
	//		"Gomez",
	//		"Morticia",
	//	},
	//}

	//	Access values
	m := f.(map[string]interface{})
	for key, value := range m {
		switch valueType := value.(type) {
		case string:
			fmt.Println(key, "is string", valueType)
		case float64:
			fmt.Println(key, "is float64", valueType)
		case []interface{}:
			fmt.Println(key, "is an array:")
			for i, u := range valueType {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(key, "is of a type I don't know how to handle")
		}
	}
}

//https://blog.golang.org/json-and-go
//By default, only fields in the json are un-marshaled. This behaviour can be overridden in cases when you want the whole json
//Can be achieved using interface. e.g. interface{}. empty interface with 0 methods. every Go type implements this interface
func main() {
	decodeAndEncodeStreamsOfJsonData()
	fmt.Println()
	fmt.Println()
	unMarshallWithUnknownType()
	fmt.Println()

	interfaceStuff()
	fmt.Println()

	msg := Message{
		Id:          1,
		Message:     "some message",
		Description: "some description",
		Sender:      Person{Id: 1, Fname: "gil", Lname: "nd"},
		Receivers: []Person{
			{Id: 1, Fname: "gil", Lname: "nd"},
			{Id: 1, Fname: "gil", Lname: "nd"},
			{Id: 1, Fname: "gil", Lname: "nd"},
		},
	}

	jsn, error := json.Marshal(msg)
	if error != nil {
		fmt.Errorf("json can't be read because of %v", error)
	}

	fmt.Println(string(jsn))

	var message Message
	json.Unmarshal(jsn, &message)

	fmt.Println(message)
}

//Fields gotta be in caps to be marshaled. Lower case are private
type Message struct {
	Id          int64
	Message     string
	Description string
	Sender      Person
	Receivers   []Person
	SomePerson  *Person //Unmarshalling without this would return a run for this field instead of being ignored!!
}

type Person struct {
	Id    int64
	Fname string
	Lname string
}
