package masker

import (
	"fmt"
	"os"
	"strings"
)

type Service struct {
	prod producer
	pres presenter
}

func NewService() *Service {
	return &Service{
		prod: NewDataProducer(getPathFromArgs()),
		pres: NewDataPresenter("maskedStrings.txt"),
	}
}

func (s *Service) Run() {
	spam, err := s.prod.produce()
	if err != nil {
		fmt.Println(err.Error())
	}
	result := strings.Join(spam, "")
	maskedStrings := make([]string, 0)
	maskedStrings = append(maskedStrings, s.CatchHttp(result))

	fmt.Println(maskedStrings)
	s.pres.present(maskedStrings)
}

func getPathFromArgs() string {
	if len(os.Args) > 1 {
		return os.Args[1]
	}
	return "stringsToMask.txt"
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
