package main

func transformCenterMiddle(centerMiddle *CenterMiddle) *CenterMiddle {
	transformedItems := make([]CenterMiddleItem, len(centerMiddle.Items))
	for i, item := range centerMiddle.Items {
		transformedItems[i] = CenterMiddleItem{
			SlotRightItems: SlotItems{
				Content: item.SlotRightItems.Content,
			},
			SlotLeftItems: SlotItems{
				Content: item.SlotLeftItems.Content,
			},
		}
	}

	return &CenterMiddle{
		BlockPosition:     centerMiddle.BlockPosition,
		Layout:            centerMiddle.Layout,
		Template:          centerMiddle.Template,
		Target:            centerMiddle.Target,
		Items:             transformedItems,
		Template50Options: centerMiddle.Template50Options,
		TextProps:         centerMiddle.TextProps,
	}
}

func transformToOutputBlockData(inputData InputBlockData) OutputBlockData {
	layoutCarrousel := make([]Image, 0, len(inputData.BlocksData.LayoutCarousel))
	for _, image := range inputData.BlocksData.LayoutCarousel {
		transformedImage := Image{
			ImageMiddleDesktopPath: image.ImageMiddleDesktopPath,
			ImageMobilePath:        image.ImageMobilePath,
		}
		layoutCarrousel = append(layoutCarrousel, transformedImage)
	}

	centerMiddle := transformCenterMiddle(&inputData.BlocksData.CenterMiddle)

	return OutputBlockData{
		BlockType: inputData.BlockType,
		BlocksData: OutputBlocksData{
			Layout:          inputData.BlocksData.Layout,
			Template:        inputData.BlocksData.Template,
			Menu:            inputData.BlocksData.Menu,
			Footer:          inputData.BlocksData.Footer,
			LayoutCarrousel: layoutCarrousel,
			CenterMiddle:    *centerMiddle,
		},
	}
}
