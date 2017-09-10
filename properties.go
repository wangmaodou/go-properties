package prop

import (
	"os"
	"bufio"
	"strings"
	"strconv"
)

type Properties struct {
	value map[string]string
	path  string
}

func NewProperties(path string) (result *Properties) {
	result = &Properties{}
	if len(path) == 0 {
		panic("The path is wrong.")
		return
	}
	result.path = path
	file, err := os.OpenFile(path, os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()
	result.value = getProperties(file)
	return
}

func getProperties(file *os.File) map[string]string {
	reader := bufio.NewReader(file)
	result := make(map[string]string)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}
		s := string(line)
		s = strings.Replace(s, " ", "", -1)
		if len(s) == 0 || strings.Index(s, "#") == 0 {
			continue
		}
		kv := strings.Split(s, "=")
		if len(kv) != 2 {
			continue
		}
		result[kv[0]] = kv[1]
	}
	return result
}

func (p Properties) GetString(key string) string {
	v, _ := p.value[key]
	return v
}

func (p Properties) GetInteger(key string) int {
	v, _ := p.value[key]
	i, _ := strconv.Atoi(v)
	return i
}

func (p Properties) GetFloat(key string) float64 {
	v, _ := p.value[key]
	f, _ := strconv.ParseFloat(v, 64)
	return f
}

func (p Properties) GetBool(key string) bool {
	v, _ := p.value[key]
	b, _ := strconv.ParseBool(v)
	return b
}

func (p *Properties) Set(k, v string) {
	p.value[k] = v
}

func (p Properties) String() string {
	result := ""
	result += "value = ["
	for k,v:=range p.value{
		result+=k+"="+v+", "
	}
	result=string([]rune(result)[0:len(result)-2])
	result+="],"
	result += "path = " + p.path
	return result
}
