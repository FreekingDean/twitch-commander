package twitch_commander

type TwitchCommander struct {
	oauthToken string
	nickname   string
}

const (
	PRIVMSG_REGEX string = "^:[a-zA-Z0-9_]+![a-zA-Z0-9_]+@[a-zA-Z0-9_]+\\.tmi\\.twitch\\.tv\\ PRIVMSG\\ #%s\\ :(.*)$"
)

func NewTwitchCommander(oauthToken, nickname, channleName string) *TwitchCommander {
	return &TwitchCommander{
		oauthToken:  oauthToken,
		nickname:    nickname,
		channelName: channelName,

		messageMatcher: regexp.MustCompile(fmt.Sprintf(PRIVMSG_REGEX, channelName)),
	}
}
