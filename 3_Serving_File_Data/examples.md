# Reading a file

```go
rawBytes, err := os.ReadFile("./path/to/file.whatever")
```

# Unmarshalling JSON (slice of bytes in JSON -> go struct)

Note: fields MUST be uppercased for the unmarshalling

TODO: aliasing vs default

```go
rawData := []bytes{`{"numberField":1, "booleanField":false}`}

type SomeStruct struct {
	AliasOne int `json:"numberField"`
	AliasTwo bool `json:"booleanField"`
}

targetInstance := SomeStruct{}
finalStruct, err := json.Unmarshal(rawData, &targetInstance)
```

Structs can be nested
and fields can have json aliases when converting to/from structs

```go
type author struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type Book struct {
	Title string // no need to alias if field names match
	Author author
}
/* json equivalent:
	{
		"title": "something",
		"author": {
			"first_name": "Bob",
			"last_name": "Jones"
		}
	}
*/
```
