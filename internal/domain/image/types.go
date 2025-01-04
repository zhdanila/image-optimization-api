package image

type ImageGroup string

const (
	AllowedImageGroupGeneral ImageGroup = "images"
)

func ImageGroups() []string {
	return []string{
		string(AllowedImageGroupGeneral),
	}
}
