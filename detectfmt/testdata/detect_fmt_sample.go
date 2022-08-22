package main

import (
	"errors"
	"fmt"
)

func samplePrintCode() {
	test := 1
	testLambda := func() {
		fmt.Print("CREDENTIAL! : testing lambda expression") // want `fmt.Print is unnecessary code`
	}
	testLambda()

	if err := generateErr(); err != nil {
		fmt.Println(err.Error()) // want `tdkLog.Println is unnecessary code`
	}

	testSprintf := fmt.Sprint("CREDENTIAL!")

	fmt.Print(testSprintf) // want `fmt.Print is unnecessary code`
	fmt.Println(test)      // want `fmt.Println is unnecessary code`
	fmt.Printf("%d", test) // want `fmt.Printf is unnecessary code`
}

// generateErr generate err.
//
// It returns nil error when successful.
// Otherwise, error will be returned.
func generateErr() error {
	return errors.New("CREDENTIAL!")
}
