package time_meter

import "time"

type TimeMeter struct {
	preTime time.Time
}

func NewTimeMeter() *TimeMeter {
	return &TimeMeter{
		preTime: time.Now(),
	}
}

func (tm *TimeMeter) Cost() time.Duration {
	cost := time.Since(tm.preTime)
	tm.preTime = time.Now()
	return cost
}

