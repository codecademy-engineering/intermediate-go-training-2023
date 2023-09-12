We have one final task for your contract with us. In the event that an agent
is confronted by one of these squirrels, we need a way for them to send an alert
to other agents, indicating which squirrel is now posing a threat.

Another developer will be implementing a "beacon" later for this purpose, which
will send a notification to all smartwatches that the agents are wearing.

In the interim, we need for you to stub out this functionality. This way we'll
be ready to connect the real beacon implementation as soon as it's ready.

# Your mission (should you choose to accept it)

1. Refactor lookup logic from the GetByID handler into a reusable private method
that takes an id (string) and returns a pointer to a squirrel and an error if no
squirrel matching the id is found.

2. Create an alert package with a beacon interface. This interface should have
an "AlertAllAgents" method that accepts a pointer to a squirrel and has a return
type of error.

3. Write a mock implementation of the beacon interface within the same alert
package. The AlertAllAgents method can always return a nil error, simulating
a sucessful alert.

4. Refactor the SquirrelHandler struct so that contains a field that is a
beacon interface. Change the factory func to take a beacon instance and attach
it to the SquirrelHandler struct when initializing.

5. In your main func, create an instance of the mock beacon, and pass it into the
SquirrelHandler factory.

6. Create an Alert handler method on the SquirrelHandler struct, for
"/squirrel_alert/:id". This handler should use the same lookup method to find a
squirrel for the provided id, and then call the "AlertAllAgents" method on the
beacon interface with that squirrel pointer.

7. Return the following http responses
- 404 and a "squirrel not found" message, if no squirrel is found matching the id
- 500 and an error message, if the beacon AlertAllAgents method returns an error
- 200 and a success messsage, if agents were successfully alerted.
