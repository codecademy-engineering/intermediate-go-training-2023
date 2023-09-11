# Reading a file

```go
rawBytes, err := os.ReadFile("./path/to/file.whatever")
```

# Unmarshalling JSON (slice of bytes in JSON -> go struct)

Fields MUST be uppercased for the unmarshalling, but fields can have json aliases
when converting to/from structs if you want to have lowercase names (or any other names) in JSON.

```json
{"numberField":1, "booleanField":false}
```

```go
rawData := // json above as a slice of bytes (e.g. read from a json file or an http response)

type SomeStruct struct {
	AliasOne int `json:"numberField"`
	AliasTwo bool `json:"booleanField"`
}

targetInstance := SomeStruct{}
finalStruct, err := json.Unmarshal(rawData, &targetInstance)
```

Structs can be nested. If your json data has a nested structure, you might want to
unmarshall it into a struct that also uses nesting.

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
