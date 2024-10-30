package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Email string
}

// Here we will try to take a simple Person struct, convert it to JSON, and then convert it back to the struct.
func main() {
	person := &Person{
		Name:  "John Doe",
		Age:   30,
		Email: "john.doe@example.com",
	}

	personJSON, err := getJson(*person)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("JSON:", personJSON)

	personStruct, err := getStruct(personJSON)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Struct: ", personStruct)

}

// getJson takes a Person struct as input and returns its JSON string representation.
// If an error occurs during the marshaling process, it returns the error message and the error itself.
//
// Parameters:
//   - payload: Person struct to be converted to JSON.
//
// Returns:
//   - string: JSON string representation of the input Person struct.
//   - error: Error encountered during the marshaling process, if any.
func getJson(payload Person) (string, error) {
	// JSON string
	jsonString, err := json.Marshal(payload)

	if err != nil {
		return err.Error(), err
	}
	return string(jsonString), nil
}

// getStruct takes a JSON payload as a string and unmarshals it into a Person struct.
// It returns the populated Person struct and an error if the unmarshalling fails.
//
// Parameters:
//   - payload: A string containing the JSON data.
//
// Returns:
//   - Person: The unmarshalled Person struct.
//   - error: An error if the JSON unmarshalling fails.
func getStruct(payload string) (Person, error) {
	person := Person{}
	err := json.Unmarshal([]byte(payload), &person)

	if err != nil {
		return person, err
	}
	return person, nil
}
