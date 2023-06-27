package utils

import "strings"

func ConvertToNullString(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}

func ConvertDateToIndonesia(date string) string {
	dateInd := map[string]string{
		"Monday":    "Senin",
		"Tuesday":   "Selasa",
		"Wednesday": "Rabu",
		"Thursday":  "Kamis",
		"Friday":    "Jumat",
		"Saturday":  "Sabtu",
		"Sunday":    "Minggu",
	}

	for eng, ind := range dateInd {
		date = strings.Replace(date, eng, ind, -1)
	}
	return date
}
