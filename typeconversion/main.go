// package main

// import (
// 	"fmt"
// 	// "os"
// 	"encoding/json"
// 	"strings"
// )

// // type Input struct {
// // 	TypeName string                 `json:"type name"`
// // 	Type     map[string]interface{} `json:"type"`
// // }

// // func Literal() {
// // 	// Example JSON input
// // 	jsonInput := `
// // 	{
// // 	"type name": "Product",
// // 	"type": {
// // 		"name": "string",
// // 		"price": "float64",
// // 		"stock": "int",
// // 		"available": "bool"
// // 	}
// // }
// // `

// // 	err := processAndAppend(jsonInput, "types.go")
// // 	if err != nil {
// // 		fmt.Printf("Error: %v\n", err)
// // 		return
// // 	}

// // 	fmt.Println("Struct appended to types.go successfully")
// // }

// // func processAndAppend(jsonInput string, fileName string) error {
// // 	var input Input
// // 	err := json.Unmarshal([]byte(jsonInput), &input)
// // 	if err != nil {
// // 		return fmt.Errorf("failed to parse JSON: %w", err)
// // 	}

// // 	structCode, err := generateStruct(input)
// // 	if err != nil {
// // 		return fmt.Errorf("failed to generate struct: %w", err)
// // 	}

// // 	err = appendToFile(fileName, structCode)
// // 	if err != nil {
// // 		return fmt.Errorf("failed to append to file: %w", err)
// // 	}

// // 	return nil
// // }

// // func generateStruct(input Input) (string, error) {

// // 	structCode := fmt.Sprintf("type %s struct {\n", input.TypeName)

// // 	for fieldName, fieldType := range input.Type {

// // 		goFieldName := toCamelCase(fieldName)

// // 		structCode += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", goFieldName, fieldType, fieldName)
// // 	}

// // 	structCode += "}\n"

// // 	return structCode, nil
// // }

// // func appendToFile(fileName, content string) error {
// // 	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModePerm)
// // 	if err != nil {
// // 		return fmt.Errorf("failed to open file: %w", err)
// // 	}
// // 	defer file.Close()
// // 	_, err = file.Write([]byte(content))
// // 	if err != nil {
// // 		return fmt.Errorf("failed to write to file: %w", err)
// // 	}

// // 	return nil
// // }

// //	func toCamelCase(s string) string {
// //		words := strings.Fields(s)
// //		for i, word := range words {
// //			words[i] = strings.ToUpper(string(word[0])) + word[1:]
// //		}
// //		return strings.Join(words, "")
// //	}
// func capitalizeFirstLetter(s string) string {
// 	if len(s) == 0 {
// 		return s
// 	}
// 	return strings.ToUpper(string(s[0])) + s[1:]
// }

// func recur(m *map[string]interface{}, key1 string, s *string, validateLogic *string, parent string) {
// 	structDef := "type " + capitalizeFirstLetter(key1) + "Type struct {\n"
// 	validationLogic := ""

// 	for key, value := range *m {
// 		switch val := value.(type) {
// 		case string:
// 			var fieldType string
// 			var min, max int
// 			minProvided, maxProvided := false, false

// 			parts := strings.Fields(val)
// 			fieldType = parts[0]

// 			for _, part := range parts[1:] {
// 				if strings.HasPrefix(part, "min=") {
// 					fmt.Sscanf(part, "min=%d", &min)
// 					minProvided = true
// 				} else if strings.HasPrefix(part, "max=") {
// 					fmt.Sscanf(part, "max=%d", &max)
// 					maxProvided = true
// 				}
// 			}
// 			structDef += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", capitalizeFirstLetter(key), fieldType, key)
// 			fieldPath := parent + capitalizeFirstLetter(key)
// 			if fieldType == "string" {
// 				if minProvided {
// 					validationLogic += fmt.Sprintf("gt(len(%s), %d) && ", fieldPath, min)
// 				}
// 				if maxProvided {
// 					validationLogic += fmt.Sprintf("lt(len(%s), %d) && ", fieldPath, max)
// 				}
// 			} else {
// 				if minProvided {
// 					validationLogic += fmt.Sprintf("gt(%s, %d) && ", fieldPath, min)
// 				}
// 				if maxProvided {
// 					validationLogic += fmt.Sprintf("lt(%s, %d) && ", fieldPath, max)
// 				}
// 			}

// 		case map[string]interface{}:
// 			nestedType := capitalizeFirstLetter(key1 + capitalizeFirstLetter(key))
// 			structDef += fmt.Sprintf("\t%s %sType `json:\"%s\"`\n", capitalizeFirstLetter(key), nestedType, key)
// 			recur(&val, nestedType, s, validateLogic, parent+capitalizeFirstLetter(key)+".")
// 		default:
// 			fmt.Println("Error!")
// 		}
// 	}
// 	structDef += "}\n\n"
// 	*s = structDef + *s
// 	*validateLogic += validationLogic
// }

// func main() {
// 	Schema := map[string]interface{}{
// 		"userName": "string min=5 max=50",
// 		"id":       "int min=1 max=100",
// 		"email":    "string min=10 max=100",
// 		"address": map[string]interface{}{
// 			"street": "string min=10 max=100",
// 			"city":   "string min=3 max=50",
// 			"zip":    "int max=9999",
// 			"country": map[string]interface{}{
// 				"name": "string min=3 max=50",
// 				"code": "string min=2 max=10",
// 			},
// 		},
// 		"preferences": map[string]interface{}{
// 			"theme":                "string min=3 max=20",
// 			"notificationsEnabled": "bool",
// 			"language":             "string min=2 max=10",
// 		},
// 	}
// 	var schema map[string]interface{}
// 	json.Unmarshal([]byte(Schema), &schema)
// 	code := ""
// 	validateLogic := ""
// 	recur(&schema, "User", &code, &validateLogic, "u.")
// 	validateFunction := "func Validate(u *UserType) bool {\n\treturn " + validateLogic[:len(validateLogic)-4] + "\n}\n\n"
// 	code = validateFunction + code
// 	fmt.Println(code)
// }
package main

import (
	"fmt"
	"encoding/json"
	"strings"
)

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
func recur(m *map[string]interface{}, key1 string, s *string, validateLogic *string, parent string) {
	structDef := "type " + capitalizeFirstLetter(key1) + "Type struct {\n"
	validationLogic := ""

	for key, value := range *m {
		switch val := value.(type) {
		case string:
			var fieldType string
			var min, max ,length int
			var startStr, endStr, containsStr, regexPattern string
			minProvided, maxProvided, lengthProvided, startStrProvided, endStrProvided, containsStrProvided, regexProvided   := false, false, false, false, false, false, false

			parts := strings.Fields(val)
			fieldType = parts[0]

			for _, part := range parts[1:] {
				if strings.HasPrefix(part, "min=") {
					fmt.Sscanf(part, "min=%d", &min)
					minProvided = true
				} else if strings.HasPrefix(part, "max=") {
					fmt.Sscanf(part, "max=%d", &max)
					maxProvided = true
				}else if strings.HasPrefix(part, "length=") {
					fmt.Sscanf(part, "length=%d", &length)
					lengthProvided = true
			}else if strings.HasPrefix(part, "startswith=") {
				fmt.Sscanf(part, "startswith=%s", &startStr)
				startStrProvided  = true
			}else if strings.HasPrefix(part, "startswith=") {
				fmt.Sscanf(part, "startswith=%s", &endStr)
				endStrProvided  = true
			}else if strings.HasPrefix(part, "contains=") {
				fmt.Sscanf(part, "contains=%s", &containsStr)
				containsStrProvided = true
			}else if strings.HasPrefix(part, "regex=") {
				fmt.Sscanf(part, "regex=%s", &regexPattern)
				regexProvided = true
			}
			}
			structDef += fmt.Sprintf("\t%s %s `json:\"%s\"`\n", capitalizeFirstLetter(key), fieldType, key)
			fieldPath := parent + capitalizeFirstLetter(key)
			if fieldType == "string" {
				if lengthProvided {
					validationLogic += fmt.Sprintf("len(%s) == %d && ", fieldPath, length)
				} else{
					if minProvided {
						validationLogic += fmt.Sprintf("gt(len(%s), %d) && ", fieldPath, min)
					}
					if maxProvided {
						validationLogic += fmt.Sprintf("lt(len(%s), %d) && ", fieldPath, max)
					}
					
				}
				if startStrProvided{
					validationLogic += fmt.Sprintf("ststr(%s, %s) && ", fieldPath, startStr)
				}
				if endStrProvided{
					validationLogic += fmt.Sprintf("endstr(%s, %s) && ", fieldPath, endStr)
				}
				if containsStrProvided{
					validationLogic += fmt.Sprintf("constr(%s, %s) && ", fieldPath, containsStr)
				}
				if regexProvided {
                    validationLogic += fmt.Sprintf("reg(%s, \"%s\") && ", fieldPath, regexPattern)
                }
				if key == "email" {
					validationLogic += fmt.Sprintf("em(%s) && ", fieldPath)
				}
				if key == "url" {
					validationLogic += fmt.Sprintf("url(%s) && ", fieldPath)
				}
			
				} else {
				if minProvided {
					validationLogic += fmt.Sprintf("gt(%s, %d) && ", fieldPath, min)
				}
				if maxProvided {
					validationLogic += fmt.Sprintf("lt(%s, %d) && ", fieldPath, max)
				}
			}

		case map[string]interface{}:
			nestedType := capitalizeFirstLetter(key1 + capitalizeFirstLetter(key))
			structDef += fmt.Sprintf("\t%s %sType `json:\"%s\"`\n", capitalizeFirstLetter(key), nestedType, key)
			recur(&val, nestedType, s, validateLogic, parent+capitalizeFirstLetter(key)+".")
		default:
			fmt.Println("Error!")
		}
	}
	structDef += "}\n\n"
	*s = structDef + *s
	*validateLogic += validationLogic
}
func main() {
	Schema := `{
	
 
    "empName": "string min=5 max=50",
    "id": "int min=1 max=100",
    "email": "string min=10 max=100",
    "phone": "string length=10",
    "website": "string regex=^https?://[a-zA-Z0-9.-]+$",
    "address": {
        "street": "string min=10 max=100",
        "city": "string min=3 max=50",
        "zip": "int max=9999",
        "country": {
            "name": "string min=3 max=50",
            "code": "string min=2 max=10"
        }
    },
    "preferences": {
        "theme": "string min=3 max=20",
        "notificationsEnabled": "bool",
        "language": "string min=2 max=10",
        "receiveEmails": "bool"
    },
    "socialLinks": {
        "linkedin": "string regex=^https?://(www\\.)?linkedin\\.com/.*$",
        "github": "string regex=^https?://(www\\.)?github\\.com/.*$"
    }



	  }`
	var schema map[string]interface{}
	json.Unmarshal([]byte(Schema), &schema)
	code := ""
	validateLogic := ""
	recur(&schema, "User", &code, &validateLogic, "u.")
	validateFunction := "func Validate(u *UserType) bool {\n\treturn " + validateLogic[:len(validateLogic)-4] + "\n}\n\n"
	code = validateFunction + code
	fmt.Println(code)
}
