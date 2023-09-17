package main

import (
	"os"

	"github.com/deltachat-bot/deltabot-cli-go/botcli"
	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
	"github.com/deltachat/deltachat-rpc-client-go/deltachat/option"
	"github.com/spf13/cobra"
)

var cli = botcli.New("sipbot")

func onBotInit(cli *botcli.BotCli, bot *deltachat.Bot, cmd *cobra.Command, args []string) {
	bot.OnNewMsg(onNewMsg)

	// set message auto-deletion for to cleanup
	accounts, _ := bot.Rpc.GetAllAccountIds()
	for _, accId := range accounts {
		isConf, err := bot.Rpc.IsConfigured(accId)
		if isConf || err != nil {
			continue
		}

		err = bot.Rpc.SetConfig(accId, "delete_device_after", option.Some("86400")) // one day
		if err != nil {
			cli.Logger.Error(err)
		}
		err = bot.Rpc.SetConfig(accId, "delete_server_after", option.Some("1"))
		if err != nil {
			cli.Logger.Error(err)
		}
		err = bot.Rpc.SetConfig(accId, "displayname", option.Some("SIPBot"))
		if err != nil {
			cli.Logger.Error(err)
		}
	}
}

func onBotStart(cli *botcli.BotCli, bot *deltachat.Bot, cmd *cobra.Command, args []string) {
	dsn := os.Getenv("SIPBOT_DBDSN")
	domain := os.Getenv("SIPBOT_DOMAIN")
	if dsn == "" {
		cli.Logger.Fatal("SIPBOT_DBDSN env var is not set")
	}
	if domain == "" {
		cli.Logger.Fatal("SIPBOT_DOMAIN env var is not set")
	}
	if err := initDB(dsn, domain); err != nil {
		cli.Logger.Error(err)
		bot.Stop()
	}
}

func main() {
	cli.OnBotInit(onBotInit)
	cli.OnBotStart(onBotStart)
	if err := cli.Start(); err != nil {
		cli.Logger.Error(err)
	}
}
