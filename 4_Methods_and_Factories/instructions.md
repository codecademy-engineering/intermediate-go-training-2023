Ok, not bad. We can access the list of all squirrels at once. BUT...

This implementation could be more flexible and efficient. Plus we'd like to do more than just read a long list...

We'd like this service to read the datafile only once on boot and keep the contents available in memory.

In addition, this same data could be shared by another endpoint that provides our agents a quick way to request data on just 1 squirrel as opposed to all squirrels at once.

This could be achieved with a handler STRUCT rather than a standalone handler func.

# Your mission (should you choose to accept it)

1. Refactor the logic that unmarshalls the squirrel intel contained within the `data.json` file into a stand alone func. This func should be called first in your main func and panic if the data can't be read.

2. Write a factory func that takes in the successfully unmarshalled data and uses it to initialize a `SquirrelHandler` struct.

This handler struct should have a method called `GetAll` which accesses the in-memory squirrel data housed in the struct in order to return the same JSON response as the original gin handler func you wrote previously.

3. Replace your original handler func on the `/squirrels` route with a call to your new handler struct method instead.

4. Add another method to the SquirrelHandler struct that will read an "id" parameter.

- Range through each squirrel in the data set and return the squirrel with matching ID if found.
- Return a 404 nil result if no matching squirrel is found.

5. Register this new handler method with the Gin router so it calls this method when going to "/squirrels/:id".
