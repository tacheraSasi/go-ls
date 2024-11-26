package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "time"
	

)

// TODO: Add in the functionality to count the number of line in all files of a certain type
func main() {	
	// Checking if I have an option argument
	if len(os.Args) < 2 {
		defaultOutput()
		return
	}

	// Setting the option and directory/file path if available
	option := os.Args[1]
	dir := "."

	// Updating dir if I have a specific path
	if len(os.Args) > 2 {
		dir = os.Args[2]
	}

	// Using the chosen option to either list files or count lines
	switch option {
	case "list":
		listFiles(dir)
	case "count-lines":
		countLines(dir)
	default:
		defaultOutput()
	}
}

func countLines(filePath string) {
	if filePath == "" {
		log.Fatal("Please specify a file")
		return
	}

	// Opening the file to read line-by-line
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Something went wrong while reading %s: %v", filePath, err)
		return
	}
	defer file.Close() // will close the file when the function is done

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	//  counting the lines in the file
	var count int
	for scanner.Scan() {
		count++
	}

	// Check if there was an error scanning the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println("Number of lines in the file is", count)
}

func listFiles(dirPath string) {
	// opening the specified directory to read its contents
	directory, err := os.Open(dirPath)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer directory.Close() // Closing the directory after reading its contents

	//reads all entries in the directory
	files, err := directory.Readdir(-1) // -1 reads all entries
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	//  loops through each file and display its details
	// fmt.Println("Listing files in directory:", dirPath)
	tableTitle := fmt.Sprintf("LISTING FILES IN DIRECTORY: %s", dirPath)
	// for _, file := range files {
	// 	// Displaying file details
	// 	fmt.Printf("%-25s", file.Name()) // File name
	// 	fmt.Printf("%-10d", file.Size()) // File size
	// 	fmt.Printf("%-10s", file.Mode().String())
	// 	fmt.Print("     ")                             // File permissions
	// 	modTime := file.ModTime().Format(time.RFC1123) // Last modified time
	// 	fmt.Printf("%-20s\n", modTime)
	// }

	PrintFiles(files,tableTitle)
}

func defaultOutput() {
	// Displaying usage instructions if the user doesn't provide a valid option
	fmt.Println("Please enter a valid option:\n-- list <dirPath>\n-- count-lines <filePath>")
}
