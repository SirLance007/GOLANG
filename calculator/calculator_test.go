package calculator

import "testing"

func TestMultiply(t *testing.T){
	// 1. Inputs aur Expected Output decide karo
	a, b := 3, 4
	expected := 12

	// 2. Real function call karo
	result := Multiply(a, b)

	// 3. Match check karo
	if result != expected {
		t.Errorf("Multiply(%d, %d) = %d; want %d", a, b, result, expected)
	}
}

func TestUserProfile(t *testing.T){
	input := "Prankur"
	want := "Prankur"

	got := UserProfile(input)
	if got != want{
		t.Errorf("The FirstName is not Prankur")
	}
}

func TestIsEven(t *testing.T){
	// Data Table
	tests := []struct{
		input int 
		want bool 
	}{
		{input: 2 , want: true},
		{input: 7 , want: false},
		{input: 0 , want: true},
		{input: -4 , want: true},
	}

	// 2. Loop through table
	for _, tt := range tests{
		got := IsEven(tt.input)
		if got != tt.want {
			t.Errorf("IsEven(%d) = %v; want %v" , tt.input , got , tt.want)
		}
	}
}

func TestIsEven_Subsets(t *testing.T){
	tests := []struct {
		name string 
		input int 
		want bool 
	}{
		{name: "Positive Even" , input : 2 , want : true},
		{name: "Positive Odd" , input: 7 , want: false},
	}

	for _, tt := range tests{
		t.Run(tt.name , func(t *testing.T){
			got := IsEven(tt.input)
			if got != tt.want{
				t.Errorf("got %v; want %v" , got , tt.want)
			}
		})
	}
}

