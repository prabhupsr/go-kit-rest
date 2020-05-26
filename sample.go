package users

import (
	json "encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
)

func add(x int, y int) int {
	return x + y
}

func addSimplified(x, y int) int {
	return x + y
}

func addAndMultiply(x, y int) (sum, product int) {
	sum = x + y
	product = x * y
	return
}

func pointers() {
	var a string = "abc"
	var aPointer = &a
	fmt.Println(aPointer)
	fmt.Println(*aPointer)
}

func structFn() Student {
	return CreateStudent(1, "prabhu")
}

func main() {
	fmt.Println("hello world with random ", rand.Intn(50))
	fmt.Println("hello world with random ", rand.Intn(50))
	fmt.Println("hello world with random ", math.Pi)
	fmt.Println("add ", add(10, 11))
	fmt.Println("add ", addSimplified(10, 11))

	sum, product := addAndMultiply(2, 5)
	fmt.Println(sum, product)

	/* var a, b = 5, false
	var a1, b1 = 5, false
	c1 := 4 */
	pointers()
	student := structFn()
	fmt.Println(student.Id)
	//fmt.Println(student.name)
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8090", nil)
}

type Person struct {
	Id   int
	Name string
}

type Employee struct {
	Name   string
	Age    int
	salary int
}

func handle(writer http.ResponseWriter, request *http.Request) {
	handleRequest(request.Body, request.Header.Get("myHeader"))
}

func handleRequest(body io.ReadCloser, myHeader string) (string, error) {
	if myHeader != "abc" {
		return "", errors.New("invalid Header")
	}
	person, err := decodePerson(body)
	if err != nil {
		return "", errors.New("invalid request body")
	}
	marshal, _ := json.Marshal(person)
	return string(marshal), nil

	/*if myheader := request.Header.Get("myHeader"); myheader == "abc" {
		person, err := decodePerson(request)
		if err != nil {
			http.Error(writer, "invalid input", http.StatusBadRequest)
		}
		writer.WriteHeader(200)
		marshal, _ := json.Marshal(person)
		fmt.Fprintf(writer, string(marshal))
	} else {
		http.Error(writer, "invalid input", http.StatusBadRequest)
	}*/
}

func decodePerson(body io.ReadCloser) (Person, error) {
	var person Person
	err := json.NewDecoder(body).Decode(&person)
	return person, err
}
