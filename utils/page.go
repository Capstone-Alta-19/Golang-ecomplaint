package utils

func GetOffset(limit, page int) int {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	return (page - 1) * limit
}
