package main

import "fmt"

/**
享元模式
享元模式从对象中剥离出不发生改变且多个实例需要的重复数据，独立出一个享元，
使多个对象共享，从而节省内存以及减少对象数量。
*/

type ImageFlyweightFactory struct {
	maps map[string]*ImageFlyweight
}

var imageFactory *ImageFlyweightFactory

func GetImageFlyweightFactory() *ImageFlyweightFactory {
	if imageFactory == nil {
		imageFactory = &ImageFlyweightFactory{
			make(map[string]*ImageFlyweight),
		}
	}
	return imageFactory
}

func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	image := f.maps[filename]
	if image == nil {
		image = NewImageFlyweight(filename)
		f.maps[filename] = image
	}

	return image
}

type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	// load image file
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}

type ImageViewer struct {
	*ImageFlyweight
}

func (i *ImageFlyweight) Data() string {
	return i.data
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{image}
}

func (i *ImageViewer) Display() {
	fmt.Println("display ", i.data)
}

func main() {
	fviewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")

	fmt.Println(fviewer1)
	fmt.Println(viewer2)
}
