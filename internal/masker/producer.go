package masker

import "os"

type DataProducer struct {
	Path string
}

func NewDataProducer(path string) *DataProducer {
	return &DataProducer{
		Path: path,
	}
}

func (d *DataProducer) produce() ([]string, error) {
	var stringsOfSpam []string = make([]string, 0)
	data, err := os.ReadFile(d.Path)
	if err != nil {
		panic(err.Error())
	}
	for _, v := range data {
		stringsOfSpam = append(stringsOfSpam, string(v))
	}
	return stringsOfSpam, err
}
