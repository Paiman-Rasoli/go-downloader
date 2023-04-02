package request

import (
	"time"

	"fmt"

	"github.com/briandowns/spinner"
)

func lunchSpinner() *spinner.Spinner {
	fmt.Print("Downloading ..... ")
	s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)  // Build our new spinner
	return s
}