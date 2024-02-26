package photo

const (
	prefix = "ch3"
)

type Photo struct {
	ID  uint64
	Url string
}

func getPrefix() string {
	return prefix
}

func GetPhotoPrefix(ph Photo) string {
	return prefix + ph.Url
}
