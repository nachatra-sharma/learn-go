# learn-go

My Go learning notes and practice, organised by topic. Each folder is one small runnable
program (`package main` + `main.go`), so you can run any single topic on its own:

```sh
go run ./06-pointers/01-pointer-basics
go vet ./...          # type-check everything at once
```

## Topics so far

| Topic              | Folder                               | Covers                                                                                                                     |
| ------------------ | ------------------------------------ | -------------------------------------------------------------------------------------------------------------------------- |
| 1. Basics          | [01-basics](01-basics)               | printing, `Printf` verbs, variables & constants, `for` loops, `if/else`, `switch`, ranging over a string                   |
| 2. Functions       | [02-functions](02-functions)         | variadic functions (`args ...int`)                                                                                         |
| 3. Arrays & slices | [03-arrays-slices](03-arrays-slices) | fixed arrays vs slices, `append`, `len`, sum, max, count even, reverse in place, linear search, remove duplicates          |
| 4. Maps            | [04-maps](04-maps)                   | `make(map[k]v)`, frequency counter, find duplicate, two sum, intersection of two slices                                    |
| 5. Strings & runes | [05-strings-runes](05-strings-runes) | `map[rune]int`, count characters, first repeated character, anagram check                                                  |
| 6. Pointers        | [06-pointers](06-pointers)           | `&x` vs `*p`, mutating through a pointer, comparing pointers, swap via pointers                                            |
| 7. Structs         | [07-structs](07-structs)             | defining structs, field access, the three init styles, constructor returning `*T`, methods with value vs pointer receivers |
| 8. Interfaces      | [08-interfaces](08-interfaces)       | defining an interface, satisfying it with two types (stripe / razorpay), accepting an interface as a parameter             |
