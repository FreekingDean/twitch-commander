package twitch_commander

import (
	"net"
)

const (
	twitch_username_regex_string string = "[a-zA-Z0-9_]+"
)

// TwitchCommander is the base object to interact
// with Twitch chat.
type TwitchCommander struct {
	conn    net.Conn
	scanner *bufio.Scanner
}

// NewTwitchCommander builds a TwitchCommander object the OAuthToken, and Nickname.
// You can retreive your OAuth token here https://twitchapps.com/tmi.
func NewTwitchCommander() *TwitchCommander {
	return &TwitchCommander{
	//messageMatcher: regexp.MustCompile(fmt.Sprintf(PRIVMSG_REGEX, channelName)),
	}
}

// ConnectAndListen will call the Connect(), Authenticate(oAuthToken, nickname),
// and finally Listen() methods. The Listen() method is a
// blocking call. This will also `defer Close()`.
func (tc *TwitchCommander) ConnectAndListen(oAuthToken, nickname string) error {
	if err := tc.Connect(); err != nil {
		return err
	}
	defer tc.Close()
	if err := tc.Authenticate(); err != nil {
		return err
	}
	return tc.Listen()
}

// Connect will connect the commander to the Twitch IRC
// host. It will return an error if connection is unsuccessful.
func (tc *TwitchCommander) Connect() error {
	if tc.conn != nil {
		return errors.New("Already Connected!")
	}
	conn, err := net.Dial("tcp", twitch_hostname)
	if err != nil {
		return err
	}
	tc.conn = conn
	tc.scanner = bufio.NewScanner(conn)
	return nil
}

// Close will close the Twitch IRC connection.
func (tc *TwitchCommander) Close() error {
	if tc.conn == nil {
		return errors.New("Not Connected!")
	}
	if err := tc.conn.Close(); err != nil {
		return err
	}
	tc.conn = nil
	return nil
}

// Authenticate will send the valid authentication messages
// to the Twitch IRC described here https://dev.twitch.tv/docs/irc#connecting-to-twitch-irc.
func (tc *TwitchCommander) Authenticate(oAuthToken, nickname string) error {
	if err := tc.send("PASS oauth:", oAuthToken); err != nil {
		return err
	}
	if err := tc.send("NICK ", nickname); err != nil {
		return err
	}
	return nil
}

// Listen is a blocking call that will listen to all available
// messages from the Twitch IRC connection.
func (tc *TwitchCommander) Listen() error {
	for tc.scanner.Scan() {
		line := scanner.Text()
		for _, handler := range lineHandlers {
			err, handled := handler(tc, line)
			if err != nil {
				return err
			}
			if handled {
				break
			}
		}
	}
	return tc.scanner.Error()
}

// AddChannel will add a channel for TwitchCommander to join
// and listen to. If there is an active connection TwitchCommander
// will join this channel immediatly. If there is not an active
// connection TwitchCommander will join on ConnectAndListen().
func (tc *TwitchCommander) AddChannel(c *Channel) error {
	if _, ok := tc.channels[c.name]; ok {
		return errors.New("Channel exists")
	}
	tc.channels[c.name] = c
	return tc.JoinChannel(c)
}

// RemoveChannel will remove a channel that TwitchCommander knows
// about. If there is an active connection, and TwitchCommander has
// joined this channel TwitchCommander will leave this channel before
// removal.
func (tc *TwitchCommander) RemoveChannel(c *Channel) {
	if channel, ok := tc.channels[c.name]; ok {
		if err := tc.Send("LEAVE #", tc.name); err != nil {
			delete(tc.channels, c.name)
		} else {
			return err
		}
	} else {
		return errors.New("Cant find channel")
	}
}

// Sends the JOIN command with the channel name to the Twitch
// IRC.
func (tc *TwitchCommander) JoinChannel(c *Channel) error {
	if channel, ok := tc.channels[c.name]; ok {
		return tc.Send("JOIN #", tc.name)
	} else {
		return errors.New("Add channel before joining")
	}
}

// Sends the LEAVE command with the channel name to the Twitch
// IRC.
func (tc *TwitchCommander) LeaveChannel(c *Channel) {
	if channel, ok := tc.channels[c.name]; ok {
	} else {
		return errors.New("Cant find channel")
	}
}

func (tc *TwitchCommander) send(message string, args ...interface{}) error {
	_, err := fmt.Fprintf(tc.conn, message, args...)
	return err
}

func (tc *TwitchCommander) receivedMessage(m *Message) error {
}
