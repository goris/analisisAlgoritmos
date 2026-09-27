package main

import "testing"

func TestPower(t *testing.T) {
	tests := []struct {
		name   string
		base   int64
		expo   int64
		result int64
	}{
		{
			name:   "2 pow 10",
			base:   2,
			expo:   10,
			result: 1024,
		},
		{
			name:   "3 pow 13",
			base:   3,
			expo:   13,
			result: 1594323,
		},
		{
			name:   "5 pow 25",
			base:   5,
			expo:   25,
			result: 298023223876953125,
		},
		{
			name:   "1 pow 1",
			base:   1,
			expo:   1,
			result: 1,
		},
		{
			name:   "0 pow 1",
			base:   0,
			expo:   1,
			result: 0,
		},
		{
			name:   "1 pow 0",
			base:   1,
			expo:   0,
			result: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := power(tt.base, tt.expo)

			if result != tt.result {
				t.Errorf(
					"power(%d, %d) = %d; expected: %d",
					tt.base,
					tt.expo,
					result,
					tt.result,
				)
			}
		})
	}
}
