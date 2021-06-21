package bizdays

import (
	"math"
	"time"
)

func Between(a, b time.Time) (days int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}

	if a.Weekday() == time.Saturday {
		a = a.Add(time.Hour * 48)
	} else if a.Weekday() == time.Sunday {
		a = a.Add(time.Hour * 24)
	}

	if b.Weekday() == time.Saturday {
		b = b.Add(-(time.Hour * 24))
	} else if b.Weekday() == time.Sunday {
		b = b.Add(-(time.Hour * 48))
	}

	days = int(math.Floor(b.Sub(a).Hours()/24/7))*5 + (5+int(b.Weekday())-int(a.Weekday()))%5

	return
}
