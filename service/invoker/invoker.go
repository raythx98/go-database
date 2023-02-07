package invoker

import (
	"github.com/raythx98/go-database/service/command"
)

type Invoker struct {
	Command command.Command
}

func (b *Invoker) Invoke() (string, error) {
	/*
		start goroutine to execute command,
		suitable if response is not needed - EDA
	*/
	// go b.Command.Execute()
	return b.Command.Execute()
}
