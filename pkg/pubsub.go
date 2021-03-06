package pkg

type PubSubAuthorization struct {
	Nonce        string
	TwitchUserID string
	admin        bool
}

func PubSubAdminAuth() *PubSubAuthorization {
	return &PubSubAuthorization{
		admin: true,
	}
}

func (p PubSubAuthorization) Admin() bool {
	return p.admin
}

type PubSubBan struct {
	Channel string
	Target  string
	Reason  string
}

type PubSubTimeout struct {
	Channel  string
	Target   string
	Reason   string
	Duration uint32
}

type PubSubUntimeout struct {
	Channel string
	Target  string
}

type PubSubUser struct {
	ID   string
	Name string
}

type PubSubBanEvent struct {
	Channel PubSubUser
	Target  PubSubUser
	Source  PubSubUser
	Reason  string
}

type PubSubTimeoutEvent struct {
	Channel  PubSubUser
	Target   PubSubUser
	Source   PubSubUser
	Duration int
	Reason   string
}
