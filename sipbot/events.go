package main

import (
	"encoding/json"

	"github.com/deltachat/deltachat-rpc-client-go/deltachat"
)

type ErrResponse struct {
	Error string `json:"error"`
}

func onNewMsg(bot *deltachat.Bot, accId deltachat.AccountId, msgId deltachat.MsgId) {
	logger := cli.GetLogger(accId).With("msg", msgId)

	msg, err := bot.Rpc.GetMessage(accId, msgId)
	if err != nil {
		logger.Error(err)
		return
	}

	if msg.IsBot || msg.FromId <= deltachat.ContactLastSpecial {
		return
	}

	chat, err := bot.Rpc.GetBasicChatInfo(accId, msg.ChatId)
	if err != nil {
		logger.Error(err)
		return
	}

	if chat.ChatType == deltachat.ChatSingle {
		if msg.Text == "/register" {
			onRegister(bot.Rpc, accId, msg.ChatId, msg.Sender.Address)
		} else {
			logger.Debugf("Got new unsupported 1:1 message: %#v", msg)
			reportError(bot.Rpc, accId, msg.ChatId)
		}
	}
}

// register in the SIP server
func onRegister(rpc *deltachat.Rpc, accId deltachat.AccountId, chatId deltachat.ChatId, address string) {
	logger := cli.GetLogger(accId).With("chat", chatId)
	sipAcc, err := getAccount(address)
	if err != nil {
		logger.Error(err)
		return
	}
	data, err := json.Marshal(sipAcc)
	if err != nil {
		logger.Error(err)
		return
	}
	_, err = rpc.SendMsg(accId, chatId, deltachat.MsgData{Text: string(data)})
	if err != nil {
		logger.Error(err)
		return
	}
}

// Send an error message if command is unknown
func reportError(rpc *deltachat.Rpc, accId deltachat.AccountId, chatId deltachat.ChatId) {
	logger := cli.GetLogger(accId).With("chat", chatId)
	data, err := json.Marshal(ErrResponse{Error: "Unknown command"})
	if err != nil {
		logger.Error(err)
		return
	}
	_, err = rpc.SendMsg(accId, chatId, deltachat.MsgData{Text: string(data)})
	if err != nil {
		logger.Error(err)
		return
	}
}
