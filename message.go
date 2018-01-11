package twitch_commander

type Message struct {
	Channel  string
	Username string
	Body     string
}

var privmsgRegexp *regexp.Regexp = regexp.MustCompile(`^:[a-zA-Z0-9_]+![a-zA-Z0-9_]+@([a-zA-Z0-9_]+)\.tmi\.twitch\.tv\ PRIVMSG\ #(.*)\ :(.*)$`)

func parseMessage(message) *Message {
}
