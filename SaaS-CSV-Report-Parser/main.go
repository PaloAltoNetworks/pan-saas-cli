package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func main() {
	data := readCSVFile("SaaSApps.csv")
	grepApps(data)
}

// readCSVFile takes a CSV file and return a file descriptor. The file descriptor is used by the grepApps function.
func readCSVFile(fname string) *os.File {
	data, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

//grepApps reads the content in the file descriptor 
//grepApps then creates a text file with the same name as the input file and writes the application names to the file
func grepApps(fname *os.File) {

	r := csv.NewReader(fname)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	newFileName := strings.Replace(fname.Name(), "csv", "txt", 1)
	f, err := os.Create(newFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for i, s := range records {
		//Skips processing the title line
		if i == 0 {
			continue
		}
		//Greps the second item in the list i.e. the application name and writes to file
		f.WriteString(s[1] + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}

