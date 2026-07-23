package main 

import "testing"

// Test functions must be followed 
func TestAdd(t *testing.T){
	got := Add(2 , 3)
	want := 5 
	if got !=  want {
		t. , got , amdErrorf("Add(2 , 3) = %d" , got ,want)
	}
}

