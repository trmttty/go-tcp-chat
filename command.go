package main

type commandID int

const (
	CmdNick commandID = iota
	CmdJoin
	CmdInvite
	CmdMembers
	CmdRooms
	CmdMsg
	CmdQuit
)

type command struct {
	id     commandID
	client *client
	args   []string
}
