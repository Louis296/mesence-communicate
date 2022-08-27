package communicate_service

type OnlineMessageData struct {
	Users []string
}

type OfflineMessageData struct {
	Users []string
}

type WordMessage struct {
	Type string
	Data WordMessageData
}

type FriendRequestMessage struct {
	Type string
	Data FriendRequestData
}

type WordMessageData struct {
	To       string
	From     string
	Content  string
	SendTime string
}

type FriendRequestData struct {
	Sender        string
	Candidate     string
	Content       string
	StartTime     string
	RequestStatus string
}
