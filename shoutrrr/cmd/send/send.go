package send

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"

	"github.com/circa10a/shoutrrr/internal/dedupe"
	internalUtil "github.com/circa10a/shoutrrr/internal/util"
	"github.com/circa10a/shoutrrr/pkg/router"
	"github.com/circa10a/shoutrrr/pkg/types"
	"github.com/circa10a/shoutrrr/pkg/util"
	cli "github.com/circa10a/shoutrrr/shoutrrr/cmd"
)

// MaximumNArgs defines the maximum number of arguments accepted by the command.
const (
	MaximumNArgs     = 2
	MaxMessageLength = 100
)

// Cmd sends a notification using a service URL.
var Cmd = &cobra.Command{
	Use:    "send",
	Short:  "Send a notification using a service url",
	Args:   cobra.MaximumNArgs(MaximumNArgs),
	PreRun: internalUtil.LoadFlagsFromAltSources,
	RunE:   Run,
}

func init() {
	Cmd.Flags().BoolP("verbose", "v", false, "")
	Cmd.Flags().StringArrayP("url", "u", []string{}, "The notification url")
	_ = Cmd.MarkFlagRequired("url")
	Cmd.Flags().
		StringP("message", "m", "", "The message to send to the notification url, or - to read message from stdin")

	_ = Cmd.MarkFlagRequired("message")
	Cmd.Flags().StringP("title", "t", "", "The title used for services that support it")
}

func logf(format string, a ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", a...)
}

func run(cmd *cobra.Command) error {
	flags := cmd.Flags()
	verbose, _ := flags.GetBool("verbose")

	urls, _ := flags.GetStringArray("url")
	urls = dedupe.RemoveDuplicates(urls)
	message, _ := flags.GetString("message")
	title, _ := flags.GetString("title")

	if message == "-" {
		logf("Reading from STDIN...")

		stringBuilder := strings.Builder{}

		count, err := io.Copy(&stringBuilder, os.Stdin)
		if err != nil {
			return fmt.Errorf("failed to read message from stdin: %w", err)
		}

		logf("Read %d byte(s)", count)

		message = stringBuilder.String()
	}

	var logger *log.Logger

	if verbose {
		urlsPrefix := "URLs:"
		for i, url := range urls {
			logf("%s %s", urlsPrefix, url)

			if i == 0 {
				// Only display "URLs:" prefix for first line, replace with indentation for the subsequent
				urlsPrefix = strings.Repeat(" ", len(urlsPrefix))
			}
		}

		logf("Message: %s", util.Ellipsis(message, MaxMessageLength))

		if title != "" {
			logf("Title: %v", title)
		}

		logger = log.New(os.Stderr, "SHOUTRRR ", log.LstdFlags)
	} else {
		logger = util.DiscardLogger
	}

	serviceRouter, err := router.New(logger, urls...)
	if err != nil {
		return cli.ConfigurationError(fmt.Sprintf("error invoking send: %s", err))
	}

	params := make(types.Params)
	if title != "" {
		params["title"] = title
	}

	errs := serviceRouter.SendAsync(message, &params)
	for err := range errs {
		if err != nil {
			return cli.TaskUnavailable(err.Error())
		}

		logf("Notification sent")
	}

	return nil
}

// Run executes the send command and handles its result.
func Run(cmd *cobra.Command, _ []string) error {
	err := run(cmd)
	if err != nil {
		var result cli.Result
		if errors.As(err, &result) && result.ExitCode != cli.ExUsage {
			// If the error is not related to CLI usage, report error and exit to avoid cobra error output
			_, _ = fmt.Fprintln(os.Stderr, err.Error())

			os.Exit(result.ExitCode)
		}
	}

	return err
}
