package botApp

import ("gopkg.in/telegram-bot-api.v4")

func Handle(message tgbotapi.Message) error {
  if message.IsCommand() {
    // handle command
  } else {
    // handle text
  }
}