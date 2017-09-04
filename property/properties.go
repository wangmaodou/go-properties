package property

import (
	"os"
	"bufio"
	"strings"
)

type Properties struct {
	value map[string]string
	path  string
}

func FromFile(path string) *Properties {
	result := &Properties{}
	if len(path) == 0 {
		panic("The path is null.")
	}
	result.path = path
	file, err := os.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	result.value = getProperties(file)
	return result

}

func getProperties(file *os.File) map[string]string {
	reader := bufio.NewReader(file)
	result := make(map[string]string)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			panic(err)
			break
		}
		s := string(line)
		strings.Replace(s, " ", "", -1)
		if len(s) == 0 || strings.Index(s, "#") == 0 {
			continue
		}
		kv := strings.Split(s, "=")
		if len(kv) < 2 {
			continue
		}
		result[kv[0]] = kv[1]
	}
	return result
}

func (p Properties) Get(key string) string {
	v, _ := p.value[key]
	return v
}

func (p *Properties) Set(k, v string) {
	p.value[k] = v
}
