package image

type Group string

const (
	AllowedImageGroupGeneral Group = "images"
)

func Groups() []string {
	return []string{
		string(AllowedImageGroupGeneral),
	}
}
