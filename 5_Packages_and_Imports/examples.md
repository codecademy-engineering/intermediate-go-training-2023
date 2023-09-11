# Creating a new package
under /mypackage/somefile.go
```go
package mypackage

// Structs and functions starting with an uppercase letter can be
// accessed in other places where this package is imported.
type MyExportedStruct struct {}

func MyExportedFunc() {}

// Structs and functions starting with a lowercase letter can only be
// accessed internally within the package where they are defined.
type myPrivateStruct struct {}

func myPrivateFunc() {}
```


# Importing from the package
from main.go
```go
import "github.com/mymodulepath/mypackage"

func main() {
  s := mypackage.MyExportedStruct{}
  mypackage.MyExportedFunc()
  // If you try to import mypackage.myPrivateStruct...it won't work
}
```