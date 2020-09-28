package exit

import (
	"syscall"
	"testing"
	"time"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		called   bool
		callback func() bool
	}{
		{
			desc:   "",
			called: false,
			callback: func() bool {
				return true
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			callback := func() {
				tC.called = tC.callback()
			}
			go func() {
				<-time.After(200 * time.Millisecond)
				syscall.Kill(syscall.Getpid(), syscall.SIGINT)
			}()
			Init(callback)
			if !tC.called {
				t.Error("expected `called` to be true")
			}
		})
	}
}
