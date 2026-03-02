/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	tea "charm.land/bubbletea/v2"
	"github.com/spf13/cobra"
	"github.com/yorukot/termuru/internal"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "termuru",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(internal.IntinalModel())
		finalModel, err := p.Run()
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(1)
		}

		model, ok := finalModel.(internal.Model)
		if !ok {
			return
		}
		sshArgs := model.SSHArgs()
		if len(sshArgs) == 0 {
			return
		}

		// Ensure we are back on the normal screen and clear any TUI artifacts
		// before handing control to OpenSSH.
		fmt.Print("\x1b[?1049l\x1b[0m\x1b[?25h\x1b[2J\x1b[H")

		sshCmd := exec.Command("ssh", sshArgs...)
		sshCmd.Stdin = os.Stdin
		sshCmd.Stdout = os.Stdout
		sshCmd.Stderr = os.Stderr
		if err := sshCmd.Run(); err != nil {
			var exitErr *exec.ExitError
			if errors.As(err, &exitErr) {
				// Remote command failures should not be treated as application crashes.
				return
			}
			fmt.Println("error:", err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
