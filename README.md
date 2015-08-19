# go-config
When working with microservice, Docker and go, I experienced a really ununified way for configuration.

## Example
```
import (
  config "github.com/nikriek/go-config"
)

func main() {

  // Provide a list of required variables
  c, err := c.Requires([]]string{"DB_HOST", "PUBLIC_KEY"})
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