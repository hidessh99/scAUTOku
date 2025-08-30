package utils

import (
	"fmt"
	"os/exec"
	"strings"
)

// ExecuteShellCommand executes a shell command and returns the output
func ExecuteShellCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("command execution failed: %v, output: %s", err, string(output))
	}
	return strings.TrimSpace(string(output)), nil
}

// ExecuteProjectScript executes a script from the project directory
func ExecuteProjectScript(scriptName string, args ...string) (string, error) {
	scriptPath := fmt.Sprintf("./project/%s", scriptName)
	allArgs := append([]string{}, args...)
	allArgs = append(allArgs, scriptPath)

	// On Windows, we might need to use bash or WSL
	cmd := exec.Command("bash", allArgs...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("script execution failed: %v, output: %s", err, string(output))
	}
	return strings.TrimSpace(string(output)), nil
}
