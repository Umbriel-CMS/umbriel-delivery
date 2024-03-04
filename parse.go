package main

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
	BlockPosition int         `json:"blockPosition"`
	BlockType     string      `json:"blockType"`
	BlocksData    ContentData `json:"blocksData"`
}
type PageBlock struct {
	BlockPosition int       `json:"blockPosition"`
	PageId        string    `json:"pageId"`
	BlockType     string    `json:"blockType"`
	BlocksData    BlockData `json:"blocksData"`
}

type SimplifiedPageBlock struct {
	BlocksData SimplifiedBlockData `json:"blocksData"`
}

type SimplifiedBlockData struct {
	BlockPosition int          `json:"blockPosition"`
	BlockType     string       `json:"blockType"`
	Layout        string       `json:"layout"`
	Template      string       `json:"template"`
	MenuItems     []MenuItem   `json:"menuItems,omitempty"`
	Content       *ContentData `json:"content,omitempty"`
}

type ContentData struct {
	Layout   string `json:"layout"`
	Template string `json:"template"`
	Items    []Item `json:"items,omitempty"`
}

type Item struct {
	CarrouselImagePath string `json:"carrousel_image_path"`
	CarrouselImageAlt  string `json:"carrousel_image_alt"`
}

func transformToSimplifiedBlocks(pageBlocks []PageBlock) []SimplifiedPageBlock {
	var simplifiedBlocks []SimplifiedPageBlock
	for _, block := range pageBlocks {
		var items []Item
		var contentData *ContentData

		if block.BlocksData.Content != nil {
			for _, item := range block.BlocksData.Content.BlocksData.Items {
				items = append(items, Item{
					CarrouselImagePath: item.CarrouselImagePath,
					CarrouselImageAlt:  item.CarrouselImageAlt,
				})
			}
			contentData = &ContentData{
				Layout:   block.BlocksData.Content.BlocksData.Layout,
				Template: block.BlocksData.Content.BlocksData.Template,
				Items:    items,
			}
		}

		simplifiedBlock := SimplifiedPageBlock{
			BlocksData: SimplifiedBlockData{
				BlockPosition: block.BlockPosition,
				BlockType:     block.BlockType,
				Layout:        block.BlocksData.Layout,
				Template:      block.BlocksData.Template,
				MenuItems:     block.BlocksData.MenuItems,
				Content:       contentData,
			},
		}
		simplifiedBlocks = append(simplifiedBlocks, simplifiedBlock)
	}
	return simplifiedBlocks
}
