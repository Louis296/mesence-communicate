package enum

const (
	HeartPackageMessageType  = "heartPackage"
	WordPackageMessageType   = "word"
	OnlineMessageType        = "online"
	OfflineMessageType       = "offline"
	FriendRequestMessageType = "friendRequest"
)

var FriendRequestStatusMap = map[int]string{
	0: "waiting",
	1: "accepted",
	2: "refused",
}
