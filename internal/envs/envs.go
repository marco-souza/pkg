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
	return e.processEnvs(ProcessEnvsOpts{
		LookupKey: key,
		ProcessFound: func(k, _ string) string {
			return fmt.Sprintf("%s='%s'\n", key, value)
		},
		OnNotFound: func() string {
			return fmt.Sprintf("%s='%s'\n", key, value)
		},
	})
}

func (e *Env) DetEnv(key string) error {
	return e.processEnvs(ProcessEnvsOpts{
		LookupKey: key,
		ProcessFound: func(k, _ string) string {
			return ""
		},
	})
}

func (e *Env) GenerateExample() error {
	return e.processEnvs(ProcessEnvsOpts{
		OutputFile: e.Filepath + ".example",
		ParseLine: func(key, _ string) string {
			return fmt.Sprintf("%s='%s'\n", key, "********")
		},
	})
}

type ProcessEnvsOpts struct {
	ParseLine    func(key, val string) string
	ProcessFound func(key, val string) string
	OnNotFound   func() string
	OutputFile   string
	LookupKey    string
}

func (e *Env) processEnvs(opts ProcessEnvsOpts) error {
	lines, err := e.readFile()
	if err != nil {
		return err
	}

	if opts.ParseLine == nil {
		opts.ParseLine = func(k, v string) string {
			return fmt.Sprintf("%s='%s'\n", k, strings.Trim(v, "'"))
		}
	}
	if opts.OnNotFound == nil {
		opts.OnNotFound = func() string {
			return ""
		}
	}

	output := ""
	found := false
	for _, line := range lines {
		if strings.Trim(line, " ") == "" {
			continue
		}

		parsed := strings.Split(string(line), "=")
		if len(parsed) != 2 {
			output += string(line) + "\n"
			continue
		}

		key, val := parsed[0], parsed[1]
		if opts.LookupKey != "" && key == opts.LookupKey {
			output += opts.ProcessFound(key, val)
			found = true
			continue
		}

		parsedLine := opts.ParseLine(key, val)
		if parsedLine == "" {
			continue
		}

		output += parsedLine
	}

	if opts.LookupKey != "" && !found {
		output += opts.OnNotFound()
	}

	outputFile := opts.OutputFile
	if outputFile == "" {
		outputFile = e.Filepath
	}

	if err := os.WriteFile(outputFile, []byte(output), 0644); err != nil {
		return err
	}

	return nil
}

func (e *Env) readFile() ([]string, error) {
	lines, err := os.ReadFile(e.Filepath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(lines), "\n"), nil
}
