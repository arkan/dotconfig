# dotconfig

Package dotconfig provides an easy way to load/save configuration to the ~/.config/APP_NAME/config.yml file.

[![GoDoc](https://godoc.org/github.com/arkan/dotconfig?status.svg)](https://godoc.org/github.com/arkan/dotconfig)
[![Go Report Card](https://goreportcard.com/badge/github.com/arkan/dotconfig)](https://goreportcard.com/report/github.com/arkan/dotconfig)


## Usage
```go
type settings struct {
	Name     string `yaml:"name"`
	Endpoint string `yaml:"endpoint"`
	Token    string `yaml:"token"`
}

ss := settings{}

if err := dotconfig.Load("my-app", &ss); err != nil {
  if err == dotconfig.ErrConfigNotFound {
    ss.Name = "Name"
    ss.Endpoint = "http://google.com"
    ss.Token = "Some token"
    if err := dotconfig.Save("my-app", ss); err != nil {
      panic(err)
    }
  }
}else if err != nil {
  panic(err)
}

fmt.Printf("Hello %v\n", ss.Name)

```


See the [documentation](https://godoc.org/github.com/arkan/dotconfig) for information.

## Licence
[MIT](./LICENSE)
