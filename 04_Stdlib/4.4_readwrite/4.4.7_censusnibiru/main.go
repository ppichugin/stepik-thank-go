package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
)

// алфавит планеты Нибиру
const alphabet = "aeiourtnsl"

// Census реализует перепись населения.
// Записи о рептилоидах хранятся в каталоге census, в отдельных файлах,
// по одному файлу на каждую букву алфавита.
// В каждом файле перечислены рептилоиды, чьи имена начинаются
// на соответствующую букву, по одному рептилоиду на строку.
type Census struct {
	dir string
	fds map[rune]*os.File
}

// Count возвращает общее количество переписанных рептилоидов.
func (c *Census) Count() int {
	var count int
	for _, fd := range c.fds {
		info, err := fd.Stat()
		if err != nil {
			panic(err)
		}
		count += int(info.Size() / 2)
	}
	return count
}

// Add записывает сведения о рептилоиде.
func (c *Census) Add(name string) {
	if len(name) == 0 || name[0] < 'a' || name[0] > 'z' {
		return
	}
	r := rune(name[0])
	if _, ok := c.fds[r]; !ok {
		fd, err := os.OpenFile(filepath.Join(c.dir, string(r)+".txt"), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		c.fds[r] = fd
	}
	fd := c.fds[r]
	_, err := fd.WriteString(name + "\n")
	if err != nil {
		panic(err)
	}
}

// Close закрывает файлы, использованные переписью.
func (c *Census) Close() {
	for _, fd := range c.fds {
		err := fd.Close()
		if err != nil {
			panic(err)
		}
	}
}

// NewCensus создает новую перепись и пустые файлы
// для будущих записей о населении.
func NewCensus() *Census {
	dir := "census"
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		panic(err)
	}
	fds := make(map[rune]*os.File)
	for _, r := range alphabet {
		fd, err := os.OpenFile(filepath.Join(dir, string(r)+".txt"), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		fds[r] = fd
	}
	return &Census{
		dir: dir,
		fds: fds,
	}
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// randomName возвращает имя очередного рептилоида.
func randomName(n int) string {
	chars := make([]byte, n)
	for i := range chars {
		chars[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(chars)
}

func main() {
	rand.Seed(0)
	census := NewCensus()
	defer census.Close()
	for i := 0; i < 1024; i++ {
		reptoid := randomName(5)
		census.Add(reptoid)
	}
	fmt.Println(census.Count())
	ReptoidByName(census, 'n', 42)
}

func ReptoidByName(c *Census, name rune, order int) {
	m := c.fds
	n := m[name]

	file, err := os.Open(n.Name())
	if err != nil {
		return
	}

	var count int
	var reptoid string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		word := scanner.Text()
		count++
		if count == order {
			reptoid = word
		}
	}

	fmt.Println(count, reptoid)
}
