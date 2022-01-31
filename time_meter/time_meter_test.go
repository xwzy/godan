package time_meter

import (
	"fmt"
	"testing"
	"time"
)

func TestNewTimeMeter(t *testing.T) {
	timer := NewTimeMeter()
	time.Sleep(time.Second * 2)
	fmt.Println(timer.Cost().Milliseconds())

	time.Sleep(time.Microsecond * 19283)
	fmt.Println(timer.Cost().Milliseconds())

	time.Sleep(time.Millisecond * 237)
	fmt.Println(timer.Cost().Milliseconds())
}
