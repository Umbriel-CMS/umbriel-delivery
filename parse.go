package main

import (
	"encoding/json"
	"io"
)

type BlockData struct {
	Layout    string     `json:"layout"`
	Template  string     `json:"template"`
	MenuItems []MenuItem `json:"menuItems"`
	Content   *Content   `json:"content"`
}

type MenuItem struct {
	Title string `json:"title"`
	Href  string `json:"href"`
}

type Content struct {
	BlockPosition int       `json:"blockPosition"`
	BlockType     string    `json:"blockType"`
	BlocksData    BlockData `json:"blocksData"`
}

type PageBlock struct {
	BlockPosition int       `json:"blockPosition"`
	PageId        string    `json:"pageId"`
	BlockType     string    `json:"blockType"`
	BlocksData    BlockData `json:"blocksData"`
}

func parsePageBlock(body io.Reader) (*PageBlock, error) {
	var block PageBlock
	err := json.NewDecoder(body).Decode(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}
