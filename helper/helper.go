package helper

import (
	"fmt"
	"github.com/merveuluser/go-converter-app/conversions"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var conversionType string
var valueStr string

func RoundFloat(v float64) float64 {
	ratio := math.Pow(10, 2)
	return math.Round(v*ratio) / ratio
}

func Web() bool {
	if len(os.Args) == 2 && os.Args[1] == "web" {
		return true
	}
	return false
}

func Cli() bool {
	if len(os.Args) == 4 && os.Args[1] == "convert" {
		return true
	}
	return false
}

func WebConversion() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		parts := strings.Split(strings.TrimPrefix(path, "/"), "=")
		if len(parts) != 2 {
			http.Error(w, "Invalid URL format. Usage: /<conversion_type>=<value>", http.StatusBadRequest)
			return
		}

		conversionType := parts[0]
		valueStr := parts[1]

		conversionFunc, exists := ConversionFuncExist(conversionType)
		if !exists {
			fmt.Fprintf(w, "Unsupported conversion type. Supported conversions:\n\nmeterstoinches\nkilometerstomiles\nkilometerstomiles\nkilogramstopounds\ncelsiustofahrenheit\nfahrenheittocelsius\ncelsiustokelvin\nkelvintocelsius\nfahrenheittokelvin\nkelvintofahrenheit")
			return
		}
		value, err := ValueStrToValue(valueStr)
		if err != nil {
			fmt.Fprintf(w, "Invalid value. Please provide a numeric value.")
			return
		}

		result := conversionFunc(value)
		fmt.Fprintf(w, "Converted value: %.2f\n", result)
		// find a way to format this like x fahrenheit is y celsius based on the conversion type and print the result with its correct unit
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func CliConversion() {
	conversionType := os.Args[2]
	valueStr := os.Args[3]

	conversionFunc, exists := ConversionFuncExist(conversionType)
	if !exists {
		fmt.Println("Unsupported conversion type. Supported conversions:\n\nmeterstoinches\nkilometerstomiles\nkilometerstomiles\nkilogramstopounds\ncelsiustofahrenheit\nfahrenheittocelsius\ncelsiustokelvin\nkelvintocelsius\nfahrenheittokelvin\nkelvintofahrenheit")
		return
	}
	value, err := ValueStrToValue(valueStr)
	if err != nil {
		fmt.Println("Invalid value. Please provide a numeric value.")
		return
	}

	result := conversionFunc(value)
	fmt.Printf("Converted value: %.2f\n", result)
	// find a way to format this like x fahrenheit is y celsius based on the conversion type and print the result with its correct unit
}

func ConversionFuncExist(conversionType string) (conversions.ConversionFunc, bool) {
	conversionFunc, exists := conversions.ConversionMap[conversionType]
	return conversionFunc, exists
}

func ValueStrToValue(valueStr string) (float64, error) {
	value, err := strconv.ParseFloat(valueStr, 64)
	return value, err
}
