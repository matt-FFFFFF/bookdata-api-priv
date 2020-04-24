# Coach notes for Golang Beginners Hack

## Challenge 1

Outside of the success criteria, make sure to highlight the interface used by the server.
The interface is defined in ```datastore/datastore.go``` and the ```server.go``` create a var (books) when it starts up.

In the init function, the books var gets assigned to a pointer to an empty datastore.Books struct.
This is allowed because, in ```datastore/memory.go``` the Books type implements all the methods of the interface.

This means that the datastore can change, as long as it implements the requried method signature.

Point the group towards <https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c> for further reading if interested.

## Challenge 2

### On defer

The critical thing to convey here is that defer is a function that is run at the end of the current function, wherever it exits.
This is tidier than try-catch-finally, as there are no braces to worry about.

Although the function is executed when the current function exits, the values passed into the defer are set a the time of the defer call.
This way we can pass in the current time when we call defer. When defer runs we can perform a simple ```time.Since``` calculation.

## Challenge 3

Suggest the teams split the tasks into pairs.

There will probably be a lot of code repetition. Once answer to this is the ```Filter``` func in ```memory.go```.
The first parameter to this is the in-memory store itself.
The second paramater is an inline function that takes in a (pointer to a) single item from the book store ```*loader.BookData```, then performs some comparison with the supplied API request and returns a bool.
This way we can filter the store in any way we like without lots of repetition (for loops).

It may be that the teams do not come up with this concept, especially if they are not experienced in go or development in general.
Point this out to them but it is by no means essential to pass the challenge.

## Challenge 4

Key thing is to use multiple stages in the Dockerfile.
A builder container, with the tools required to build the binary, and a run container that the binary is copied from.

The build stage should use the same base as the dev container, ensuring consistency between dev and build.
