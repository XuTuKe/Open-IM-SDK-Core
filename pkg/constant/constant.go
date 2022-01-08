package constant

import (
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"open_im_sdk/internal/controller/conversation_msg"
	"open_im_sdk/internal/controller/friend"
	"open_im_sdk/internal/controller/group"
	"open_im_sdk/pkg/server_api_params"
	"open_im_sdk/pkg/utils"
	"sync"
)

func initAddr() {
	ginAddress = SvrConf.IpApiAddr
	getUserInfoRouter = ginAddress + "/user/get_user_info"
	updateUserInfoRouter = ginAddress + "/user/update_user_info"
	addFriendRouter = ginAddress + "/friend/add_friend"
	getFriendApplicationListRouter = ginAddress + "/friend/get_friend_apply_list"
	getSelfApplicationListRouter = ginAddress + "/friend/get_self_apply_list"
	deleteFriendRouter = ginAddress + "/friend/delete_friend"
	getFriendInfoRouter = ginAddress + "/friend/get_friends_info"
	getFriendListRouter = ginAddress + "/friend/get_friend_list"
	sendMsgRouter = ginAddress + "/chat/send_msg"
	getBlackListRouter = ginAddress + "/friend/get_blacklist"
	addFriendResponse = ginAddress + "/friend/add_friend_response"
	addBlackListRouter = ginAddress + "/friend/add_blacklist"
	removeBlackListRouter = ginAddress + "/friend/remove_blacklist"
	//getFriendApplyListRouter = ginAddress + "/friend/get_friend_apply_list"
	pullUserMsgRouter = ginAddress + "/chat/pull_msg"
	pullUserMsgBySeqRouter = ginAddress + "/chat/pull_msg_by_seq"

	newestSeqRouter = ginAddress + "/chat/newest_seq"
	setFriendComment = ginAddress + "/friend/set_friend_comment"
	tencentCloudStorageCredentialRouter = ginAddress + "/third/tencent_cloud_storage_credential"

	//group
	createGroupRouter = ginAddress + "/group/create_group"
	setGroupInfoRouter = ginAddress + "/group/set_group_info"
	joinGroupRouter = ginAddress + "/group/join_group"
	quitGroupRouter = ginAddress + "/group/quit_group"
	getGroupsInfoRouter = ginAddress + "/group/get_groups_info"
	getGroupMemberListRouter = ginAddress + "/group/get_group_member_list"
	getGroupAllMemberListRouter = ginAddress + "/group/get_group_all_member_list"
	getGroupMembersInfoRouter = ginAddress + "/group/get_group_members_info"
	inviteUserToGroupRouter = ginAddress + "/group/invite_user_to_group"
	getJoinedGroupListRouter = ginAddress + "/group/get_joined_group_list"
	kickGroupMemberRouter = ginAddress + "/group/kick_group"
	transferGroupRouter = ginAddress + "/group/transfer_group"
	getGroupApplicationListRouter = ginAddress + "/group/get_group_applicationList"
	acceptGroupApplicationRouter = ginAddress + "/group/group_application_response"
	refuseGroupApplicationRouter = ginAddress + "/group/group_application_response"
	//conversation
	setReceiveMessageOptRouter = ginAddress + "/conversation/set_receive_message_opt"
	getReceiveMessageOptRouter = ginAddress + "/conversation/get_receive_message_opt"
	getAllConversationMessageOptRouter = ginAddress + "/conversation/get_all_conversation_message_opt"

}

var (
	ginAddress                          = ""
	getUserInfoRouter                   = ""
	updateUserInfoRouter                = ""
	addFriendRouter                     = ""
	getFriendInfoRouter                 = ""
	getFriendApplicationListRouter      = ""
	getSelfApplicationListRouter        = ""
	deleteFriendRouter                  = ""
	getFriendListRouter                 = ""
	sendMsgRouter                       = ""
	getBlackListRouter                  = ""
	addFriendResponse                   = ""
	addBlackListRouter                  = ""
	removeBlackListRouter               = ""
	setFriendComment                    = " "
	pullUserMsgRouter                   = ""
	pullUserMsgBySeqRouter              = ""
	newestSeqRouter                     = ""
	tencentCloudStorageCredentialRouter = ""
	//group
	createGroupRouter                  = ""
	setGroupInfoRouter                 = ""
	joinGroupRouter                    = ""
	quitGroupRouter                    = ""
	getGroupsInfoRouter                = ""
	getGroupMemberListRouter           = ""
	getGroupAllMemberListRouter        = ""
	getGroupMembersInfoRouter          = ""
	inviteUserToGroupRouter            = ""
	getJoinedGroupListRouter           = ""
	kickGroupMemberRouter              = ""
	transferGroupRouter                = ""
	getGroupApplicationListRouter      = ""
	acceptGroupApplicationRouter       = ""
	refuseGroupApplicationRouter       = ""
	setReceiveMessageOptRouter         = ""
	getReceiveMessageOptRouter         = ""
	getAllConversationMessageOptRouter = ""
)

func (u *UserRelated) initListenerCh() {
	u.ch = make(chan utils.cmd2Value, 1000)
	u.ConversationCh = u.ch

	u.wsNotification = make(map[string]chan utils.GeneralWsResp, 1)
	u.seqMsg = make(map[int32]*server_api_params.MsgData, 1000)

	u.receiveMessageOpt = make(map[string]int32, 1000)
}

const (
	CmdFriend                     = "001"
	CmdBlackList                  = "002"
	CmdFriendApplication          = "003"
	CmdDeleteConversation         = "004"
	CmdNewMsgCome                 = "005"
	CmdGeyLoginUserInfo           = "006"
	CmdUpdateConversation         = "007"
	CmdForceSyncFriend            = "008"
	CmdFroceSyncBlackList         = "009"
	CmdForceSyncFriendApplication = "010"
	CmdForceSyncMsg               = "011"
	CmdForceSyncLoginUerInfo      = "012"
	CmdReLogin                    = "013"
	CmdUnInit                     = "014"
	CmdAcceptFriend               = "015"
	CmdRefuseFriend               = "016"
	CmdAddFriend                  = "017"
)

const (
	//ContentType
	Text           = 101
	Picture        = 102
	Voice          = 103
	Video          = 104
	File           = 105
	AtText         = 106
	Merger         = 107
	Card           = 108
	Location       = 109
	Custom         = 110
	Revoke         = 111
	HasReadReceipt = 112
	Typing         = 113
	Quote          = 114
	//////////////////////////////////////////
	SingleTipBegin             = 200
	AcceptFriendApplicationTip = 201
	AddFriendTip               = 202
	RefuseFriendApplicationTip = 203
	SetSelfInfoTip             = 204

	SingleTipEnd = 399
	/////////////////////////////////////////
	GroupTipBegin             = 500
	TransferGroupOwnerTip     = 501
	CreateGroupTip            = 502
	JoinGroupTip              = 504
	QuitGroupTip              = 505
	SetGroupInfoTip           = 506
	AcceptGroupApplicationTip = 507
	RefuseGroupApplicationTip = 508
	KickGroupMemberTip        = 509
	InviteUserToGroupTip      = 510

	GroupTipEnd = 599
	////////////////////////////////////////
	//MsgFrom
	UserMsgType = 100
	SysMsgType  = 200

	/////////////////////////////////////
	//SessionType
	SingleChatType = 1
	GroupChatType  = 2

	//MsgStatus
	MsgStatusSending     = 1
	MsgStatusSendSuccess = 2
	MsgStatusSendFailed  = 3
	MsgStatusHasDeleted  = 4
	MsgStatusRevoked     = 5
	MsgStatusFiltered    = 6

	//OptionsKey
	IsHistory            = "history"
	IsPersistent         = "persistent"
	IsUnreadCount        = "unreadCount"
	IsConversationUpdate = "conversationUpdate"
)

const (
	ckWsInitConnection  string = "ws-init-connection"
	ckWsLoginConnection string = "ws-login-connection"
	ckWsClose           string = "ws-close"
	ckWsKickOffLine     string = "ws-kick-off-line"
	ckTokenExpired      string = "token-expired"
	ckSelfInfoUpdate    string = "self-info-update"
)

const (
	ErrCodeInitLogin    = 1001
	ErrCodeFriend       = 2001
	ErrCodeConversation = 3001
	ErrCodeUserInfo     = 4001
	ErrCodeGroup        = 5001
)

const (
	SdkInit      = 0
	LoginSuccess = 101
	Logining     = 102
	LoginFailed  = 103

	LogoutCmd = 201

	TokenFailedExpired       = 701
	TokenFailedInvalid       = 702
	TokenFailedKickedOffline = 703
)

const (
	DeFaultSuccessMsg = "ok"
)

const (
	AddConOrUpLatMsg          = 2
	UnreadCountSetZero        = 3
	IncrUnread                = 5
	TotalUnreadMessageChanged = 6
	UpdateFaceUrlAndNickName  = 7
	UpdateLatestMessageChange = 8
	NewConChange              = 9
	NewCon                    = 10

	HasRead = 1
	NotRead = 0

	IsFilter  = 1
	NotFilter = 0

	Pinned    = 1
	NotPinned = 0
)

const (
	GroupActionCreateGroup            = 1
	GroupActionApplyJoinGroup         = 2
	GroupActionQuitGroup              = 3
	GroupActionSetGroupInfo           = 4
	GroupActionKickGroupMember        = 5
	GroupActionTransferGroupOwner     = 6
	GroupActionInviteUserToGroup      = 7
	GroupActionAcceptGroupApplication = 8
	GroupActionRefuseGroupApplication = 9
)
const ZoomScale = "200"
const MaxTotalMsgLen = 2048
const (
	FriendAcceptTip  = "You have successfully become friends, so start chatting"
	TransferGroupTip = "The owner of the group is transferred!"
	AcceptGroupTip   = "%s join the group"
)

const (
	WSGetNewestSeq     = 1001
	WSPullMsg          = 1002
	WSSendMsg          = 1003
	WSPullMsgBySeqList = 1004
	WSPushMsg          = 2001
	WSKickOnlineMsg    = 2002
	WSDataError        = 3001
)

const (
	//MsgReceiveOpt
	ReceiveMessage          = 0
	NotReceiveMessage       = 1
	ReceiveNotNotifyMessage = 2
)

// key = errCode, string = errMsg
type ErrInfo struct {
	ErrCode int32
	ErrMsg  string
}

var (
	OK = ErrInfo{0, ""}

	ErrParseToken = ErrInfo{200, ParseTokenMsg.Error()}

	ErrTencentCredential = ErrInfo{400, ThirdPartyMsg.Error()}

	ErrTokenExpired     = ErrInfo{701, TokenExpiredMsg.Error()}
	ErrTokenInvalid     = ErrInfo{702, TokenInvalidMsg.Error()}
	ErrTokenMalformed   = ErrInfo{703, TokenMalformedMsg.Error()}
	ErrTokenNotValidYet = ErrInfo{704, TokenNotValidYetMsg.Error()}
	ErrTokenUnknown     = ErrInfo{705, TokenUnknownMsg.Error()}

	ErrAccess = ErrInfo{ErrCode: 801, ErrMsg: AccessMsg.Error()}
	ErrDB     = ErrInfo{ErrCode: 802, ErrMsg: DBMsg.Error()}
	ErrArgs   = ErrInfo{ErrCode: 803, ErrMsg: ArgsMsg.Error()}
	ErrApi    = ErrInfo{ErrCode: 804, ErrMsg: ApiMsg.Error()}
)

var (
	ParseTokenMsg       = errors.New("parse token failed")
	TokenExpiredMsg     = errors.New("token is timed out, please log in again")
	TokenInvalidMsg     = errors.New("token has been invalidated")
	TokenNotValidYetMsg = errors.New("token not active yet")
	TokenMalformedMsg   = errors.New("that's not even a token")
	TokenUnknownMsg     = errors.New("couldn't handle this token")

	AccessMsg = errors.New("no permission")
	DBMsg     = errors.New("db failed")
	ArgsMsg   = errors.New("args failed")
	ApiMsg    = errors.New("api failed")

	ThirdPartyMsg = errors.New("third party error")
)

func (e *ErrInfo) Error() string {
	return e.ErrMsg
}

const SuccessCallbackDefault = ""

var UserSDKRwLock sync.RWMutex
var UserRouterMap map[string]*UserRelated
var SvrConf utils.IMConfig
var SdkLogFlag int32
var hearbeatInterval int32 = 5
