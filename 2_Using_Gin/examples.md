# Gin

Gin is one of the most popular web frameworks for golang.

We use it often in repos at Codecademy (it even ships with the [go service cookie-cutter template](https://github.com/codecademy-engineering/cookiecutter-service-go)).

Creating a Gin router:

```go
r := gin.Default()
```

Handler funcs for Gin take a pointer to a gin 'Context' which is a simplified way of using the ResponseWriter and request pointer from standard lib.

Additionally, the Gin context includes convenient helpers such as `c.JSON` which makes rendering JSON responses very simple.

```go
func myHandler(c *gin.Context) {
	myStructInstance := SomeStruct{KeyA: "ValueA", KeyB: false }
	// Content-Type header automatically set!
	c.JSON(200, myStructInstance)
}
```

Registering a handler func to a route on the Gin router is done on an HTTP-method-specific basis:

```go
r.GET("/something", somethingGetHandler)
r.POST("/something", somethingPostHandler)
```

Starting up the server on a specified port:

```go
r.Run("localhost:4321")
```

# Go mod file

Think of this like `package.json`.

Initializing a new go mod for your project is done with the `go mod init` command.

Run this at the root level of your repo:

```bash
go mod init github.com/codecademy-engineering/intermediate-go-training-2023
```

Naming your module with the the repo url allows your code to become a dependency in another project.

# Installing 3rd party dependencies

Adding a dependency to your project is done via the terminal.
Run the following at the root of your project.

```bash
go get github.com/gin-gonic/gin
```
