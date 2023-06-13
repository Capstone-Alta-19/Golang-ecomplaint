package utils

func ConvertToNullString(str string) *string {
	if str == "" {
		return nil
	}
	return &str
}
