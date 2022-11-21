package main

import (
	"bufio"
	"fmt"
	"os"
)

func writeFile() {
	f, err := os.Create("file.txt")

	if err != nil {
		panic(err)
	}

	size, err := f.Write([]byte("Writing data to file"))
	// size, err := f.WriteString("Hello, World!")

	if err != nil {
		panic(err)
	}

	fmt.Printf("File successfuly created! Size: %d bytes\n", size)

	f.Close()
}

func readFile() {
	f, err := os.ReadFile("file.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(f))
}

func readSlices() {
	file, err := os.Open("file.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)

	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}
}

func removeFile() {
	err := os.Remove("file.txt")

	if err != nil {
		panic(err)
	}
}

func main() {
	removeFile()
}
