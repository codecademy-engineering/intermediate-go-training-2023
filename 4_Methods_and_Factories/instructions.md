Ok, not bad. We can access the list of all squirrels at once now. BUT...

This implementation could be more flexible and efficient. Plus we'd like to do more than just read a long list...

We'd like this service to read the datafile only once on boot and keep the contents available in memory.

In addition, this same data could be shared by another endpoint that provides our agents a quick way to request data on just 1 squirrel as opposed to all squirrels at once.

This could be achieved with a handler STRUCT with handler METHODS rather than a standalone handler FUNC.

# Your mission (should you choose to accept it)

1. Create a SquirrelHandler struct with a field to keep a slice of all the squirrels.

2. Add a method on the struct called `GetAll` which takes the gin context as an arg. This method should accesses the
   squirrel data housed in the struct instance and return the same JSON response as before with the list of all squirrels.

3. Add another method to the SquirrelHandler struct called `GetByID` that reads an "id" from the gin context. This method
   should find the squirrel with the matching id and include it in a 200 response. If no matching squirrel is
   found, a `{"message": "squirrel not found"}` response should be given with a 404 status code.

4. Create a factory func that reads and unmarshals the squirrel data and creates a new instance
   of the SquirrelHandler struct containing the slice of Squirrels. This factory should panic
   if the data can't be read or unmarshalled.

5. Call the factory func from main() to create a SquirrelHandler instance.

6. Register both handler methods so that `GetAll` handles `/squirrels` and `GetByID` handles `/squirrels/:id`.
