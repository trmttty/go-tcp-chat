# go-tcp-chat
This is a simple chat using TCP.
## usage
- Setting up a chat server.
```
go build && ./go-tcp-chat
```
- Accessing the chat server using telnet in another terminal.
```
telnet localhost 8888
```
- You will be asked to register your name
```
> welcome! please enter your name
<your name>

then...

> all right, I will call you <your name>
```
- You can view the list of chat rooms with the following command
```
/rooms
```
- There are two types of chat rooms: public and private.  
Private chat rooms are only visible to those who have been granted access.
```
> available public rooms are: general random
> available private rooms are: project1 project2 
```
