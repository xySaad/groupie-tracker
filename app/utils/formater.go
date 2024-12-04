package utils

import "unicode"

func FormatLocations(locations Object) Object {
	formatedRelation := Object{}

	for location, dates := range locations {
		var locationR []rune
		spaceFound := true
		for _, char := range location {
			switch char {
			case '-':
				locationR = append(locationR, ',', ' ')
				spaceFound = true
			case '_':
				locationR = append(locationR, ' ')
				spaceFound = true
			default:
				if spaceFound {
					locationR = append(locationR, unicode.ToUpper(char))
					spaceFound = false
					continue
				}
				locationR = append(locationR, char)
			}
		}
		formatedLocation := string(locationR)
		formatedRelation[formatedLocation] = dates
	}
	return formatedRelation
}
