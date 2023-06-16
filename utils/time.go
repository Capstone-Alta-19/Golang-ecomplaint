package utils

import "fmt"

func GetTimeAgo(timeNow int64) string {
	switch {
	case timeNow < 60:
		return "Just Now"
	case timeNow < 3600:
		return fmt.Sprintf("%dm", timeNow/60)
	case timeNow < 86400:
		return fmt.Sprintf("%dh", timeNow/3600)
	case timeNow < 604800:
		return fmt.Sprintf("%dd", timeNow/86400)
	case timeNow < 2592000:
		return fmt.Sprintf("%dw", timeNow/604800)
	case timeNow < 31536000:
		return fmt.Sprintf("%dmo", timeNow/2592000)
	default:
		return fmt.Sprintf("%dy", timeNow/31536000)
	}
}
