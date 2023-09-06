# Your mission (should you choose to accept it)

1. Add another method to the SquirrelHandler struct that will read an "id" parameter.
- Range through each squirrel in the data set and return the squirrel with matching ID if found.
- Return a 404 nil result if no matching squirrel is found.

2. Register this new handler method with the Gin router so it calls this method when going to "/squirrels/:id".
