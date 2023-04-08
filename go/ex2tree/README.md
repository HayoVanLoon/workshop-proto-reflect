# Example 2: Using a Tree

This example features a library with a limited implementation of a tree.

It stores and checks a node type's value, so there is some measure of type
safety.

By storing its children as a list (not a map) and defining an order (insertion
order), it provides a stable result for the traversal method.

## Features

|           | + | 0 | - |
|-----------|---|---|---|
| easy      | x |   |   |
| get       | x |   |   |
| set       | x |   |   |
| traversal |   | x |   |
| type safe |   | x |   |
| schema*   |   |   | x |

*) a (possibly) non-recursive schema fit for use in databases
