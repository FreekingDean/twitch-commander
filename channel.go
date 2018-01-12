package twitch_commander

type Channel struct {
	name string

	connected bool
}

// NewChannel builds a new Channel with the given name.
func NewChannel(name string) *Channel {
	return &Channel{
		name:      name,
		connected: false,
	}
}
