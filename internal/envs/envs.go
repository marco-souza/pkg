package envs

import (
	"fmt"
	"os"
	"strings"
)

type Env struct {
	Filepath string
}

func NewEnv(filepath string) *Env {
	return &Env{Filepath: filepath}
}

func (e *Env) GetEnv(key string) (string, error) {
	lines, err := e.readFile()
	if err != nil {
		return "", err
	}

	for _, line := range lines {
		parsed := strings.Split(string(line), "=")
		if len(parsed) != 2 {
			continue
		}

		if parsed[0] == key {
			value := strings.Trim(parsed[1], "'")
			return value, nil
		}
	}

	return "", fmt.Errorf("key %s not found", key)
}

func (e *Env) SetEnv(key, value string) error {
	lines, err := e.readFile()
	if err != nil {
		return err
	}

	output := ""
	found := false
	for _, line := range lines {
		parsed := strings.Split(string(line), "=")
		if len(parsed) != 2 {
			output += string(line) + "\n"
			continue
		}

		if parsed[0] == key {
			parsed[1] = value
			found = true
		}

		output += fmt.Sprintf("%s='%s'\n", parsed[0], parsed[1])
	}

	if !found {
		output += fmt.Sprintf("%s='%s'\n", key, value)
	}

	fmt.Println(output)

	if err := os.WriteFile(e.Filepath, []byte(output), 0644); err != nil {
		return err
	}

	return nil
}

func (e *Env) DetEnv(key string) error {
	lines, err := e.readFile()
	if err != nil {
		return err
	}

	output := ""
	for _, line := range lines {
		parsed := strings.Split(string(line), "=")
		if len(parsed) != 2 {
			output += string(line) + "\n"
			continue
		}

		if parsed[0] == key {
			continue
		}

		output += fmt.Sprintf("%s='%s'\n", parsed[0], parsed[1])
	}

	if err := os.WriteFile(e.Filepath, []byte(output), 0644); err != nil {
		return err
	}

	return nil
}

func (e *Env) GenerateTemplate() error {
	lines, err := e.readFile()
	if err != nil {
		return err
	}

	output := ""
	for _, line := range lines {
		parsed := strings.Split(string(line), "=")
		if len(parsed) != 2 {
			output += string(line) + "\n"
			continue
		}

		output += fmt.Sprintf("%s='%s'\n", parsed[0], "********")
	}

	outputFile := e.Filepath + ".template"
	os.WriteFile(outputFile, []byte(output), 0644)

	return nil
}

func (e *Env) readFile() ([]string, error) {
	lines, err := os.ReadFile(e.Filepath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(lines), "\n"), nil
}

func init() {
	fmt.Println("initializing package")
}
