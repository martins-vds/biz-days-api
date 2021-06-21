package bizdays

import (
	"testing"
	"time"
)

func TestBizDaysCalculation(t *testing.T) {
	cases := []struct {
		from, to time.Time
		days     int
	}{
		{from: time.Now(), to: time.Now(), days: 0},
		{from: time.Date(2021, time.June, 18, 0, 0, 0, 0, time.Local), to: time.Date(2021, time.June, 21, 0, 0, 0, 0, time.Local), days: 1},
		{from: time.Date(2021, time.June, 19, 0, 0, 0, 0, time.Local), to: time.Date(2021, time.June, 25, 0, 0, 0, 0, time.Local), days: 4},
		{from: time.Date(2021, time.June, 21, 0, 0, 0, 0, time.Local), to: time.Date(2021, time.June, 26, 0, 0, 0, 0, time.Local), days: 4},
	}

	for _, c := range cases {
		got := Between(c.from, c.to)
		want := c.days

		if got != c.days {
			t.Errorf("Between() = %d; want = %d", got, want)
		}
	}
}
