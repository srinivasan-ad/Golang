package producer

import (
    "encoding/csv"
    "os"
    "log"
)
func ReadCSV() ([][]string, error) {
    file, err := os.Open("./csv/data.csv")
    if err != nil {
        log.Fatalf("Could not open the CSV file: %v", err)
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatalf("Error while reading file: %v", err)
        return nil, err
    }

    return records, nil
}