package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	_framework_ "github.com/JosephABudd/kickfyne/commands/framework"
	_frontend_ "github.com/JosephABudd/kickfyne/commands/frontend"
	_help_ "github.com/JosephABudd/kickfyne/commands/help"
	_message_ "github.com/JosephABudd/kickfyne/commands/message"
	_record_ "github.com/JosephABudd/kickfyne/commands/record"
	_utils_ "github.com/JosephABudd/kickfyne/source/utils"
)

func main() {

	// Build the args to pass on to the handlers.
	lArgs := len(os.Args)
	if lArgs < 2 {
		fmt.Println(_help_.Usage())
		return
	}

	var err error
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer func() {
		ctxCancel()
		if err != nil {
			log.Println("Error: ", err.Error())
			os.Exit(1)
		}
	}()

	var pathWD string
	if pathWD, err = os.Getwd(); err != nil {
		return
	}

	var isBuilt bool
	if isBuilt, err = _utils_.IsBuilt(pathWD); err != nil {
		fmt.Println("NOT BUILT")
		return
	}
	fmt.Println("BUILT")
	var importPrefix string
	if importPrefix, err = _utils_.ImportPrefix(pathWD); err != nil {
		fmt.Println("Failure: Unable to read a go.mod file.")
		return
	}
	go notify(ctx, ctxCancel)

	switch os.Args[1] {
	case _framework_.Cmd:
		var handlerArgs []string
		if lArgs > 2 {
			handlerArgs = os.Args[2:]
		}
		err = _framework_.Handler(pathWD, handlerArgs, isBuilt, importPrefix)
	case _frontend_.CmdScreen:
		var handlerArgs []string
		if lArgs > 2 {
			handlerArgs = os.Args[1:]
		}
		err = _frontend_.Handler(pathWD, handlerArgs, isBuilt, importPrefix)
	case _help_.Cmd:
		var handlerArgs []string
		if lArgs > 2 {
			handlerArgs = os.Args[2:]
		}
		err = _help_.Handler(handlerArgs)
	case _message_.Cmd:
		var handlerArgs []string
		if lArgs > 2 {
			handlerArgs = os.Args[2:]
		}
		err = _message_.Handler(pathWD, handlerArgs, isBuilt, importPrefix)
	case _record_.Cmd:
		var handlerArgs []string
		if lArgs > 2 {
			handlerArgs = os.Args[2:]
		}
		err = _record_.Handler(pathWD, handlerArgs, isBuilt, importPrefix)
	default:
		fmt.Println(_help_.Usage())
	}
}

func notify(ctx context.Context, ctxCancel context.CancelFunc) {

	signalChan := make(chan os.Signal, 1)
	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)
	for {
		select {
		case <-ctx.Done():
			return
		case <-signalChan:
			ctxCancel()
			// terminate after second signal before callback is done
			go func() {
				<-signalChan
				os.Exit(1)
			}()
			return
		}
	}
}
