package util

import (
	"fmt"
	"time"
)

// Duration mearsures and logs the cost time
func Duration(start time.Time, name string, f func()) {
	f()
	fmt.Println(name, " cost: ", time.Since(start))
}
