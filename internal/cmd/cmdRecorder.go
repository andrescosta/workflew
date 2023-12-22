package cmd

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/andrescosta/jobico/api/pkg/remote"
)

var cmdRecorder = &command{
	name:      "recorder",
	usageLine: "cli recorder [tenant] [-lines <number>]",
	short:     "gets result information for the tenant",

	long: `
Recorder gets result information for the tenant.
The -lines <number> gets the logs lines  `,
}

var cmdRecorderflagLines *int

func initRecorder() {
	cmdRecorder.flag = *flag.NewFlagSet("recorder", flag.ContinueOnError)
	cmdRecorderflagLines = cmdRecorder.flag.Int("lines", 0, "")
	cmdRecorder.flag.Usage = func() {}
	cmdRecorder.run = runRecorder
}
func runRecorder(ctx context.Context, cmd *command, _ []string) {
	ch := make(chan string)
	go func(mc <-chan string) {
		for {
			select {
			case <-ctx.Done():
				return
			case l := <-mc:
				fmt.Println(l)
			}
		}
	}(ch)
	fmt.Printf("getting results at proc: %d \n", os.Getpid())
	client, err := remote.NewRecorderClient(ctx)
	if err != nil {
		printError(os.Stderr, cmd, err)
	}
	if err := client.StreamJobExecutions(ctx, "", int32(*cmdRecorderflagLines), ch); err != nil {
		printError(os.Stderr, cmd, err)
	}
	fmt.Println("command stoped.")
}