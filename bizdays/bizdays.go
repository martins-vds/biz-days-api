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
		a = a.Add(2)
	} else if a.Weekday() == time.Sunday {
		a = a.Add(1)
	}

	if b.Weekday() == time.Saturday {
		b = b.Add(-1)
	} else if b.Weekday() == time.Sunday {
		b = b.Add(-2)
	}

	days = int(math.Floor(b.Sub(a).Hours()/24/7))*5 + (5+int(b.Weekday())-int(a.Weekday()))%5

	return
}
