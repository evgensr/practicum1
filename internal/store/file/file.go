// Package file stores data in file
package file

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/evgensr/practicum1/internal/store"
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

//New init
func New(param string) *Box {

	fileStoragePath := param

	box := &Box{
		fileStoragePath: fileStoragePath,
	}

	// open file
	file, err := os.Open(fileStoragePath)
	// close file
	defer func() {
		errCloseFile := file.Close()
		if err == nil {
			err = errCloseFile
		}
	}()

	// if the file is not found, we return an empty box
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
	// spew.Dump(box)

	return box

}

//save saving the data to a file
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
