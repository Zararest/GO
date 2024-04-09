# Task description
In this lecture xml parser should be implemented.

Program should:
- parse yaml file with input/output files info
- parse input file and sort it by value filed

## Notes
- Yaml parers are based on structure tags:
```go
type _YAMLStructure struct {
	InputFile string `yaml:"input-file"`
	OutFile   string `yaml:"output-file"`
}
```
- In order to create one representation fro manother you should use:
```go
json.Marshal(person)
```
- Imports are related to modules, not the directories:
```go
*module-name*/path/to/dir
```