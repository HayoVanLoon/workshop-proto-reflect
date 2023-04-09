# Example 1: Using a Map (JSON object)

We start of with a basic map of string to interface. This is basically the
deserialised version of a JSON object.

## Read a Value

We start with reading a value indicated by a path of names. There is a standard
pattern here that we will expand upon for other operations and in other
examples.

1. follow the path
2. check if we are at the destination
3. if so: operate (return value)
4. otherwise: try to go deeper

## Updating a Value

The updating of a value follows the same pattern. However, instead of
returning the value, we set it.

Another difference is that we can always go deeper; we can just initialise a
new section of the path (aka: a new map). We can add anything really.

## Traversal

Traversal allows us to operate on all values. For the example, we chose to let
it flatten the structure, return a list of all its values. Like with the Read
and Update examples, we also could have made it update fields instead of
returning them.

The pattern is even simpler:

1. walk over the fields
2. if the field is a simple value: operate
3. otherwise: go deeper

But it is here that the first cracks in our approach start to appear. The order
of map entries is not defined. This means that the traversal would be random,
which in turn means that the order of items in the resulting list might vary.
This might be a deal-breaker for your use case, or it might not. When updating
via traversal, this randomness is very likely not an issue.

## Features

|              | + | 0 | - |
|--------------|---|---|---|
| easy         | x |   |   |
| get          | x |   |   |
| set          | x |   |   |
| add anything | x |   |   |
| traversal    |   | x |   |
| type safe    |   |   | x |
