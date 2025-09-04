package masker

type producer interface {
	produce() ([]string, error)
}

type presenter interface {
	present([]string) error
}
