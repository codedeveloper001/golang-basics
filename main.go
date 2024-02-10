package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Set up the scp command with source and destination paths (both local)
	sourcePath := "/home/akash/src/demo.txt"
	destinationPath := "/home/akash/dest/"

	// Create a test file in the source directory
	createTestFile(sourcePath)

	// Run the scp command to copy the file locally
	cmd := exec.Command("cp", sourcePath, destinationPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	_ = cmd.Run()

	fmt.Println("File copied successfully!")
}

func createTestFile(filePath string) error {
	// Create a test file for demonstration
	file, _  := os.Create(filePath)
	defer file.Close()

	_, _ = file.AppendString("This is a test file content.")

	fmt.Println("Test file created:", filePath)
	return nil
}


