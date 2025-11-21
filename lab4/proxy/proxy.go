package proxy

import "fmt"

// Интерфейс субъекта
type Image interface {
	Display() string
}

// Реальный субъект
type RealImage struct {
	Filename string
}

func NewRealImage(filename string) *RealImage {
	image := &RealImage{Filename: filename}
	image.loadFromDisk()
	return image
}

func (r *RealImage) Display() string {
	return fmt.Sprintf("Отображение изображения: %s", r.Filename)
}

func (r *RealImage) loadFromDisk() {
	fmt.Printf("Загрузка изображения %s с диска...\n", r.Filename)
}

// Прокси
type ImageProxy struct {
	realImage *RealImage
	Filename  string
}

func NewImageProxy(filename string) *ImageProxy {
	return &ImageProxy{Filename: filename}
}

func (p *ImageProxy) Display() string {
	if p.realImage == nil {
		p.realImage = NewRealImage(p.Filename)
	}
	return p.realImage.Display()
}