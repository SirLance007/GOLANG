package checkout

import (
	"os"
	"testing"
)

func setupMockLogger(t *testing.T) *os.File {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "checkout_log_*.txt")
	if err != nil {
		t.Fatalf("Failed to create tme")
	}

	// t.Cleanup registers teardown logic that runs AFTER the test completes
	t.Cleanup(func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	})

	return tmpFile
}

func TestCheckoutWithLoggin(t *testing.T) {
	logFile := setupMockLogger(t)

	// Write mock data...
	_, err := logFile.WriteString("Processing Order #1001..\n")
	if err != nil {
		t.Errorf("Failed to write to log: %v", err)
	}
}