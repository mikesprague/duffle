package main

import (
	"fmt"
	"io"

	"github.com/deis/duffle/pkg/version"
	"github.com/spf13/cobra"
)

func newVersionCmd(w io.Writer) *cobra.Command {
	const usage = `Prints current version of the Duffle CLI`

	cmd := &cobra.Command{
		Use:   "version",
		Short: usage,
		Long:  usage,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.Version)
		},
	}

	return cmd
}