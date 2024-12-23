package main

import (
	"errors"
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("echo", "hello", "world")
	if errors.Is(cmd.Err, exec.ErrDot) {
		cmd.Err = nil
	}

	/* 	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	} */
	stdout, outErr := cmd.Output()
	if outErr != nil {
		log.Fatalln(outErr)
	}

	fmt.Println(string(stdout))

	/*
		 	err := repository.Migrate()
			if err != nil {
				log.Fatalln(err)
			}
	*/
}
