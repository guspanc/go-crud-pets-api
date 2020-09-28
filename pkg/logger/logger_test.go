package logger

import "testing"

func Test(t *testing.T) {
	testCases := []struct {
		desc    string
		message string
	}{
		{
			desc:    "default",
			message: "message",
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			INFO.Println(tC.message)
		})
	}
}
