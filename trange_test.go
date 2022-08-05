package trange_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/worldline-go/trange"
)

const (
	DateLayout = "2006-01-02"
)

func TestDay(t *testing.T) {
	testCases := map[string]struct {
		DateStr   string
		Location  string
		WantStart func(loc *time.Location) time.Time
		WantEnd   func(loc *time.Location) time.Time
		HasError  bool
	}{
		"Test with Europe/Amsterdam": {
			DateStr:  "2022-08-05",
			Location: "Europe/Amsterdam",
			WantStart: func(loc *time.Location) time.Time {
				return time.Date(2022, 8, 5, 0, 0, 0, 0, loc)
			},
			WantEnd: func(loc *time.Location) time.Time {
				return time.Date(2022, 8, 5, 23, 59, 59, 999_999_999, loc)
			},
			HasError: false,
		},
		"Test with Invalid Date": {
			DateStr:  "Invalid",
			Location: "Europe/Amsterdam",
			WantStart: func(loc *time.Location) time.Time {
				return time.Time{}
			},
			WantEnd: func(loc *time.Location) time.Time {
				return time.Time{}
			},
			HasError: true,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			loc, err := time.LoadLocation(tc.Location)
			assert.NoError(t, err)

			start, end, err := trange.Day(tc.DateStr, loc)

			if tc.HasError {
				assert.NotNil(t, err)
			} else {
				assert.NoError(t, err)
			}

			assert.Equal(t, start, tc.WantStart(loc))
			assert.Equal(t, end, tc.WantEnd(loc))
		})
	}
}
