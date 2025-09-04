package masker

import "os"

type Service struct {
	prod producer
	pres presenter
}

func NewService(producer, presenter) *Service {
	return &Service{
		prod: NewDataProducer(getPathFromArgs()),
	}
}

func getPathFromArgs() string {
	return os.Args[1]
}

func findTheLastSpaceOrStop(message []byte, startIndex int) int {
	for i := startIndex; i < len(message); i++ {
		if i == len(message)-1 {
			return len(message)
		}
		if message[i] == ' ' || i == len(message) {
			return i
		}
	}
	return 0
}

const httpConst string = "http://"

func changeLinkToStars(bufMes []byte, httpBytes []byte) {
	for i := 0; i < len(bufMes); i++ {
		for _, val := range httpBytes {
			if val == bufMes[i] {
				if (!(i+7 > len(bufMes))) && string(bufMes[i:i+7]) == string(httpBytes) {
					for index := i + 7; index < findTheLastSpaceOrStop(bufMes, i+7); index++ {
						bufMes[index] = '*'
					}
				}
			}
		}

	}
}

func (s *Service) CatchHttp(message string) string {
	httpBytes := []byte(httpConst)
	bufMes := []byte(message)
	changeLinkToStars(bufMes, httpBytes)
	return string(bufMes)
}
