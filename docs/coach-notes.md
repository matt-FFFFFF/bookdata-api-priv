# Coach notes for Golang Beginners Hack

## Challenge 1

Outside of the success criteria, make sure to highlight the interface used by the server.
The interface is defined in ```datastore/datastore.go``` and the ```server.go``` create a var when it starts up.

In the init function the books var gets assigned to an pointer to an empty datastore.Books struct.
This is allowed because, in ```datastore/memory.go``` the Books type implements all the methods of the interface.

THis means that the datastore can change, as long as it implements the requried method signature.

Point the group towards <https://medium.com/golangspec/interfaces-in-go-part-i-4ae53a97479c> for further reading if interested.

## Challenge 2
