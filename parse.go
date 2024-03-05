package main

type BlockData struct {
	Layout       string        `json:"layout"`
	Template     string        `json:"template"`
	MenuItems    []MenuItem    `json:"menuItems"`
	Content      *Content      `json:"content"`
	CenterMiddle *CenterMiddle `json:"centerMiddle"`
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
	PageId     string    `json:"pageId"`
	BlockType  string    `json:"blockType"`
	BlocksData BlockData `json:"blocksData"`
}

type SimplifiedPageBlock struct {
	BlocksData SimplifiedBlockData `json:"blocksData"`
}

type SimplifiedBlockData struct {
	BlockPosition int           `json:"blockPosition"`
	BlockType     string        `json:"blockType"`
	Layout        string        `json:"layout"`
	Template      string        `json:"template"`
	MenuItems     []MenuItem    `json:"menuItems,omitempty"`
	Content       *ContentData  `json:"content,omitempty"`
	CenterMiddle  *CenterMiddle `json:"centerMiddle,omitempty"`
}

type ContentData struct {
	Layout   string `json:"layout"`
	Template string `json:"template"`
	Items    []Item `json:"items,omitempty"`
}

type CenterMiddle struct {
	BlockPosition int                `json:"blockPosition"`
	Layout        string             `json:"layout"`
	Template      string             `json:"template"`
	Items         []CenterMiddleItem `json:"items,omitempty"`
}

type CenterMiddleItem struct {
	Title    string               `json:"title,omitempty"`
	Subtitle string               `json:"subtitle,omitempty"`
	Images   []CenterMiddleImage  `json:"images,omitempty"`
	Footer   []CenterMiddleFooter `json:"footer,omitempty"`
}

type CenterMiddleImage struct {
	ImageMiddleDesktopPath string `json:"image_middle_desktop_path,omitempty"`
	ImageMobilePath        string `json:"image_mobile_path,omitempty"`
}

type CenterMiddleFooter struct {
	TitleTop    string `json:"title_top,omitempty"`
	TitleBottom string `json:"title_bottom,omitempty"`
}

type Item struct {
	CarrouselImagePath string `json:"carrousel_image_path"`
	CarrouselImageAlt  string `json:"carrousel_image_alt"`
}

func transformCenterMiddle(centerMiddle *CenterMiddle) *CenterMiddle {
	items := transformCenterMiddleItems(centerMiddle.Items)
	return &CenterMiddle{
		BlockPosition: centerMiddle.BlockPosition,
		Layout:        centerMiddle.Layout,
		Template:      centerMiddle.Template,
		Items:         items,
	}
}

func transformToSimplifiedBlocks(pageBlocks []PageBlock) []SimplifiedPageBlock {
	var simplifiedBlocks []SimplifiedPageBlock

	for _, block := range pageBlocks {
		var contentData *ContentData
		var centerMiddle *CenterMiddle
		if block.BlocksData.CenterMiddle != nil {
			centerMiddle = transformCenterMiddle(block.BlocksData.CenterMiddle)
		}

		simplifiedBlock := SimplifiedPageBlock{
			BlocksData: SimplifiedBlockData{
				BlockType:    block.BlockType,
				Layout:       block.BlocksData.Layout,
				Template:     block.BlocksData.Template,
				MenuItems:    block.BlocksData.MenuItems,
				Content:      contentData,
				CenterMiddle: centerMiddle,
			},
		}
		simplifiedBlocks = append(simplifiedBlocks, simplifiedBlock)
	}
	return simplifiedBlocks
}

func transformCenterMiddleItems(items []CenterMiddleItem) []CenterMiddleItem {
	var transformedItems []CenterMiddleItem
	for _, item := range items {
		transformedItem := CenterMiddleItem{
			Title:    item.Title,
			Subtitle: item.Subtitle,
			Images:   item.Images,
			Footer:   item.Footer,
		}
		transformedItems = append(transformedItems, transformedItem)
	}
	return transformedItems
}
