package main

import "testing"

type User struct {
	ID   int
	Name string
	Age  int      
}

func FetchUser(id int) (User, error) {
	return User{ID: 1, Name: "Alice", Age: 30}, nil
}

func TestUser(t *testing.T) {
	user, err := FetchUser(1)
	if err != nil {
		t.Fatalf("Failed to fetch user: %v", err)
	}
	if user.Name != "Alice" {
		t.Errorf("Got name %q, want %q", user.Name, "Alice")
	}
	if user.Age != 30 {
		t.Errorf("Got Age %d; want %d", user.Age, 30)
	}
}

