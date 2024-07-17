package main

import (
	"fmt"
	"github.com/merveuluser/go-converter-app/conversions"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type ConversionFunc func(float64) float64

var conversionMap = map[string]ConversionFunc{
	"meterstoinches":      conversions.MetersToInches,
	"kilometerstomiles":   conversions.KilometersToMiles,
	"gramstoounces":       conversions.GramsToOunces,
	"kilogramstopounds":   conversions.KilogramsToPounds,
	"celsiustofahrenheit": conversions.CelsiusToFahrenheit,
	"fahrenheittocelsius": conversions.FahrenheitToCelsius,
	"celsiustokelvin":     conversions.CelsiusToKelvin,
	"kelvintocelsius":     conversions.KelvinToCelsius,
	"fahrenheittokelvin":  conversions.FahrenheitToKelvin,
	"kelvintofahrenheit":  conversions.KelvinToFahrenheit,
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			parts := strings.Split(strings.TrimPrefix(path, "/"), "=")
			if len(parts) != 2 {
				http.Error(w, "Invalid URL format. Usage: /<conversion_type>=<value>", http.StatusBadRequest)
				return
			}

			conversionType := parts[0]
			valueStr := parts[1]

			conversionFunc, exists := conversionMap[conversionType]
			if !exists {
				fmt.Fprintf(w, "Unsupported conversion type. Supported conversions:\n\nmeterstoinches\nkilometerstomiles\nkilometerstomiles\nkilogramstopounds\ncelsiustofahrenheit\nfahrenheittocelsius\ncelsiustokelvin\nkelvintocelsius\nfahrenheittokelvin\nkelvintofahrenheit")
				return
			}
			// get the keys from the map and print it

			value, err := strconv.ParseFloat(valueStr, 64)
			if err != nil {
				fmt.Fprintf(w, "Invalid value. Please provide a numeric value.")
				return
			}

			result := conversionFunc(value)
			fmt.Fprintf(w, "Converted value: %.2f\n", result)
			// find a way to format this like x fahrenheit is y celsius based on the conversion type and print the result with its correct unit
		}
		http.HandleFunc("/", handler)
		// ?is it better to assign the func to the variable and then use it or to define the func and then call it

		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return

	} else {
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run main.go <conversion_type> <value>")
			return
		}

		command := os.Args[1]
		conversionType := os.Args[2]
		valueStr := os.Args[3]

		if command != "convert" {
			fmt.Println("Invalid command. Use: go run main.go convert <conversion_type> <value>")
			return
		}

		conversionFunc, exists := conversionMap[conversionType]
		if !exists {
			fmt.Println("Unsupported conversion type.")
			return
		}

		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			fmt.Println("Invalid value. Please provide a numeric value.")
			return
		}

		result := conversionFunc(value)
		fmt.Printf("Converted value: %.2f\n", result)
		// find a way to format this like x fahrenheit is y celsius based on the conversion type and print the result with its correct unit
	}
}
