Initializing a new go mod for your project is done with the `go mod init` command.
Run this at the root level of your repo:

```bash
go mod init github.com/codecademy-engineering/intermediate-go-training-2023
```

Adding a dependency to your project is done via the terminal.
Run the following at the root of your project.

```bash
go get -u github.com/gin-gonic/gin
```

## Gin

One of the most popular web frameworks for golang. We use it often in repos at Codecademy (it even ships with the service cookie-cutter for golang).

Creating a Gin router:

```go
r := gin.Default()
```

Handler funcs for Gin look slightly different.
They take a pointer to gin 'Context' which is a simplified way of using the ResponseWriter and request pointer from standard lib.

Gin context includes convenient helpers such as automatically marshalling your structs to JSON with proper response code and Content-Type headers included:

```go
func myHandler(c *gin.Context) {
	myStructInstance := SomeStruct{KeyA: "ValueA", KeyB: false }
	c.JSON(200, myStructInstance)
}
```

Registering a handler func to a route on the Gin router is done on HTTP method basis so the same URL can be easily reused:

```go
r.GET("/something", somethingGetHandler)
r.POST("/something", somethingPostHandler)
```

Starting up the server on a specified port:

```go
r.Run("localhost:4321")
```
