package pkg

type Sender interface {
	TwitchAccount() TwitchAccount

	Say(Channel, string)
	Mention(Channel, User, string)
	Whisper(User, string)
	Timeout(Channel, User, int, string)
	Ban(Channel, User, string)

	GetPoints(Channel, string) uint64

	// give or remove points from user in channel
	BulkEdit(string, []string, int32)

	AddPoints(Channel, string, uint64) (bool, uint64)
	RemovePoints(Channel, string, uint64) (bool, uint64)
	ForceRemovePoints(Channel, string, uint64) uint64

	PointRank(Channel, string) uint64

	GetUserStore() UserStore
	GetUserContext() UserContext

	MakeUser(string) User
	MakeChannel(string) Channel

	// Permanently join channel with the given channel ID
	JoinChannel(channelID string) error

	// Permanently leave channel with the given channel ID
	LeaveChannel(channelID string) error
}
