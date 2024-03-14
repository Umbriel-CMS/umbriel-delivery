package main

type InputBlockData struct {
	BlockType  string          `json:"blockType"`
	BlocksData InputBlocksData `json:"blocksData"`
}
type SlotImages struct {
	SlotLeftImages  []Image `json:"slot_left_images,omitempty"`
	SlotRightImages []Image `json:"slot_right_images,omitempty"`
}
type InputBlocksData struct {
	Layout         string       `json:"layout"`
	Template       string       `json:"template"`
	Menu           []MenuItem   `json:"menu"`
	Footer         []Footer     `json:"footer"`
	LayoutCarousel []Image      `json:"layoutCarrousel"`
	CenterMiddle   CenterMiddle `json:"centerMiddle"`
}

type MenuItem struct {
	Title string `json:"title"`
	Href  string `json:"href"`
}

type Footer struct {
	TitleTop    string `json:"title_top"`
	TitleBottom string `json:"title_bottom"`
}

type Image struct {
	ImageMiddleDesktopPath string `json:"image_middle_desktop_path"`
	ImageMobilePath        string `json:"image_mobile_path"`
}

type CenterMiddle struct {
	BlockPosition     string              `json:"blockPosition"`
	Layout            string              `json:"layout"`
	Template          string              `json:"template"`
	Target            string              `json:"target"`
	TextProps         []TextPropsStruct   `json:"textProps"`
	Template50Options []Template50Options `json:"template50Options"`
	Items             []CenterMiddleItem  `json:"items"`
}

type TextPropsStruct struct {
	Color      string `json:"color,omitempty"`
	FontFamily string `json:"fontFamily,omitempty"`
	FontSize   string `json:"fontSize,omitempty"`
	Transform  string `json:"transform,omitempty"`
	LineHeight string `json:"lineHeight,omitempty"`
	As         string `json:"as,omitempty"`
}

type Template50Options struct {
	SlotLeftBgColor    string               `json:"slot_left_bgColor,omitempty"`
	SlotRightBgColor   string               `json:"slot_right_bgColor,omitempty"`
	AlignTextSlotLeft  string               `json:"align_text_slot_left,omitempty"`
	AlignTextSlotRight string               `json:"align_text_slot_right,omitempty"`
	AnchorHandlerProps []AnchorHandlerProps `json:"anchorHandlerProps,omitempty"`
}

type AnchorHandlerProps struct {
	IsSlotImageLeftAnchor  bool   `json:"isSlotImageLeftAnchor"`
	IsSlotImageRightAnchor bool   `json:"isSlotImageRightAnchor"`
	IsSlotLeftTitleAnchor  bool   `json:"isSlotLeftTitleAnchor"`
	IsSlotRightTitleAnchor bool   `json:"isSlotRightTitleAnchor"`
	Href                   string `json:"href"`
	Index                  string `json:"index"`
}

type CenterMiddleItem struct {
	SlotRightItems SlotItems `json:"slot_right_items"`
	SlotLeftItems  SlotItems `json:"slot_left_items"`
}

type SlotItems struct {
	Content []ContentItem `json:"content"`
}

type ContentItem struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Image    any    `json:"image"`
}
type ContentItems struct {
	Content []Content `json:"content"`
}

type SlotLeftItems struct {
	ContentItems
	Images []Images `json:"images"`
}

type Content struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Image    any    `json:"image"`
}

type Images struct {
	SlotLeftImages  []Image `json:"slot_left_images"`
	SlotRightImages []Image `json:"slot_right_images"`
}

type OutputBlockData struct {
	BlockPosition string           `json:"blockPosition"`
	BlockType     string           `json:"blockType"`
	Target        string           `json:"target"`
	BlocksData    OutputBlocksData `json:"blocksData"`
}

type OutputBlocksData struct {
	Layout          string       `json:"layout"`
	Template        string       `json:"template"`
	Target          string       `json:"target"`
	Menu            []MenuItem   `json:"menu"`
	Footer          []Footer     `json:"footer"`
	LayoutCarrousel []Image      `json:"layoutCarrousel"`
	CenterMiddle    CenterMiddle `json:"centerMiddle"`
}
