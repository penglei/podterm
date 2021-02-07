package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println(os.Args)
	if len(os.Args) > 1 {
		name := os.Args[1]
		args := os.Args[1:]
		f, err := exec.LookPath(name)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		err = syscall.Exec(f, args, os.Environ())
		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Printf("sleep...")
		closeCh := make(chan os.Signal)
		signal.Notify(closeCh)
		<-closeCh
	}
}
