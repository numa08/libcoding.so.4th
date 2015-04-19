package main
import (
    "net/http"
    "io/ioutil"
    "github.com/bitly/go-simplejson"
    "strings"
    "golang.org/x/net/html"
    "github.com/PuerkitoBio/goquery"
)

type Performer struct {
    Name string
}

func LoadPerformers() ([]Performer, error){
    url := "http://connpass.com/api/v1/event/?event_id=13232"
    resp, err := http.Get(url)
    defer resp.Body.Close()

    if err != nil {
        return nil, err
    }
    rawBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    jsBody, err := simplejson.NewJson(rawBody)
    if err != nil {
        return nil, err
    }

    description, err := jsBody.Get("events").GetIndex(0).Get("description").String()
    if err != nil {
        return nil, err
    }
    reader := strings.NewReader(description)
    node, err := html.Parse(reader)
    if err != nil {
        return nil, err
    }
    document := goquery.NewDocumentFromNode(node)

    performers := make([]Performer, 4)
    document.Find("h3>a").Each(func(idx int, s *goquery.Selection){
        name := s.Text()
        p := Performer{Name: name,}
        performers[idx] = p
    })
    return performers, nil
}