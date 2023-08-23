Creating a struct
```go
type MyStruct struct {
  MyProp string `json:"What this will be when serialized"`
} 
```

Instantiating a struct
```go
myInstance := MyStruct{ MyProp: "some text" }
```

JSON Marshalling (struct instance -> JSON string)
```go
jsonString, err := json.Marshal(myInstance)
if err != nil {
  // handle the error
}
```

Creating a Mux (let's be honest, it's a Router)
```go
mux := http.NewServeMux()
```

Create a route handler func
```go
func handleRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello client! From, server"))
}
```

Set a status code (will be 200 if not set)
```go
w.WriteHeader(500) // w is an http.ResponseWriter 
```

Set response content type
```go
// w is an http.ResponseWriter 
w.Header().Set("Content-Type", "application/json")
```

Attach the handler to the mux
```go
mux.Handle("/sayHello", http.HandlerFunc(handleRoute))
```

Start the server
```go
http.ListenAndServe("localhost:3000", mux)
```

Panic! (not for production)
It will print the error or message passed to it and end the program immediately.
```go
panic("Something has gone very wrong...halt all the things")
```
