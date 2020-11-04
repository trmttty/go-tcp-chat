package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn         net.Conn
	name         string
	currentRoom  *room
	privateRooms map[string]*room
	commands     chan<- command
}

func newClient(conn net.Conn, command chan<- command) *client {
	return &client{
		conn:         conn,
		privateRooms: make(map[string]*room),
		commands:     command,
	}
}

func (c *client) readInput() {
	c.msg("welcom! plese enter your nickname")
	reader := bufio.NewReader(c.conn)
	name, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	name = strings.Trim(name, "\r\n")
	c.commands <- command{
		id:     CmdNick,
		client: c,
		args:   []string{name},
	}

	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/rename":
			c.commands <- command{
				id:     CmdRename,
				client: c,
				args:   args,
			}
		case "/create":
			c.commands <- command{
				id:     CmdCreate,
				client: c,
				args:   args,
			}
		case "/join":
			c.commands <- command{
				id:     CmdJoin,
				client: c,
				args:   args,
			}
		case "/invite":
			c.commands <- command{
				id:     CmdInvite,
				client: c,
				args:   args,
			}
		case "/members":
			c.commands <- command{
				id:     CmdMembers,
				client: c,
				args:   args,
			}
		case "/rooms":
			c.commands <- command{
				id:     CmdRooms,
				client: c,
				args:   args,
			}
		case "/msg":
			c.commands <- command{
				id:     CmdMsg,
				client: c,
				args:   args,
			}
		case "/quit":
			c.commands <- command{
				id:     CmdQuit,
				client: c,
				args:   args,
			}
		default:
			c.err(fmt.Errorf("unknown command: %s", cmd))
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("ERR: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}
