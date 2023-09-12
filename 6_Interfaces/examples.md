# Interfaces

Abstractions that define which methods (and signatures) an object should have.

Other code can then rely on the interface rather than a concrete struct.
The rest of the code therefore doesn't know or care which concrete implementation it's working with.

ANY struct that has method signatures matching the interface are said to "satisfy the interface" and can be passed in its stead.

```go
type MyInterface interface {
	// note uppercase S
	SomePublicMethod(argName string) error
	// lowercase s
	somePrivateMethod(argName string) (string, error)

	MethodWithAnonArgName(string, int) error
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
func (s *myInterfaceImpl) SomePublicMethod(argName string) error {
	//...impl logic
}

// so when calling the func expecting an interface, we can give it the implementation

myImpl := myInterfaceImpl{}
someFunc(myImpl)
```
