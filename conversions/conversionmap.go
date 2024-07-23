package conversions

type ConversionFunc func(float64) float64

type Conversion struct {
	conversionType string
	funcName       ConversionFunc
}

var ConversionList = []Conversion{
	{"meterstoinches", MetersToInches},
	{"kilometerstomiles", KilometersToMiles},
	{"gramstoounces", GramsToOunces},
	{"kilogramstopounds", KilogramsToPounds},
	{"celsiustofahrenheit", CelsiusToFahrenheit},
	{"fahrenheittocelsius", FahrenheitToCelsius},
	{"celsiustokelvin", CelsiusToKelvin},
	{"kelvintocelsius", KelvinToCelsius},
	{"fahrenheittokelvin", FahrenheitToKelvin},
	{"kelvintofahrenheit", KelvinToFahrenheit},
}

func ConversionTypes() []string {
	conversionTypes := make([]string, 0, len(ConversionList))
	for _, conversion := range ConversionList {
		conversionTypes = append(conversionTypes, conversion.conversionType)
	}
	return conversionTypes
}

func GetConversionFunc(conversionType string) (ConversionFunc, bool) {
	for _, conversion := range ConversionList {
		if conversion.conversionType == conversionType {
			return conversion.funcName, true
		}
	}
	return nil, false
}
