package run_test

import (
	"io"
	"log"
	"os"

	"github.com/amonks/run"
)

// ui implements MultiWriter
var _ run.MultiWriter = ui{}

type ui struct{}

func (w ui) Writer(string) io.Writer {
	return os.Stdout
}

// In this example, we build a version of the run CLI tool that uses a UI we
// provide ourselves.
func Example_bringYourOwnUI() {
	tasks, err := run.Load(".")
	if err != nil {
		log.Fatal(err)
	}

	run, err := run.RunTask(".", tasks, "dev")
	if err != nil {
		log.Fatal(err)
	}

	ui := ui{}

	run.Start(ui)

	if err := <-run.Wait(); err != nil {
		log.Fatal(err)
	}
}
