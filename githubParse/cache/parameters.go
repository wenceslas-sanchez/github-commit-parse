package cache

import "time"

type Parameters struct {
	d time.Duration // in seconds
	t time.Time
}

//func (p *Parameters) Default() *Parameters {
//	return &Parameters{
//		0, time.Time{},
//	}
//}
