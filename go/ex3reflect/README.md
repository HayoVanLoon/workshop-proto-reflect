# Example 3: Using a Reflection

Most programming language have reflection (or inspection) which allow varying 
degrees of (generic) grey to black magic on data structures. Go is no exception 
in this.

In any language (that I know of), reflection is a finicky beast, sometimes with
rather sharp claws. The operations here however need not venture too deeply 
into unsafe (pun intended) territory, if at all. 

## Features

|             | + | 0 | - |
|-------------|---|---|---|
| easy        |   |   | x |
| get         | x |   |   |
| set         | x |   |   |
| traversal   | x |   |   |
| type safe   | x |   |   |
| schema*     | x |   |   |
| annotations | x |   |   |

*) a (possibly) non-recursive schema fit for use in databases
