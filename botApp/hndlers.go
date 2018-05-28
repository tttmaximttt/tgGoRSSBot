package botApp

import ("gopkg.in/telegram-bot-api.v4")

func Handle(message tgbotapi.Message) {
  if message.IsCommand() {
    // handle command
  } else {
    // handle text
  }
}