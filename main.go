package main

import (
  "bufio"
  "fmt"
  "gopkg.in/yaml.v2"
  "log"
  "os"
  "strconv"
)

var (
	scanner *bufio.Scanner
)

type T struct {
	Chords []string `yaml:"chords"`
	Keywords []string `yaml:"keywords"`
	Usages []string `yaml:"usages"`
}

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	fileName := getString("File name")

	t := T{
		Chords: getStringArray("Number of chords", func(i int) string {
			return fmt.Sprintf("Chord %d", i)
		}),
		Keywords: getStringArray("Number of keywords", func(i int) string {
			return fmt.Sprintf("Keyword %d", i)
		}),
		Usages: getStringArray("Number of usages", func(i int) string {
			return fmt.Sprintf("Usage %d", i)
		}),
	}

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	file, err := os.Create(fmt.Sprintf("%s.yaml", fileName))
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	defer file.Close()

	w := bufio.NewWriter(file)
	_, err = w.Write(d)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	
	w.Flush()
}

func getStringArray(msg string, elementMsg func(i int) string) []string {
	arrLength := getInt(msg)
	arr := make([]string, arrLength)
	for i := 0; i < arrLength; i++ {
		element := getString(elementMsg(i))
		arr[i] = element
	}
	
	return arr
}

func getString(msg string) string {
	var text string

	for len(text) == 0 {
		fmt.Println(msg)
		scanner.Scan()
		text = scanner.Text()
	}
	
	return text
}

func getInt(msg string) int {
	for {
		fmt.Println(msg)
		scanner.Scan()
		
		if intValue, err := strconv.Atoi(scanner.Text()); err == nil {
			return intValue
		}
	}
}