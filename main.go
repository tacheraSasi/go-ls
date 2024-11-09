package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	
	// Getting the current directory or take from args
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	switch os.Args[1] {
	case "list":
		listFiles(dir)
	case "count line":
		countLines(os.Args[2])
	default:
		log.Fatal("please enter a valid option")
		
	}
	

	
}

func countLines(filePath string){
	if filePath == "" {
		log.Fatal("Please specify a file")
		return
	}
	file,err := os.Open(filePath) 
	if err != nil{
		log.Fatal("Something went wrong while reading",filePath," :",err)
		return
	}

	defer file.Close()// closing the file once the function is finished

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	//getting the number of lines
	var count int 
	for scanner.Scan(){
		count ++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return
	 }
	 fmt.Println("Number of lines in file:", count)

}

func listFiles(dirPath string){
	// Opening the directory
	directory, err := os.Open(dirPath)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer directory.Close()

	// Reading the directory contents
	files, err := directory.Readdir(-1) // -1 means read all entries
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// Looping through the files and display information
	fmt.Println("Listing files in directory:", dirPath)
	for _, file := range files {
		// File name
		fmt.Printf("%-25s", file.Name())

		// File size
		fmt.Printf("%-10d", file.Size())

		// File permissions
		fmt.Printf("%-10s", file.Mode().String())

		// Last modified time
		modTime := file.ModTime().Format(time.RFC1123)
		fmt.Printf("%-20s\n", modTime)
	}

}
