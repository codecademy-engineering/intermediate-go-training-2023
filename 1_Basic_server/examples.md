# Creating a package

Your main.go file should always use a package called "main"

```go
package myPackage
```

# Creating a struct

```go
type MyStruct struct {
  MyProp string
}
```

# Instantiating a struct

```go
myInstance := MyStruct{ MyProp: "some text" }
```

# Multiple return values

It is very common in golang to receive 2+ return values from a function.

The last value will often be of type `error` and will either be:
A) some error instance if there was a problem
B) `nil` if everything was successful

By convention, the variable name developers tend to choose for the error return is `err`.

```go
val, err := someFunc()
```

You are always guaranteed that a function will return the same number of values and in the same order as stated in its function signature.

# Error handling

There is no try/catch in golang.

Instead, you must leverage the multiple returns and check the error value yourself.

```go
someVal, err := myFunc()
if err != nil {
	// handle the error
}
```

Panic! (not for production)
It will print the error or message passed to it and end the program immediately.
Useful for quick prototyping.

```go
panic("Something has gone very wrong...halt all the things")
```

# JSON Marshalling (struct instance -> JSON string)

Equivalent of javascript JSON.stringify() except it is a slice of bytes (think raw binary data)

```go
rawJSONBytes, err := json.Marshal(myInstance)
if err != nil {
  // handle the error
}
```

# Creating a route handler func

```go
func handleRoute(w http.ResponseWriter, r *http.Request) {
	// writing the response body as slice of bytes
	w.Write([]byte("Hello client! From, server"))
}
```

# Setting a status code (will be 200 if not set)

```go
w.WriteHeader(500) // w is an http.ResponseWriter
```

# Setting response content type

```go
// w is an http.ResponseWriter
w.Header().Set("Content-Type", "application/json")
```

# Creating a Mux (let's be honest, it's a Router)

```go
mux := http.NewServeMux()
```

# Attaching the handler to the mux

```go
mux.Handle("/sayHello", http.HandlerFunc(handleRoute))
```

# Starting the server

```go
http.ListenAndServe("localhost:4321", mux)
```
