# go-config
When working with microservice, Docker and go, I wanted to have a simple way for my service configuration. 

## Example
```
import (
  config "github.com/nikriek/go-config"
)

func main() {

  // Provide a list of required variables
  c, err := config.Requires([]]string{"DB_HOST", "PUBLIC_KEY"})
  if err != nil{
    panic(err) // "DB_HOST needs to be set"
  }

  // Work with a variable
  host := c.Get("DB_HOST")
  ...
}

```

## Tests?
Who needs tests?