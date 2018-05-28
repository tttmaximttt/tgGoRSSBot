package botApp

import (
  "github.com/tttmaximttt/teleGoBot/config"
  "github.com/tttmaximttt/teleGoBot/models"
  "gopkg.in/telegram-bot-api.v4"
  log "github.com/sirupsen/logrus"
)

var rssMap = map[string]string{
  "Habr": "https://habrahabr.ru/rss/hubs/all",
}

func New(config config.Config) (*tgbotapi.BotAPI, error) {
  bot, err := tgbotapi.NewBotAPI(config.BotToken)
  _, err = bot.SetWebhook(tgbotapi.NewWebhook(config.WebhookURL))
  return bot, err
}

func Run(bot tgbotapi.BotAPI, updates tgbotapi.UpdatesChannel) {
  for update := range updates {
    if url, ok := rssMap[update.Message.Text]; ok {

      rss, err := models.GetRSS(update.Message.Text)

      news, err := rss.GetNews(url)

      if err != nil {
        bot.Send(tgbotapi.NewMessage(
          update.Message.Chat.ID,
          "sorry, error happend " + err.Error(),
        ))
        log.Error("Error hapend %s", err.Error())
      }


      if rss != nil {
        for _, item := range news.Items {
          bot.Send(tgbotapi.NewMessage(
            update.Message.Chat.ID,
            item.Url+"\n"+item.Title,
          ))
        }
      }
    } else {
      bot.Send(tgbotapi.NewMessage(
        update.Message.Chat.ID,
        "there is only Habr feed availible",
      ))
    }
  }
}
