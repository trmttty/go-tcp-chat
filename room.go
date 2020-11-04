package main

type room struct {
	name    string
	members map[string]*client
	private bool
}

func (r *room) broadcast(sender *client, msg string) {
	for name, m := range r.members {
		if name != sender.name {
			m.msg(msg)
		}
	}
}
