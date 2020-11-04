# go-tcp-chat
This is a simple chat using TCP.
## Usage
- Set up a chat server.
```
go build && ./go-tcp-chat
```
- Access the chat server using telnet in another terminal.
```
telnet localhost 8888
```
- You will be asked to register your name.
```
> welcome! please enter your name
<your name>

> all right, I will call you <your name>
```
- List chat rooms.
- There are two types of rooms: public and private.  
Private rooms are only visible to those who have been granted access.
```
/rooms

> available public rooms are: general random
> available private rooms are: project1 project2 
```
- Create a public chat room.
```
/create <room name>
```
- Create a private chat room.
```
/create <room name> private
```
- Join a chat room.
```
/join <room name>
```
- Invite a new member to a private room.  
You need to be in the room you want to invite a new member to.
```
/invite <member name>
```
- List all members in the chat server.  
```
/members
```
- Send a message to the all members in the room.
```
/msg <message>
```
- Change your registered name
```
/rename <new name>
```
- Leave the chat server
```
/quit
```
