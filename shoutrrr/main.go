package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/circa10a/shoutrrr/internal/meta"
	"github.com/circa10a/shoutrrr/shoutrrr/cmd"
	"github.com/circa10a/shoutrrr/shoutrrr/cmd/docs"
	"github.com/circa10a/shoutrrr/shoutrrr/cmd/generate"
	"github.com/circa10a/shoutrrr/shoutrrr/cmd/send"
	"github.com/circa10a/shoutrrr/shoutrrr/cmd/verify"
)

var cobraCmd = &cobra.Command{
	Use:   "shoutrrr",
	Short: "Shoutrrr CLI",
}

func init() {
	viper.AutomaticEnv()
	cobraCmd.AddCommand(verify.Cmd)
	cobraCmd.AddCommand(generate.Cmd)
	cobraCmd.AddCommand(send.Cmd)
	cobraCmd.AddCommand(docs.Cmd)

	cobraCmd.Version = meta.GetMetaStr()
}

func main() {
	if err := cobraCmd.Execute(); err != nil {
		os.Exit(cmd.ExUsage)
	}
}
