# go-by-example

All Go by example exercises

- [Learn Go By Tests](https://quii.gitbook.io/learn-go-with-tests/)
- [Go by Example](https://gobyexample.com)

#### Free Courses
- [TechWorld with Nana Go Course](https://www.youtube.com/watch?v=yyUHQIec83I)
- [Golang Tutorial: Go Full Course](https://www.youtube.com/watch?v=YzLrWHZa-Kc)
## Learn Go By Tests

### Go Fundamentals

- [x] [Hello World](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world)
- [x] [Integers](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers)
- [x] [Iteration](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration)
- [x] [Arrays and slices](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices)
- [x] [Structs, methods & interfaces](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/structs-methods-and-interfaces)
- [x] [Pointers & Errors](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/pointers-and-errors)
- [x] [Maps](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps)
- [x] [Dependency Injection](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection)
- [x] [Mocking](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/mocking)
- [x] [Concurrency](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/concurrency)
- [x] [Select](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/select)
- [ ] [Reflection](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reflection)
- [x] [Sync](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/sync)
- [x] [Context](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/context)
- [x] [Intro to property based tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/roman-numerals)
- [ ] [Maths](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/math)
- [ ] [Reading Files](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/reading-files)
- [ ] [Templating](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/html-templates)
- [ ] [Generics](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/generics)
- [ ] [Arrays & Slices with generics](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/revisiting-arrays-and-slices-with-generics)

### Testing Fundamentals

- [ ] [Introduction to acceptance tests](https://quii.gitbook.io/learn-go-with-tests/testing-fundamentals/intro-to-acceptance-tests)
- [ ] [Scaling acceptance tests](https://quii.gitbook.io/learn-go-with-tests/testing-fundamentals/scaling-acceptance-tests)

### Build an application

- [ ] [HTTP Server](https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server)
- [ ] [JSON, routing and embedding](https://quii.gitbook.io/learn-go-with-tests/build-an-application/json)
- [ ] [IO and sorting](https://quii.gitbook.io/learn-go-with-tests/build-an-application/io)
- [ ] [Command Line & package structure](https://quii.gitbook.io/learn-go-with-tests/build-an-application/command-line)
- [ ] [Time](https://quii.gitbook.io/learn-go-with-tests/build-an-application/time)
- [ ] [Web Sockets](https://quii.gitbook.io/learn-go-with-tests/build-an-application/websockets)

## Golang Notes

#### Testing

- `cd ...` into the module you want to test
- run `go test` to execute the test suites for that module
- To execute a test and check for race conditions, execute `go test -race`
- To execute a test to make a benchmark execute `go test -bench=.`

#### Maps

An interesting property of maps is that you can modify them without passing as an address to it: `&myMap`

This may make them feel like a "reference type", but they are not.
A map value is a pointer to a runtime.hmap structure.

So when you pass a map to a function/method, you are indeed copying it, but just the pointer part, not the underlying data structure that contains the data.
A gotcha with maps is that they can be a nil value. A nil map behaves like an empty map when reading, but attempts to write to a nil map will cause a runtime panic.

Therefore, you should never initialize an empty map variable:
`var m map[string]string`

#### Concurrency

- You can use `close(channel)` to signal some type of event
- Always initialize a channel with `make(chan channel struct{})`
- You can use `mutex` from `sync` library to lock a property so it can only be updated by 1 Go subroutine at a time

#### Miscellaneous

- The smallest unit to save in memory is an empty struct `struct{}`, even smaller than boolean
- `interface{}` is the same thing as saying `any` in Typescript
- Context is used to propagate through requests and subrutines so that we can cancel processes in subrutines to avoid sloppy code and memory leaks
