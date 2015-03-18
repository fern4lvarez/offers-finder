package main

import (
	"fmt"
)

func ExampleJsonResponseStringOK() {
	fmt.Println(JsonResponse{"foo": "bar"})
	// Output: {"foo":"bar"}
}

func ExampleJsonResponseEmpty() {
	fmt.Println(JsonResponse{})
	// Output: {}
}
