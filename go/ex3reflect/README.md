# Example 3: Using a Reflection

Most programming language have reflection (or inspection) which allow varying 
degrees of (generic) grey to black magic on data structures. Go is no exception 
in this.

We can use this to generate a complete type schema. The only thing we need to 
take into account is that data types can be recursive, causing infinite loops 
(and infinitely large schemas). In this example we therefore set a maximum 
depth for the schema - even though the example's data type is not recursive.

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
