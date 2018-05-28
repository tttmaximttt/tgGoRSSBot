package models

import (
  "fmt"
  "github.com/sirupsen/logrus"
)

func GetRSS (typeRSS string) (Core, error) {
  logrus.Info(fmt.Sprintf("Requested RSS %s type", typeRSS))

  switch typeRSS {
  case "Habr":
    return new(rss), nil
  default:
    return  nil, fmt.Errorf("RSS type %s not found ", typeRSS)
  }
}

type Core interface {
  GetNews(url string) (*rss, error)
}