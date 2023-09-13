# Interfaces

Abstractions that define method signatures & return types that some **unspecified** concrete struct will implement.

Interfaces are like "contracts". Wherever they are used in code, they promise that **something** will be there at runtime with the same methods and return types defined by that interface.

This is helpful to modularize your code and swap implementations in/out without needing to refactor the code that relies on the interface.

ANY struct that has AT LEAST ALL (possibly more!) matching method signatures & return types are said to "satisfy the interface".

```go
type MyInterface interface {
	// note uppercase S
	SomePublicMethod(specificArgName string) error
	// lowercase s
	somePrivateMethod(specificArgName string) (string, error)

	MethodWithAnonArgNames(string, int) error
}

// using an interface is much like the struct
func someFunc(someImpl MyInterface) {
	err := someImpl.SomePublicMethod("hello")
}

type myInterfaceImpl struct {
	// maybe some stuff in here, doesn't matter!
	// interfaces don't define the fields, only the methods!
}

// this method matches the interface definition above
func (s *myInterfaceImpl) SomePublicMethod(specificArgName string) error {
	//...impl logic
}

// so when calling the func expecting an interface, we can give it the implementation

myImpl := myInterfaceImpl{}
someFunc(myImpl)
```

# Nil pointers

When your return type on a func is a pointer, a nil pointer can be used
as an empty result

```go
type Thing struct {}

func findAThing(id string) (*Thing, error) {
	// some logic to look up a thing by id
	if (thingFound) {
		return *thing, nil
	}
	return nil, fmt.Errorf("%s not found", id)
}
```

**_Warning:_** Trying to dereference a nil pointer will cause the program to crash!
For this reason, it's important to return an error so that calling code can check it and avoid
accidentally dereferencing a nil pointer!

# Gin.H

Another convenience offered by Gin.
Use this instead of a map[string]interface{} when creating something to serialize to JSON.

```go
c.JSON(200, gin.H{"a":1, "b":"two", "c":false})
```

