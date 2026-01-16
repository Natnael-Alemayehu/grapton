package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {

	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gator")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)

	var rssfeed RSSFeed
	if err := xml.Unmarshal(data, &rssfeed); err != nil {
		return nil, err
	}

	rssfeed.Channel.Title = html.UnescapeString(rssfeed.Channel.Title)
	rssfeed.Channel.Description = html.UnescapeString(rssfeed.Channel.Description)

	for i, item := range rssfeed.Channel.Item {
		item.Description = html.UnescapeString(item.Description)
		item.Title = html.UnescapeString(item.Title)
		rssfeed.Channel.Item[i] = item
	}

	return &rssfeed, nil
}
