package datetime

import (
	"fmt"
	"testing"
	"time"
)

func Test_time(t *testing.T){ 
	now:= time.Now()
	fmt.Println(now)


	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())
	fmt.Println(now.Location())
	fmt.Println(now.Weekday())

	then:= time.Date(2022,11,17,20,34,58 ,0 ,time.UTC)

	fmt.Println(then.Before(now))
	fmt.Println(then.After(now))
	fmt.Println(then.Equal(now))

	diff:= now.Sub(then)
	fmt.Println(diff)

	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Seconds())
	fmt.Println(diff.Nanoseconds())
}