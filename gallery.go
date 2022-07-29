package manygo

const (
	ImageAspectRatioHorizontal string = "horizontal"
	ImageAspectRationSquare    string = "square"
)

type GalleryBuilder struct{}

type Gallery struct {
	Elements         []*GalleryCard `json:"elements"`
	imageAspectRatio string         `json:"image_aspect_ration"`
}

type GalleryCard struct {
	Title     string   `json:"title"`
	SubTitle  string   `json:"subtitle"`
	ImageUrl  string   `json:"image_url"`
	ActionUrl string   `json:"action_url,omitempty"`
	Buttons   []Button `json:"buttons,omitempty"`
}

func (builder *GalleryBuilder) NewGallery() *Gallery {
	return &Gallery{
		Elements:         make([]*GalleryCard, 0),
		imageAspectRatio: ImageAspectRatioHorizontal,
	}
}

func (gallery *Gallery) SetImageAspectRation(imageAspectRatio string) {
	gallery.imageAspectRatio = imageAspectRatio
}

func (gallery *Gallery) GetImageAspectRation() string {
	return gallery.imageAspectRatio
}

func (gallery *Gallery) AddGalleryCard(title string, subtitle string, imageUrl string, actionUrl string, buttons []Button) {
	card := &GalleryCard{
		Title:    title,
		SubTitle: subtitle,
		ImageUrl: imageUrl,
	}

	if "" != actionUrl {
		card.ActionUrl = actionUrl
	}

	if len(buttons) > 0 {
		card.Buttons = buttons
	}

	gallery.Elements = append(gallery.Elements, card)
}
