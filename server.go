package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

type server struct {
	members  map[string]*client
	rooms    map[string]*room
	commands chan command
}

func newServer() *server {
	return &server{
		members:  make(map[string]*client),
		rooms:    make(map[string]*room),
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case CmdNick:
			s.nick(cmd.client, cmd.args)
		case CmdJoin:
			s.join(cmd.client, cmd.args)
		case CmdInvite:
			s.invite(cmd.client, cmd.args)
		case CmdMembers:
			s.listMembers(cmd.client, cmd.args)
		case CmdRooms:
			s.listRooms(cmd.client, cmd.args)
		case CmdMsg:
			s.msg(cmd.client, cmd.args)
		case CmdQuit:
			s.quit(cmd.client, cmd.args)
		}
	}
}

func (s *server) nick(c *client, args []string) {
	userName := args[1]

	if _, exists := s.members[userName]; exists {
		c.err(errors.New("username is already taken"))
	} else {
		s.members[userName] = c
		c.name = userName
		c.msg(fmt.Sprintf("all right, I will call you %s", c.name))
	}
}

func (s *server) join(c *client, args []string) {
	roomName := args[1]

	private := false
	if len(args) == 3 && args[2] == "private" {
		private = true
	}

	r, ok := s.rooms[roomName]
	if !ok {
		r = &room{
			name:    roomName,
			members: make(map[string]*client),
			private: private,
		}
		s.rooms[roomName] = r
		if private {
			c.privateRooms[roomName] = r
		}
	}

	if r.private {
		if _, exists := c.privateRooms[r.name]; !exists {
			c.err(errors.New("you are not allowed to enter"))
			return
		}
	}

	r.members[c.name] = c

	s.quitCurrentRoom(c)

	c.currentRoom = r

	r.broadcast(c, fmt.Sprintf("%s has joined the room", c.name))
	c.msg(fmt.Sprintf("welcome to %s", r.name))
}

func (s *server) invite(c *client, args []string) {
	if c.currentRoom == nil {
		c.err(errors.New("to invite a member, please enter a room"))
		return
	}

	memberName := args[1]
	invitedMember, ok := s.members[memberName]
	if !ok {
		c.err(errors.New("no such member exists"))
		return
	}
	if c.currentRoom.private {
		invitedMember.privateRooms[c.currentRoom.name] = c.currentRoom
	}
	c.msg(fmt.Sprintf("%s has been invited to %s", invitedMember.name, c.currentRoom.name))
	invitedMember.msg(fmt.Sprintf("you have been invited to %s", c.currentRoom.name))
}

func (s *server) listMembers(c *client, args []string) {
	var members []string
	for member := range s.members {
		members = append(members, member)
	}

	c.msg(fmt.Sprintf("available members are: %s", strings.Join(members, ", ")))
}

func (s *server) listRooms(c *client, args []string) {
	var publicRooms []string
	var privateRooms []string
	for roomName, room := range s.rooms {
		if room.private {
			if _, exists := c.privateRooms[roomName]; exists {
				privateRooms = append(privateRooms, roomName)
			}
		} else {
			publicRooms = append(publicRooms, roomName)
		}
	}

	c.msg(fmt.Sprintf("available public rooms are: %s", strings.Join(publicRooms, ", ")))
	c.msg(fmt.Sprintf("available private rooms are: %s", strings.Join(privateRooms, ", ")))
}

func (s *server) msg(c *client, args []string) {
	if c.currentRoom == nil {
		c.err(errors.New("you must join the room first"))
		return
	}

	c.currentRoom.broadcast(c, c.name+": "+strings.Join(args[1:len(args)], " "))
}

func (s *server) quit(c *client, args []string) {
	log.Printf("client has disconnected: %s", c.conn.RemoteAddr().String())

	s.quitCurrentRoom(c)

	c.msg("sad to see tou go :(")
	c.conn.Close()
}

func (s *server) quitCurrentRoom(c *client) {
	if c.currentRoom != nil {
		delete(c.currentRoom.members, c.name)
		c.currentRoom.broadcast(c, fmt.Sprintf("%s has left the room", c.name))
	}
}
