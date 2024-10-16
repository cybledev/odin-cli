package decorator

import (
	"encoding/json"
	"fmt"
	"time"

	bspinner "github.com/briandowns/spinner"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/fatih/color"
	"github.com/nwidger/jsoncolor"
	"github.com/spf13/cobra"
)

func RunEWrapper(f func(cmd *cobra.Command, args []string) (string, error)) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		s := spinner.New()
		s.Spinner = spinner.Dot

		m := model{spinner: s, dataChan: make(chan string)}
		go func() {
			result, err := f(cmd, args)
			if err != nil {
				return
			}
			m.dataChan <- result
		}()

		p := tea.NewProgram(m)
		if _, err := p.Run(); err != nil {
			return fmt.Errorf("error starting Bubble Tea program: %w", err)
		}

		return nil
	}
}

func RunEColorWrapper(f func(cmd *cobra.Command, args []string) (string, error)) func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		s := bspinner.New(bspinner.CharSets[9], 100*time.Millisecond)
		s.Suffix = " Fetching data..."
		s.Start()

		defer s.Stop()

		result, err := f(cmd, args)
		if err != nil {
			return fmt.Errorf("error fetching data: %v", err)
		}

		var data map[string]interface{}
		err = json.Unmarshal([]byte(result), &data)
		if err != nil {
			return err
		}
		isRawOutput, err := cmd.Flags().GetBool("raw-output")
		if err == nil && isRawOutput {
			fmt.Println(result)
			return nil
		}

		s.Stop()
		f := jsoncolor.NewFormatter()

		// set custom colors
		f.StringColor = color.New(color.FgBlack, color.Bold)
		f.TrueColor = color.New(color.FgWhite, color.Bold)
		f.FalseColor = color.New(color.FgRed)
		f.NumberColor = color.New(color.FgWhite)
		f.NullColor = color.New(color.FgWhite, color.Bold)

		dst, err := jsoncolor.MarshalIndentWithFormatter(data, "", "  ", f)
		if err != nil {
			return err
		}
		fmt.Println(string(dst))
		return nil
	}
}
