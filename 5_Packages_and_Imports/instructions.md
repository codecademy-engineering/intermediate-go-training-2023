Our field agents are very happy that they can at last get information
about the particular squirrels they are tracking. For future developers,
it would be more maintainable for this code to be split into packages.

# Your mission (should you choose to accept it)

1. Create new packages for handlers and models

2. Move the relevant parts of your code into their respective packages and update imports where necessary.

3. At the ends of this, your main.go file should only contain your main function. It should instantiate your SquirrelHandler, register handler methods with the Gin router, and start the server.
