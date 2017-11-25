package cmd

import (
	"os"
	"fmt"
	"bufio"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "go-keisan",
	Short: "This tool is pretty cool.",
	Long:  "This tool is a great convenience.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	cobra.OnInitialize()
	RootCmd.AddCommand(versionCmd)

	showMenu(RootCmd.Commands())
}

func showMenu(cmds []*cobra.Command) {
	fmt.Println("----------")
	for i, c := range cmds {
		fmt.Printf("[%d] %s --- %s\n", i, c.Use, c.Short)
	}
	fmt.Println("----------")
	fmt.Printf(">>> Please select number [0] - [%d] or [99]", len(cmds))
	fmt.Println(">>>") 
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	fmt.Println(input) 
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of go-keisan",
	Long:  `All software has versions. This is go-keisan's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-keisan v1.0")
	},
}
