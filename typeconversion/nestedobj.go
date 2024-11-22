package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"


)

type Input struct {
	TypeName string                 `json:"type name"`
	Type     map[string]interface{} `json:"type"`
}

func amain() {
	// Example JSON input
	jsonInput := `
	{
  "type name": "Order",
  "type": {
    "orderID": "int",
    "customer": {
      "name": "string",
      "email": "string",
      "address": {
        "street": "string",
        "city": "string",
        "state": "string",
        "zip": "string"
      }
    },
    "items": [
      {
        "product": {
          "name": "string",
          "price": "float64",
          "description": "string"
        },
        "quantity": "int",
        "totalPrice": "float64"
      }
    ],
    "status": "string",
    "orderDate": "string"
  }
}
`

	err := processAndAppend(jsonInput, "types.go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Struct appended to types.go successfully")
}

func processAndAppend(jsonInput string, fileName string) error {
	var input Input
	err := json.Unmarshal([]byte(jsonInput), &input)
	if err != nil {
		return fmt.Errorf("failed to parse JSON: %w", err)
	}

	structCode, err := generateStruct(input)
	if err != nil {
		return fmt.Errorf("failed to generate struct: %w", err)
	}

	err = appendTofile(fileName, structCode)
	if err != nil {
		return fmt.Errorf("failed to append to file: %w", err)
	}

	return nil
}
func generateStruct(input Input) (string,error) {
	structCode := fmt.Sprintf("type %s struct {\n", input.TypeName)
	var nestedStructs []string
	for fieldName, fieldType := range input.Type {
		fieldName := toCamelCase(fieldName)

		switch fieldtype := fieldType.(type) {
		case string:
			structCode += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName, fieldtype, fieldName)
		case map[string]interface{}:
			nestedStructName := toCamelCase(fieldName)
			nestedStruct, err := generateStruct(Input{
				TypeName: nestedStructName,
				Type:     fieldtype,
			})
			if err != nil {
				return "", fmt.Errorf("failed to generate nested struct: %w", err)
			}
			structCode += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", fieldName, nestedStructName, fieldName)
			nestedStructs = append(nestedStructs, nestedStruct)
		}
	}
	structCode += "}\n"
	for _, nestedStruct := range nestedStructs {
		structCode += nestedStruct
	}

	return structCode, nil
}

func appendTofile(fileName, content string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()
	_, err = file.Write([]byte(content))
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func toCamelCase(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = strings.ToUpper(string(word[0])) + word[1:]
	}
	return strings.Join(words, "")
}
