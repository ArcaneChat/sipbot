package main

import (
	"testing"

	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestErrResponse(t *testing.T) {
	acfactory.WithOnlineBot(func(bot *deltachat.Bot, botAcc deltachat.AccountId) {
		acfactory.WithOnlineAccount(func(userRpc *deltachat.Rpc, userAcc deltachat.AccountId) {
			bot.OnNewMsg(onNewMsg)
			go bot.Run() //nolint:errcheck

			chatWithBot := acfactory.CreateChat(userRpc, userAcc, bot.Rpc, botAcc)

			_, err := userRpc.MiscSendTextMessage(userAcc, chatWithBot, "hi")
			require.Nil(t, err)

			msg := acfactory.NextMsg(userRpc, userAcc)
			assert.JSONEq(t, "{\"error\":\"Unknown command\"}", msg.Text)
		})
	})
}
