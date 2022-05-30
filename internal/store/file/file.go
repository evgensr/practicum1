package file

import (
	"bufio"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"github.com/evgensr/practicum1/internal/store"
	"log"
	"os"
	"sync"
)

type Box struct {
	sync.RWMutex
	Items           []Line
	fileStoragePath string
}

type Line = store.Line

// RWMap структура Mutex
type RWMap struct {
	sync.RWMutex
	row             map[string]string
	fileStoragePath string
}

type Row struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func New(param string) *Box {

	fileStoragePath := param

	box := &Box{
		fileStoragePath: fileStoragePath,
	}

	// открывам файл
	file, err := os.Open(fileStoragePath)
	// закрываем файл
	defer func() {
		errCloseFile := file.Close()
		if err == nil {
			err = errCloseFile
		}
	}()

	// если файл не найдет, возврощаем пустую box
	if err != nil {
		return box
	}

	// создаем новый сканер
	fileScanner := bufio.NewScanner(file)
	// создаем переменную для сканера из структуры Row
	line := Line{}

	// скаринуем по 1 строчке
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		// в структуру row разбиаем json
		err := json.Unmarshal(fileScanner.Bytes(), &line)
		if err != nil {
			log.Println("err unmarshal ", err)
		}
		// заполняем box
		box.addItem(line)
	}
	spew.Dump(box)

	return box

}

func save(nameFile string, line Line) error {

	file, errFile := os.OpenFile(nameFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if errFile != nil {
		log.Println(errFile)
		return errFile
	}
	data, err := json.Marshal(line)
	if err != nil {
		log.Println(err)
		return err
	}

	// добавляем перенос строки
	data = append(data, '\n')
	_, err = file.Write(data)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil

}

func (box *Box) addItem(item Line) []Line {
	box.Items = append(box.Items, item)
	return box.Items
}

func removeDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}

	return result
}

// fineDuplicate true если находим дубликат
func fineDuplicate(items *Box, str string) bool {

	if len(items.Items) == 0 {
		return false
	}
	for _, item := range items.Items {
		if item.Short == str {
			return true
		}
	}
	return false
}
