/*
Copyright © 2025 DANYIL BAZAROV <danyil.bazarov@gmail.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
	telebot "gopkg.in/telebot.v3"
)

var (
	// Telegram bot auth token
	teletoken = os.Getenv("TELEBOT_TOKEN")
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the telegram bot",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started \n", appVersion)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  teletoken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatal(err)
			return
		}

		kbot.Handle(telebot.OnText, func(context telebot.Context) error {
			log.Print(context.Message().Payload, " ", context.Text())
			switch context.Message().Payload {
			case "hello":
				err = context.Send(fmt.Sprintf("Hello, i'm Kbot, %s!", appVersion))
			}
			return err
		})
		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
