# Workshop Protobuf Reflect

Example code for a small workshop or presentation on (primarily) Golang
protobuf reflection.

## Introduction

These examples demonstrate four different approaches to generic handling of
data.

For the sake of simplicity, map and list fields have not been included and
field types are limited to strings, integers and nested objects.

The four examples use:

1. Map structures
2. Tree structures
3. Standard reflection
4. Protobuf reflection

Every example expands upon the feature set provided by the previous approach.
As this example set was designed around (reasons for) using Protocol Buffers
reflection, it culminates there. This is by no means intended to suggest that
Protocol Buffers is 'better'.

For demonstrating the cross-language nature of Protocol Buffers' annotation
system, a small Python example is included as well.

## Features Sets

|                        | map | tree | reflect | proto |
|------------------------|-----|------|---------|-------|
| easy                   | +   | +    | -       | o     |
| get                    | +   | +    | +       | +     |
| set                    | +   | +    | +       | +     |
| traversal              | -   | +    | +       | +     |
| type safe              | -   | +    | +       | +     |
| schema*                | -   | -    | +       | +     |
| annotations            | -   | -    | +       | +     |
| external** annotations | -   | -    | -       | +     |

*) a (possibly) non-recursive schema fit for use in databases

**) outside of code base
