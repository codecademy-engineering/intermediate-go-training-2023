# Methods on Structs

A `method` in golang is a special function that is attached to a struct instance.

Methods are most commonly 'pointer receivers' meaning they can directly access shared, internal data on the struct instance they belong to.

```go
type StructWithAbilities struct {
	data []int
}

// pointer receiver denoted here by the *...
// 's' inside this method will refer to the specific instance of StructWithAbilities rather than a copy of the data.
func (s *StructWithAbilities) GetData() []int {
	return s.data
}
```

# Factory Pattern

Common in golang to initialize a new instance of some struct that also has methods on it.

By convention, begin the function name with `New` or `new` then the name of the struct.

Often times this is used to initialize the entity with some starting values that will then be utilized by the methods called on the returned struct instance.

```go
type SomeStruct struct {
	FieldA string
	FieldB boolean
}

func (s *SomeStruct) CoolMethodHere() (string, boolean) {
	return s.FieldA, s.FieldB
}

func NewSomeStruct(a string, b boolean) *SomeStruct {
	// most common to return a ptr to the instance
	return &SomeStruct{
		FieldA: a,
		FieldB: b,
	}
}
```
