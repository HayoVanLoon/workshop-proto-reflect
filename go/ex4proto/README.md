# Example 4: Using Protobuf

Protobuf's reflection API feels similar to the standard reflect library.
However, as it will only operate on Protobuf structures, there are way fewer
(edge) cases to take into account.

Its annotation system (via extensions) allows for richer annotations compared
to the standard library. And being a language-agnostic system, annotation
functionality can be implemented for different languages, such as
for [Python](../../py/README.md).

## Features

|                        | + | 0 | - |
|------------------------|---|---|---|
| easy                   |   | x |   |
| get                    | x |   |   |
| set                    | x |   |   |
| traversal              | x |   |   |
| type safe              | x |   |   |
| schema*                | x |   |   |
| annotations            | x |   |   |
| external** annotations | x |   |   |

*) a (possibly) non-recursive schema fit for use in databases

**) outside of code base
