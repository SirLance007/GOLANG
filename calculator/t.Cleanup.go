package calculator

import (
	"fmt"
	"os"
	"testing"
)

func TestCleanupDemo(t *testing.T) {
	// 1. Terminal par seedha print karne ke liye fmt.Println use kar rahe hain
	fmt.Println("➡️ [STEP 1] Test Start hua: Ek temporary file bana rahe hain...")

	file, err := os.CreateTemp("", "demo_*.txt")
	if err != nil {
		t.Fatalf("File nahi bani: %v", err)
	}

	// 2. Go ko 't.Cleanup' register karwaya
	t.Cleanup(func() {
		os.Remove(file.Name()) // File disc se delete karo
		fmt.Println("🧹 [STEP 3] t.Cleanup RUN HUA: Temp file deleted successfully!")
	})

	// 3. Test ka real kaam
	fmt.Println("⚡ [STEP 2] Test ka logic chal raha hai...")
}