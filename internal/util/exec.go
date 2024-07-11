package util

import (
	"bufio"
	"errors"
	"io"
	"os"
	"os/exec"
	
	"github.com/daarlabs/hrx/internal/log"
)

func Exec(name string, args ...string) {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	if err := cmd.Start(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
	go func(stdout io.ReadCloser) {
		defer func(stdout io.ReadCloser) {
			if err := stdout.Close(); err != nil {
				log.Error(err)
				os.Exit(1)
			}
		}(stdout)
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			log.Success(scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			log.Error(err)
		}
	}(stdout)
	go func(stderr io.ReadCloser) {
		defer func(stderr io.ReadCloser) {
			if err := stderr.Close(); err != nil {
				log.Error(err)
				os.Exit(1)
			}
		}(stderr)
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			log.Error(errors.New(scanner.Text()))
		}
	}(stderr)
	if err := cmd.Wait(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
