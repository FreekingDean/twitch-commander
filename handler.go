package twitch_commander

const (
	privmsg_regex_string string = fmt.Sprintf(`^:%s!%s@(%s)\.tmi\.twitch\.tv\ PRIVMSG\ #(%s)\ :(.*)$`, twitch_username_regex_string, twitch_username_regex_string, twitch_username_regex_string, twitch_username_regex_string)
)

type handlerFunc func(*TwitchCommander, string) (bool, error)

var (
	privmsg_regex *regexp.Regexp = regexp.MustCompile(privmsg_regex_string)
	lineHandlers  []handlerFunc  = []handlerFunc{handlePING, handlePRIVMSG}
)

func handlePING(tc *TwitchCommander, line string) (bool, error) {
	if line == "PING :tmi.twitch.tv" {
		return true, tc.send("PONG :tmi.twitch.tv")
	}
	return false
}

func handlePRIVMSG(tc *TwitchCommander, line string) bool {
	matches := privmsg_regex.FindStringSubmatch(line)
	if len(matches) > 0 {
		message := NewMessage(matches[1], matches[2], matches[3])
		return true, tc.receivedMessage(message)
	}
	return false
}
