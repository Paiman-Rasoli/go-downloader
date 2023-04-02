package request

import (
	"time"

	"github.com/briandowns/spinner"
)

func lunchSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
      s.Prefix = " Downloading   "
	s.Suffix = "\n"
	return s
}