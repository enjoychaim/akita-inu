package number

import (
	"fmt"
)

func FormatNumber(n float64) string {
	if n < 1000 {
		return fmt.Sprintf("%.0f", n)
	}
	if n < 1000000 {
		return fmt.Sprintf("%.0fK", float64(n)/1000)
	}
	if n < 1000000000 {
		return fmt.Sprintf("%.0fM", float64(n)/1000000)
	}
	return fmt.Sprintf("%.0fB", float64(n)/1000000000)
}
