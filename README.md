# H2H-Matrix
This projects attempts to implement a head-to-head matrix displaying the W-L results of matches across different teams, which reads data from a JSON file.

## Run
Run the following command to see the output matrix:

```bash
go run main.go <example.json>
```

## Input Data Format
The input data is in JSON format, structured as follows:

- The top-level key is a team's name
- Each team maps to its opponents
- Each opponent maps to a W-L record

```json
{
  "BRO": {
    "BSN": { "W": 10, "L": 12 },
    ...
  },
  "BSN": {
    "BRO": { "W": 12, "L": 10 },
    ...
  },
  ...
}
```

## Explaination of My Choices
- **Language**: I have been delving into Go recently and completed a few projects with it, as I was suprised to find out that Sports Reference lists Go as one of the languages they use.
- **Data Model**:
  - The nested map structure ***H2H*** mirrors the JSON input
  - A small ***Record*** struct makes the W-L concept explicit. In this project, only Ws are used for display, but having both is always for an easy future.
  - Parsing JSON in Go is straightforward! I love it. 
```go
type Record struct {
    W int `json:"W"`
    L int `json:"L"`
}

type H2H map[string]map[string]Record

```
- **Processing Steps**:
  - Parse the JSON file into memory in main()
  - getTeams() use a nested loop iterating ***H2H*** to extract all teams' names from the matrix headers into a set
  - Some teams may only appear as opponents in the data and not as keys in the outer map, so iterating the nested maps is necessary to avoid missing teams
```go
  set := make(map[string]struct{})

  for team, opponents := range h2h {
      set[team] = struct{}{}
      for opponent := range opponents {
          set[opponent] = struct{}{}
      }
  }
```
 - Convert the set into a sorted slice then return it. In main() we get the slice of teams.
 - printMatrix() formats and prints the matrix to the console.
   - The function takes the ***H2H*** data and the slice of ***teams*** as input
   - First, print the header row with all team names
   - Then, for each team (row), print the W counts against each opponent (column)
   - If no record exists, leave the cell empty
   - If the team is the same as the opponent, print a dash ("--"), as we see in the diagonal cells
   - The only trick for formatting is to use a fixed width for each column

## Some Thoughts Afterwards
- I am once again amazed by Go's simplicity and efficiency after completing this small project.
- There's one small trap in handling the data somehow reflects a real-world scenario. Say for Major League Baseball there might be some cases we have to store the data sharded by division, then when you load just one shard (say, AL East), the outer map keys will only include teams in that division;
  However, there are inter-division and interleague games, so some teams may only appear as opponents in the nested maps. Some teams may be silently dropped if we fail to consider this case in practice.