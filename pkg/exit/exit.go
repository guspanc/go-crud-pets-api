package exit

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/guspanc/go-crud-pets-api/pkg/logger"
)

// Init exit handler
func Init(callback func()) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	terminate := make(chan bool)

	go func() {
		sig := <-sigs
		logger.INFO.Println("terminating on signal:", sig)
		close(terminate)
	}()

	<-terminate
	callback()
}
