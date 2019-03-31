package cmd

import (
	"fmt"
	"os"

	"github.com/emmanuelCarre/ls/io/ioutil"
	"github.com/emmanuelCarre/ls/ls"
	"github.com/spf13/cobra"
)

var (
	data       ls.LsData
	showHidden bool
)

var rootCmd = &cobra.Command{
	Use:   "ls",
	Short: "list directory contents",
	Long: `List information about the FILEs (the current directory by default).  Sort entries alphabetically if none of -cftuvSUX nor --sort is specified.

Mandatory arguments to long options are mandatory for short options too.`,

	Run: func(cmd *cobra.Command, args []string) {
		data.ShowHidden = showHidden
		ls.Ls(ioutil.IoUtilImpl{}, data, "./")
	},
}

func Execute() {
	rootCmd.PersistentFlags().BoolVarP(&showHidden, "all", "a", false, "do not ignore entries starting with .")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
