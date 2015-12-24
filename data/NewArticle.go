package data

import "fmt"

type NewArticle struct {
    Title string
}

func NewArticleData() NewArticle {
  return NewArticle{fmt.Sprintf("New Article - %s", conf["title"])}
}
