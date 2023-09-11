You've contracted for a secret assignment. Before we reveal what it is, you'll need
to demonstrate your golang chops.

# Your assignment (should you choose to accept it)

Create a main.go file at the root of this repository.

Copy the template below into that main.go file.

Fill out the template in order to create a server that returns json containing
a key of "Status" (note the capital 'S' letter), a value of "I'm a teapot", and a status code set to 418.

## Template

```go
// declare a package

import (
	"encoding/json"
	"net/http"
)

// define a struct for your JSON response

func handleRoot(w http.ResponseWriter, r *http.Request) {
	// create an instance of your struct

	// marshall the struct instance into JSON

	// handle if err is not nil (using panic() is fine)

	// set the "Content-Type" header

	// write the response code

	// write the response body
}

func main() {
	// instantiate a mux

	// attach handleRoot to the "/" path

	// start the server on localhost:4321
}
```
