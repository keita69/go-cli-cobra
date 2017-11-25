package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-keisan",
	Short: "This tool is pretty cool.",
	Long:  "This tool is a great convenience.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// menuCmd represents the menu command
var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		showMenu()
	},
}

func init() {
	rootCmd.AddCommand(menuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// menuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// menuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showMenu() {
	fmt.Println("/////// MENU ///////////////////////////////////////////////")
	cmds := rootCmd.Commands()
	for i, c := range cmds {
		fmt.Printf("//  [%d] %s --- %s\n", i, c.Use, c.Short)
	}
	fmt.Println("////////////////////////////////////////////////////////////")
	fmt.Printf(">>> Please select number [0] - [%d] or [99]\n", len(cmds)-1)
	fmt.Print(">>> num = ")
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := sc.Text()
	switch input {
	case "0":
		addCmdHelper() // [TODO] Cmd may not be "case 0" always.
	case "1":
		helpCmdHelper() // [TODO] Cmd may not be "case 0" always.
	case "2":
		menuCmdHelper() // [TODO] Cmd may not be "case 0" always.
	case "99":
		fmt.Printf(">>> [%s] is Exit.\n", input)
		return
	default:
		fmt.Printf(">>> [%s] is Bud Number.\n", input)
		return
	}
}

func addCmdHelper() {
	var sc = bufio.NewScanner(os.Stdin)
	fmt.Print(">>> x = ")
	sc.Scan()
	x := sc.Text()
	sc = bufio.NewScanner(os.Stdin)
	fmt.Print(">>> y = ")
	sc.Scan()
	y := sc.Text()
	cmd := exec.Command("go-keisan", "add", x, y)
	cmd.Wait()
	sum, err := cmd.Output()
	if err != nil {
		fmt.Println("Error!!")
	}
	fmt.Printf(">>> %s + %s = %s\n", x, y, sum)
}

func helpCmdHelper() {
	fmt.Println("execute helpCmdHelper !")
	return
}

func menuCmdHelper() {
	fmt.Println("execute menuCmdHelper !")
	return
}
