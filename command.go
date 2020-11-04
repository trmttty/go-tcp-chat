package main

type commandID int

const (
	NAME commandID = iota
	RENAME
	CREATE
	JOIN
	INVITE
	MEMBERS
	ROOMS
	MSG
	QUIT
)

type command struct {
	id     commandID
	client *client
	args   []string
}
