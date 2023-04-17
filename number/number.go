package number

import (
	"fmt"
)

func FormatNumber(n float64) string {
	if n < 1000 {
		return fmt.Sprintf("%f", n)
	}
	if n < 1000000 {
		return fmt.Sprintf("%.1fK", float64(n)/1000)
	}
	return fmt.Sprintf("%.1fM", float64(n)/1000000)
}
