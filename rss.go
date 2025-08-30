package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type rssfeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RssItem `xml:"item"`
	} `xml:"channel"`
}
type RssItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubdate"`
}

func urlToBeFetched(url string) (rssfeed, error) {
	httpClient := http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := httpClient.Get(url)
	if err != nil {
		return rssfeed{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return rssfeed{}, err
	}
	rssfeeds := rssfeed{}

	err = xml.Unmarshal(data, &rssfeeds)
	if err != nil {
		return rssfeed{}, err
	}
	return rssfeeds, nil

}
