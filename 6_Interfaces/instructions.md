We have one final task for your contract with us. In the event that an agent
is confronted by one of these squirrels, we need a way for them to send an alert
to other agents, indicating which squirrel is now posing a threat.

Another developer will be implementing a "beacon" later for this purpose, which
will send a notification to all smartwatches that the agents are wearing.

In the interim, we need for you to stub out this functionality. This way we'll
be ready to connect the real beacon implementation as soon as it's ready.

# Your mission (should you choose to accept it)

1. Begin by extracting the lookup logic portion of the GetByID handler into a reusable **private** method
   that takes an id (string) with return types of Squirrel pointer and error.

   - If a matching Squirrel is found, return a pointer to that Squirrel and a nil error
   - Otherwise return a nil Squirrel pointer and an error with the text "squirrel not found"

   The GetByID handler should check the error value returned by the extracted lookup logic and include it as the JSON message text on a 404 scenario.

2. Create an alert package containing a Beacon interface definition. This interface should define
   an "AlertAllAgents" method that accepts a pointer to a squirrel and has a return
   type of error.

3. Write a mock implementation of the Beacon interface within the same alert
   package and associated factory func. The AlertAllAgents method should always return a nil error, simulating a sucessful alert.

4. Add a new beacon field of type Beacon interface to the SquirrelHandler struct.
   Change the factory func to take a beacon instance and attach it to the SquirrelHandler
   struct when initializing.

5. In your main func, create an instance of the mock beacon, and pass it into the
   SquirrelHandler factory.

6. Create a ThreatAlert handler method on the SquirrelHandler struct to handle
   **A POST REQUEST** to "/squirrels/:id/threat".

   This handler should use the same extracted lookup method used by the GetByID handler to find a
   squirrel for the provided id, and then call the "AlertAllAgents" method on the
   beacon interface with that squirrel pointer.

   This handler method should return the following http responses:

   - 404 and the lookup logic error message if no squirrel is found matching the id
   - 500 with a message value of "beacon failure: <error text>" with string interpolation of the actual error if the beacon AlertAllAgents method returned one
   - 200 and a success messsage of "Alerted all agents that squirrel <id> is an imminent threat!" with string interpolation if the agents were successfully alerted.
