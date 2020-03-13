package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func wg(bin *string) func(*cobra.Command, []string) error {
	return func(_ *cobra.Command, args []string) error {
		cmd := exec.Command(*bin, args...)
		f, err := os.Stdin.Stat()
		if err != nil {
			return err
		}
		if f.Mode()&os.ModeNamedPipe != 0 {
			stdin, err := cmd.StdinPipe()
			if err != nil {
				return err
			}
			if _, err := io.Copy(stdin, bufio.NewReader(os.Stdin)); err != nil {
				return err
			}
			go func() {
				defer stdin.Close()
			}()
		}
		out, err := cmd.CombinedOutput()
		fmt.Printf("%s", out)
		if err != nil {
			return err
		}
		return nil
	}
}
