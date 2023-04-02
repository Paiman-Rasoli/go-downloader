package request

import (
	"time"

	"github.com/TwiN/go-color"
	"github.com/briandowns/spinner"
)

func lunchSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
      s.Prefix = color.InBlue("  downloading  ")
	return s
}