package file

import (
	"bufio"
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"log"
	"os"
	"sync"
)

// RWMap структура Mutex
type RWMap struct {
	mutex           sync.RWMutex
	row             map[string]string
	fileStoragePath string
}

type Row struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func New(param string) *RWMap {
	fileStoragePath := param
	// открывам файл
	file, err := os.Open(fileStoragePath)
	// закрываем файл
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

	// если файл не найдет, возврощаем пустую мапу
	if err != nil {
		log.Println(err)
		return &RWMap{
			row:             make(map[string]string),
			fileStoragePath: fileStoragePath,
		}
	}

	// создаем новый сканер
	fileScanner := bufio.NewScanner(file)
	// создаем переменную для сканера из структуры Row
	line := Row{}
	// создаем мапу в которую будем писать
	row := make(map[string]string)

	// скаринуем по 1 строчке
	for fileScanner.Scan() {
		// fmt.Println(fileScanner.Text())
		// в структуру row разбиаем json
		err := json.Unmarshal(fileScanner.Bytes(), &line)
		if err != nil {
			log.Println(err)
		}
		// заполняем мапу
		row[line.Key] = line.Value
	}

	// spew.Dump(row)
	// возврощаем заполненую мапу
	return &RWMap{
		row:             row,
		fileStoragePath: fileStoragePath,
	}
}

// Get is a wrapper for getting the value from the underlying map
func (c *RWMap) Get(key string) string {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.row[key]
}

// Set is a wrapper for setting the value of a key in the underlying map
func (c *RWMap) Set(key string, val string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	line := Row{
		Key:   key,
		Value: val,
	}

	if _, err := c.row[key]; !err {
		c.row[key] = val
		c.save(line)
	}
}

func (c *RWMap) Delete(key string) error {
	return nil
}

func (c *RWMap) Debug() {
	spew.Dump(c.row)
}

// save ...
func (c *RWMap) save(row Row) {

	data, err := json.Marshal(row)

	if err != nil {
		log.Println(err)
	}

	file, errFile := os.OpenFile(c.fileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
	if errFile != nil {
		log.Println(errFile)
	}
	// добавляем перенос строки
	data = append(data, '\n')
	_, err = file.Write(data)
	if err != nil {
		log.Println(err)
	}

}
