package masker

import "os"

type dataPresenter struct {
	path string
}

func NewDataPresenter(path string) *dataPresenter {
	return &dataPresenter{
		path: path,
	}
}

func (d *dataPresenter) present(data []string) error {
	processedStrings := make([]byte, 0)
	for _, v := range data {
		processedStrings = append(processedStrings, []byte(v)...)

	}
	err := os.WriteFile(d.path, []byte(processedStrings), 0644)
	return err
}
