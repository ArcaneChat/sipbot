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
	accounts, _ := bot.Rpc.GetAllAccountIds()
	for _, accId := range accounts {
		isConf, err := bot.Rpc.IsConfigured(accId)
		if isConf || err != nil {
			continue
		}
		bot.Rpc.SetConfig(accId, "delete_device_after", option.Some("86400")) // one day
		bot.Rpc.SetConfig(accId, "delete_server_after", option.Some("1"))
		bot.Rpc.SetConfig(accId, "displayname", option.Some("SIPBot"))
	}

	bot.OnNewMsg(onNewMsg)
}

func onBotStart(cli *botcli.BotCli, bot *deltachat.Bot, cmd *cobra.Command, args []string) {
	dsn := os.Getenv("SIPBOT_DBDSN")
	domain := os.Getenv("SIPBOT_DOMAIN")
	if dsn == "" {
		cli.Logger.Error("SIPBOT_DBDSN env var is not set")
		bot.Stop()
	}
	if domain == "" {
		cli.Logger.Error("SIPBOT_DOMAIN env var is not set")
		bot.Stop()
	}
	if err := initDB(dsn, domain); err != nil {
		cli.Logger.Error(err)
		bot.Stop()
	}
}

func initCli(cli *botcli.BotCli) {
	cli.OnBotInit(onBotInit)
	cli.OnBotStart(onBotStart)
}

func main() {
	if err := cli.Start(); err != nil {
		cli.Logger.Error(err)
	}
}
