package helper

import (
	"fmt"
	"github.com/merveuluser/go-converter-app/conversions"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func WebConversion() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		parts := strings.Split(strings.TrimPrefix(path, "/"), "=")
		if len(parts) != 2 {
			http.Error(w, InvalidURLFormat, http.StatusBadRequest)
			return
		}

		conversionType := parts[0]
		valueStr := parts[1]

		conversionFunc, exists := ConversionFuncExist(conversionType)
		if !exists {
			fmt.Fprintf(w, UnsupportedConversionType)
			return
		}
		value, err := ValueStrToValue(valueStr)
		if err != nil {
			fmt.Fprintf(w, InvalidValue)
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
		fmt.Println(UnsupportedConversionType)
		fmt.Println(strings.Join(conversions.ConversionTypes(), "\n"))
		return
	}
	value, err := ValueStrToValue(valueStr)
	if err != nil {
		fmt.Println(InvalidValue)
		return
	}

	result := conversionFunc(value)
	fmt.Printf("Converted value: %.2f\n", result)
	// find a way to format this like x fahrenheit is y celsius based on the conversion type and print the result with its correct unit
}

func ConversionFuncExist(conversionType string) (conversions.ConversionFunc, bool) {
	conversionFunc, exists := conversions.GetConversionFunc(conversionType)
	return conversionFunc, exists
}

func ValueStrToValue(valueStr string) (float64, error) {
	value, err := strconv.ParseFloat(valueStr, 64)
	return value, err
}
