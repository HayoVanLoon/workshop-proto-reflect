# Example 2: Using a Tree

This example features a library with a limited tree implementation.

By storing its children as a list (not a map) and defining an order (insertion
order), it provides a stable result for the traversal method. By storing type
information along with the value, it achieves some measure of type safety. 

It follows all the patterns for data operations introduced in the first example.

But what it cannot do is describe the data when it's not (yet) there. In other 
words: it has no schema. We could design a very generic schema. The structure 
is recursive however and without any bounds, the schema would become infinitely 
wide and infinitely deep. Not usable for let's say, a database. 

This example is a bit contrived as, with enough effort, the tree structure 
could be made to do anything (like having some way of setting a schema). Though 
then it would arguably no longer be a tree. 

## Features

|              | + | 0 | - |
|--------------|---|---|---|
| easy         | x |   |   |
| get          | x |   |   |
| set          | x |   |   |
| add anything | x |   |   |
| traversal    |   | x |   |
| type safe    |   |   | x |
| schema*      |   |   | x |

*) a (possibly) non-recursive schema fit for use in databases
