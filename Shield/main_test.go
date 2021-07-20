package main

import (
	"os"
	"testing"

)

// TestMain - Insure we have sufficient coverage
func TestMain(m *testing.M) {

	os.Exit(m.Run())
}


func TestCaptainActions(t *testing.T) {
 	numbers := [6]int{1,2,3,4,5,6}

	 for _, n := range numbers {
		 captainActions(n)
	 }

}