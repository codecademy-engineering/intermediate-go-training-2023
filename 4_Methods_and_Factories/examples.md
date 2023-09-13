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

instance := StructWithAbilities{data: []int{1, 2}}

res := instance.GetData()
```

The alternative to a 'pointer' receiver is called a 'value' receiver, which receives a COPY of the data in the struct rather than the exact instance data already allocated in memory.
**(Don't worry about these for now.)**

# Factory Pattern

Common in golang to initialize a new instance of some struct that also has methods on it.

By convention, begin the function name with `New` or `new` then the name of the struct.
Typically this goes directly beneath the struct definition and above any struct methods.

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
	// OPTIONALLY, call some other code,
	// include more complicated setup, etc.

	// return a ptr to the instance (most common)
	return &SomeStruct{
		FieldA: a,
		FieldB: b,
	}
}
```

# Creating a path with a parameter in Gin

```go
r.GET("/somepath/:myparam", SomeStruct.MyHandlerMethod)
```

# Reading a path parameter in Gin

```go
func (h *SomeStruct) MyHandlerMethod(c *gin.Context) {
	myParam := c.Param("myparam")
	// ...some logic to use the param
}
```
