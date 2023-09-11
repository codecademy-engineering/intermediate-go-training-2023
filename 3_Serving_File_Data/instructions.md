You've proven your skills and are now privy to the top secret mission for which you have been chosen.

You're not just any contract hire...and we're not just any agency (\*ehem) client...

Take a look at the `data.json` file under the `static` dir.
This is super top secret squirrel data that we (the CIA) have collected for monitoring evil squirrel security threats in various parks across NYC.

# Your assignment (should you choose to accept it)

We want you to build a go service that our agents in the field can call to get up-to-date info on our furry adversaries, starting with an endpoint that returns all the data at once.

DISCLAIMER: We know this may seem strange to unmarshall and remarshall json data to serve over the endpoint, but this is just a stepping stone towards a subsequently better way to handle the data.

In your `main.go` file, write golang structs that match the shape of the data you see.

Then create a new handler func to serve on `/squirrels` that **at request time** will unmarshall the contents of the `data.json` file into those go structs, and return it as a JSON response with 200 status code.
