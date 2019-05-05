package gorbi

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func writeToCsv(data []float64,filename string) {
	file, err := os.Create(filename)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write([]string{fmt.Sprintf("%f", value)})
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}



