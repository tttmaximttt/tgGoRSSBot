package config

import (
  "github.com/kelseyhightower/envconfig"
  "github.com/pkg/errors"
  "os"
  "encoding/json"
  "fmt"
  "path"
)

type Config struct {
  Logger struct{
    LogLevel string `json:"logLevel" envconfig:"LOG_LEVEL"`
  } `json:"logger"`
  BotToken string `json:"botToken" envconfig:"BOT_TOKEN"`
  WebhookURL string `json:"webhookURL" envconfig:"WEBHOOK_URL"`
  Port int `json:"port" envconfig:"PORT" default:"8080"`
}

func LoadConfig(envPrefix string) (*Config, error ){
  env := os.Getenv("APP_ENV")
  if env == "" {
    env = "local"
  }

  configDir := os.Getenv("APP_CONFIG")

  if configDir == "" {
    configDir = "config"
  }

  configPath := path.Join(configDir, fmt.Sprintf("/config.%s.json", env))
  config := new(Config)

  if configPath != "" {
    err := loadFile(configPath, config)
    if err != nil {
      return nil, errors.Wrap(err, "error loading config from file")
    }
  }

  err := envconfig.Process(envPrefix, config)
  if err != nil {
    return nil, errors.Wrap(err, "error loading config from env")
  }

  if config.BotToken == "" {
    fmt.Println(">>> Look at config.example.json <<<")
    panic("Bot token not set")
  }

  return config, nil
}

func loadFile(path string, config interface{}) error {
  configFile, err := os.Open(path)
  if err != nil {
    return errors.Wrap(err, "error to read config file")
  }
  defer configFile.Close()

  decoder := json.NewDecoder(configFile)
  if err = decoder.Decode(config); err != nil {
    return errors.Wrap(err, "failed to decode config file")
  }

  return nil
}