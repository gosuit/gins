# GinSuit

GinSuit is a Go library designed to simplify logging and error handling in your Gin web applications. It integrates seamlessly with the e, lec and sl libraries to provide structured logging and context management.

## Installation

```zsh
go get github.com/gosuit/gins
```

## Features

- **Logger Middleware**: Automatically logs request details and duration.
- **Context Management**: Easily retrieve and manage logging contexts within your handlers.
- **Error Handling**: Simplified error logging and JSON responses for different error types.

## Usage

### Initializing the Logger Middleware

```golang
import (
    "github.com/gin-gonic/gin"
    "github.com/gosuit/lec"
    "github.com/gosuit/sl"
    "github.com/gosuit/gins"
)

func main() {
    ctx := lec.New(sl.Default())

    r := gin.New()

    // Use the logger middleware
    r.Use(gins.InitLogger(ctx))

    // Define your routes here...

    r.Run()
}
```

### Retrieving the Logging Context

```golang
func MyHandler(c *gin.Context) {
    ctx := gins.GetCtx(c)
    logger := ctx.Logger()

    logger.Info("This is a log message")
}
```

### Handling Errors

```golang
func MyHandler(c *gin.Context) {
    err := someFunctionThatMightFail()
    if err != nil {
        gins.Abort(c, err)
        return
    }

    c.JSON(200, gin.H{"message": "success"})
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.