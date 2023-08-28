Unmarshalling JSON
```go
type ExampleStruct struct {
  MyNumber int
}
myStructInstance := ExampleStruct{MyNumber: 5}
json.Unmarshall(&myStructInstance)
```