# Gin

Gin is one of the most popular web frameworks for golang.

We use it often in repos at Codecademy (it even ships with the [go service cookie-cutter template](https://github.com/codecademy-engineering/cookiecutter-service-go)).

# Gin Router

Creating a Gin router:

```go
r := gin.Default()
```

# Gin handlers

Handler funcs for Gin take a pointer to a gin 'Context' which is a simplified way of using the ResponseWriter and request pointer from standard lib.

Additionally, the Gin context includes convenient helpers such as `c.JSON` which makes rendering JSON responses very simple.

```go
func myHandler(c *gin.Context) {
	myStructInstance := SomeStruct{KeyA: "ValueA", KeyB: false }
	// Content-Type header automatically set!
	c.JSON(200, myStructInstance)
}
```

## gin.H for simple JSON

If you want to return some kind of standardized data entity (struct instance, slice), then using c.JSON with that entity (ex above) is typically what you want both for convenience and consistency.

However, sometimes we want to return small things like messages or error responses which might not warrant declaring a struct for that purpose.

For those cases, Gin offers a special `gin.H` struct.
This is a shorthand custom datatype equivalent to map[string]interface{} which means a key/val pairing where keys are strings and vals can be anything.

You can then inline your declarations like so, which incidentally feel more like writing javascript!

```go
c.JSON(200, gin.H{"a":1, "b":"two", "c":false})

c.JSON(500, gin.H{"message":"server imploded"})
```

# Registering Gin handlers

Registering a handler func to a route on the Gin router is done on an HTTP-method-specific basis:

```go
r.GET("/something", somethingGetHandler)
r.POST("/something", somethingPostHandler)
```

# Starting Gin

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
