
# intermediate-go-training-2023

Time: 2 hours

## Theme
We'll use the images API from NASA https://images-api.nasa.gov/search?q=
Our end result would be wrapping this api in a GraphQL API so that
- We can have self-documenting schema
- We can hover over hrefs and see the images

Topics to cover (9):
- Initialize project via `go mod init <repo>`
- Create a basic http server that returns a marshled struct
- Modules (basic import of a func from one module into main, examplain the dir structure of app)
  - Create a "models" package with a SearchResult struct
- Embedding And Composition (maybe within context of nested structs for JSON?)
  - Embed a Link in a SearchResult struct
  - Add JSON serializing annotations (test with fmt.Println)
- Methods on structs, pointers (auto de-referencing)
  - Create a fetcher with a Fetch method, using a constructor pattern
- Multi-returns (including errors)
  - Make the fetcher return a search results and error
- Interfaces (possibly in context of tests?)
  - Given X (struct) takes a Y (interface), test X with a MockY instead of a RealY
  - Test that we do and don't return errors when we expect
- Routing And Muxing
  - Initially return from REST (to test JSON serialization)
- HTTP Handlers
  - Create a basic status handler
- GraphQL
- Configs (viper)
  - Add an API key for the NASA API

## Notes
Each numbered step dir could contain:
1. Toy example of the concept for that step
2. Link to a place in a real CC repo with that concept
3. Some instructions to the dev to complete something with that concept
4. Example "checkpointed" code for the project at that stage where someone could copy/paste to get up to speed
5. "Hidden" answer to that step that someone could check

Root of dir is the main.go where each dev is coding the final project.

## Steps in workshop
1. Introduce the project overall
2. Make sure everyone is on certain go version??
3. Clone the repo locally and checkout a new branch
