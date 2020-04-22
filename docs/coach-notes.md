# Coach notes for Golang Beginners Hack

## Challenge 1

Outside of the success criteria, make sure to highlight the interface used by the server.
The interface is defined in ```datastore/datastore.go``` and the ```server.go``` create a var when it starts up.

In the init function, the books var gets assigned to an pointer to an empty datastore.Books struct.
This is allowed because, in ```datastore/memory.go``` the Books type implements all the methods of the interface.

This means that the datastore can change, as long as it implements the requried method signature.

Point the group towards <https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c> for further reading if interested.

## Challenge 2

### On defer

The critical thing to convey here is that defer is a function that is run at the end of the current function, wherever it exits.
This is tidier than try-catch-finally, as there are no braces to worry about.

Although the function is executed when the current function exits, the vlaues passed into the defer are set a the time of the defer.
This way we can pass in the current time when we call defer. When defer runs we can perform a simple time.Since calculation.

## Challenge 3


