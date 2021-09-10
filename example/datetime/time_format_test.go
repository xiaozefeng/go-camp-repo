package datetime

import (
	"fmt"
	"testing"
	"time"
)

func Test_time_format(t *testing.T) {
	now := time.Now()
	fmt.Println(now.Format(time.RFC3339))
	fmt.Println(now.Format(time.ANSIC))
	fmt.Println(now.Format(time.RFC1123))

	fmt.Println(now.Format("2006-01-02 15:04:05"))
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	dateStr := `2021-09-10 17:00:22`

	if t1, err := time.Parse("2006-01-02 15:04:05", dateStr); err == nil {
		fmt.Println(t1)
	} else {
		fmt.Println(err)
	}

}
