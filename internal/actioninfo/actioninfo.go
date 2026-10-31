package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает разбор данных активности и формирование отчета.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info выводит отчеты об активностях, пропуская записи с ошибками.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("actioninfo: Info: не удалось разобрать данные %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("actioninfo: Info: не удалось сформировать отчет для %q: %v", data, err)
			continue
		}
		fmt.Println(info)
	}
}
