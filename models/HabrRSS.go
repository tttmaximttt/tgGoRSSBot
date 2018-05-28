package models

import (
  "net/http"
  "io/ioutil"
  "encoding/xml"
  "github.com/sirupsen/logrus"
  "fmt"
)

type item struct {
  Url string `xml:"guid"`
  Title string `xml:"title"`
}

type rss struct {
  Items []item `xml:"channel>item"`
  Core
}

func (self *rss) GetNews(url string) (*rss, error) {
  resp, err := http.Get(url)

  if err != nil {
    return nil, err
  }

  defer resp.Body.Close()

  logrus.Info(fmt.Sprintf("%s %v", url, resp.StatusCode))

  body, _ := ioutil.ReadAll(resp.Body)

  err = xml.Unmarshal(body, self)
  if err != nil {
    return nil, err
  }

  return self, nil
}
