package config

import (
	"bufio"
	"encoding/json"
	"io"
	"os"
	"strings"
)

type JsonLoader struct {
}

func (JsonLoader) Load(path string, body interface{}) error {

	file, err := os.Open(path)

	if err != nil {
		return err
	}

	defer file.Close()

	r := bufio.NewReader(file)
	var sb strings.Builder
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		sb.WriteString(line)
	}

	err = json.Unmarshal([]byte(sb.String()), body)

	return err
}
