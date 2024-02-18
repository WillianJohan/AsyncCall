# AsyncCall
The AsyncCall package is designed to simplify asynchronous function calls in Go, inspired by the functionality of Promise.all in Node.js. It provides an easy way to execute multiple functions concurrently and handle their results efficiently.


To install, simply execute the following command in your project terminal:
```
go get github.com/WillianJohan/AsyncCall
```

### Example of usage

```go
package main

import (
	"fmt"
	"time"

	AsyncCall "github.com/WillianJohan/AsyncCall"
)

// Definition of a sample struct
type Person struct {
	ID    int
	Name  string
	Age   int
	Street string
}

func ExampleGetPeople() {
	var John Person
	var Alice Person
	var Bob Person

	ac := AsyncCall.New()
	ac.AddAsyncCall(func() { John = GetPerson(1, "John", 30, "Main St") })
	ac.AddAsyncCall(func() { Alice = GetPerson(2, "Alice", 25, "Park Ave") })
	ac.AddAsyncCall(func() { Bob = GetPerson(3, "Bob", 35, "Elm St") })

	ac.Execute()

	fmt.Printf("\n%+v\n", John)
	fmt.Printf("\n%+v\n", Alice)
	fmt.Printf("\n%+v\n", Bob)
}

func ExampleUsingPromiseAll() {
	var John Person
	var Alice Person
	var Bob Person

	AsyncCall.ExecutePromiseAll([]func(){
		func() { John = GetPerson(1, "John", 30, "Main St") },
		func() { Alice = GetPerson(2, "Alice", 25, "Park Ave") },
		func() { Bob = GetPerson(3, "Bob", 35, "Elm St") },
	})

	fmt.Printf("\n%+v\n", John)
	fmt.Printf("\n%+v\n", Alice)
	fmt.Printf("\n%+v\n", Bob)
}

// Function to simulate fetching a person
func GetPerson(id int, name string, age int, street string) Person {
	time.Sleep(5 * time.Second)
	return Person{
		ID:    id,
		Name:  name,
		Age:   age,
		Street: street,
	}
}

func main() {
	ExampleGetPeople()
	ExampleUsingPromiseAll()
}


```
This package facilitates concurrent function execution, allowing users to improve the efficiency of their Go programs when dealing with asynchronous tasks.
