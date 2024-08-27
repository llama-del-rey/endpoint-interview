package cmd

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestCLIOutput(t *testing.T) {
	expectedOutputBytes, err := os.ReadFile("../resources/expected_output.txt")
	if err != nil {
		t.Fatalf("Failed to read expected output file: %v", err)
	}
	expectedOutput := strings.TrimSpace(string(expectedOutputBytes))

	cmd := exec.Command("../dirmanager", "FROMFILE", "../resources/input.txt")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Error running CLI: %v\nOutput: %s", err, string(output))
	}

	actualOutput := strings.TrimSpace(string(output))

	if actualOutput != expectedOutput {
		t.Errorf("Expected output:\n%s\n\nBut got:\n%s\n", expectedOutput, actualOutput)
	}
}
