package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func isAdmin() bool {
	_, err := os.Open("\\\\.\\PHYSICALDRIVE0")
	return err == nil
}

func main() {
	if !isAdmin() {
		fmt.Println("This program requires administrative privileges.")
		fmt.Println("Please run it as an administrator.")
		os.Exit(1)
	}

	// Get all network adapters
	output, err := exec.Command("netsh", "interface", "show", "interface").Output()
	if err != nil {
		fmt.Println("Error getting network adapters: ", err)
		return
	}

	// Split output into lines
	lines := strings.Split(string(output), "\n")

	// Process each line
	for _, line := range lines {
		// Skip header line and empty lines
		if !strings.Contains(line, "Enabled") && !strings.Contains(line, "Disabled") {
			continue
		}

		// Extract adapter name
		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}
		adapterName := strings.Join(fields[3:], " ")

		// Disable the adapter
		fmt.Printf("Disabling %s...\n", adapterName)
		_, err := exec.Command("netsh", "interface", "set", "interface", adapterName, "disabled").Output()
		if err != nil {
			fmt.Printf("Error disabling %s: %v\n", adapterName, err)
			continue
		}
	}

	// Wait for 5 seconds
	fmt.Println("Waiting for 5 seconds...")
	time.Sleep(5 * time.Second)

	// Re-enable all adapters
	for _, line := range lines {
		if !strings.Contains(line, "Enabled") && !strings.Contains(line, "Disabled") {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) < 4 {
			continue
		}
		adapterName := strings.Join(fields[3:], " ")

		fmt.Printf("Enabling %s...\n", adapterName)
		_, err := exec.Command("netsh", "interface", "set", "interface", adapterName, "enable").Output()
		if err != nil {
			fmt.Printf("Error enabling %s: %v\n", adapterName, err)
		}
	}

	fmt.Println("Network adapters reset complete.")
}
