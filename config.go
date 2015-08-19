package config

import (
  "os"
  "errors"
  "fmt"
)

type Config struct {
  Envs map[string]string
}

func Requires(envs []string) (Config, error) {
  c := Config{make(map[string]string)}
  for _,index := range envs {
    value := os.Getenv(index)
    if value == "" {
      return c, errors.New(fmt.Sprintf("%v is missing", index))
    } 
    c.Envs[index] = value
  }
  return c, nil
}

func (c Config) Get(s string) string {
  return c.Envs[s]
}
