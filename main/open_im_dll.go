package main

//go build -buildmode=c-archive -o openim.lib ./main/open_im_dll.go

/*
#cgo LDFLAGS: -Wl,--unresolved-symbols=ignore-in-object-files -Wl,--allow-shlib-undefined

#include <stdlib.h>

typedef void (*FunWithVoid)();
typedef void (*FunWithInt)(int);
typedef void (*FunWithString)(char *);
typedef void (*FunWithIntString)(int, char *);

static inline void CallFunWithVoid(void* cb) {
    return (*(FunWithVoid)cb)();
}

static inline void CallFunWithInt(void* cb, int i) {
    return (*(FunWithVoid)cb)(i);
}

static inline void CallFunWithString(void* cb, char * s) {
    return (*(FunWithVoid)cb)(s);
}

static inline void CallFunWithIntString(void* cb, int i, char * s) {
    return (*(FunWithVoid)cb)(i, s);
}

*/
import "C"
import (
	"open_im_sdk/internal/login"
	openIM "open_im_sdk/open_im_sdk"
	"open_im_sdk/sdk_struct"
	"unsafe"
)


func main() {

}

func FunWithVoid(cb unsafe.Pointer) {
	C.CallFunWithVoid(cb)
}

func FunWithInt(cb unsafe.Pointer, i int) {
	cint := C.int(i)
	C.CallFunWithInt(cb, cint)
}

func FunWithString(cb unsafe.Pointer, str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.CallFunWithString(cb, cstr)
}

func FunWithIntString(cb unsafe.Pointer, i int, str string) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))
	C.CallFunWithIntString(cb, C.int(i), cstr)
}

//export SdkVersion
func SdkVersion(verCb unsafe.Pointer) {
	ver := openIM.SdkVersion()
	FunWithString(verCb, ver)
}

//export InitSDK
func InitSDK(onConnecting unsafe.Pointer, onConnectSuccess unsafe.Pointer,
	onConnectFailed unsafe.Pointer, onKickedOffline unsafe.Pointer,
	onUserTokenExpired unsafe.Pointer, operationID *C.char, config *C.char) bool {

	return openIM.InitSDK(&OnConnListener{
		onConnecting: func() {
			FunWithVoid(onConnecting)
		},
		onConnectSuccess: func() {
			FunWithVoid(onConnectSuccess)
		},
		onConnectFailed: func(errCode int, errMsg string) {
			FunWithIntString(onConnectFailed, int(errCode), errMsg)
		},
		onKickedOffline: func() {
			FunWithVoid(onKickedOffline)
		},
		onUserTokenExpired: func() {
			FunWithVoid(onUserTokenExpired)
		},
	},
		C.GoString(operationID),
		C.GoString(config),
	)
}
//export Login
func Login(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, userID, token *C.char) {
	openIM.Login(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(userID), C.GoString(token))
}
//export UploadImage
func UploadImage(onReturn unsafe.Pointer, onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, filePath *C.char, token, obj *C.char) {
	ret := openIM.UploadImage(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(filePath), C.GoString(token), C.GoString(obj))

	FunWithString(onReturn, ret)
}
//export Logout
func Logout(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char) {
	openIM.Logout(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID))
}
//export GetLoginStatus
func GetLoginStatus() int32 {
	return openIM.GetLoginStatus()
}
//export GetLoginStatus
func GetLoginUser(OnReturn unsafe.Pointer) {
	ret := openIM.GetLoginUser()
	FunWithString(OnReturn, ret)
}

///////////////////////user/////////////////////
//export GetLoginStatus
func GetUsersInfo(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, userIDList *C.char) {
	openIM.GetUsersInfo(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(userIDList))
}
//export GetLoginStatus
func SetSelfInfo(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, userInfo *C.char) {
	openIM.SetSelfInfo(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(userInfo))
}
//export GetLoginStatus
func GetSelfUserInfo(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char) {
	openIM.GetSelfUserInfo(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID))
}

//////////////////////////group//////////////////////////////////////////
//export GetLoginStatus
func SetGroupListener(
	onJoinedGroupAdded unsafe.Pointer,
	onJoinedGroupDeleted unsafe.Pointer,
	onGroupMemberAdded unsafe.Pointer,
	onGroupMemberDeleted unsafe.Pointer,
	onGroupApplicationAdded unsafe.Pointer,
	onGroupApplicationDeleted unsafe.Pointer,
	onGroupInfoChanged unsafe.Pointer,
	onGroupMemberInfoChanged unsafe.Pointer,
	onGroupApplicationAccepted unsafe.Pointer,
	onGroupApplicationRejected unsafe.Pointer,
) {
	openIM.SetGroupListener(&OnGroupListener{
		onJoinedGroupAdded: func(s string) {
			FunWithString(onJoinedGroupAdded, s)
		},
		onJoinedGroupDeleted: func(s string) {
			FunWithString(onJoinedGroupDeleted, s)
		},
		onGroupMemberAdded: func(s string) {
			FunWithString(onGroupMemberAdded, s)
		},
		onGroupMemberDeleted: func(s string) {
			FunWithString(onGroupMemberDeleted, s)
		},
		onGroupApplicationAdded: func(s string) {
			FunWithString(onGroupApplicationAdded, s)
		},
		onGroupApplicationDeleted: func(s string) {
			FunWithString(onGroupApplicationDeleted, s)
		},
		onGroupInfoChanged: func(s string) {
			FunWithString(onGroupInfoChanged, s)
		},
		onGroupMemberInfoChanged: func(s string) {
			FunWithString(onGroupMemberInfoChanged, s)
		},
		onGroupApplicationAccepted: func(s string) {
			FunWithString(onGroupApplicationAccepted, s)
		},
		onGroupApplicationRejected: func(s string) {
			FunWithString(onGroupApplicationRejected, s)
		},
	})
}
//export GetLoginStatus
func CreateGroup(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupBaseInfo *C.char, memberList *C.char) {
	openIM.CreateGroup(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupBaseInfo), C.GoString(memberList))
}
//export GetLoginStatus
func JoinGroup(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupID, reqMsg *C.char) {
	openIM.JoinGroup(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupID), C.GoString(reqMsg))
}
//export GetLoginStatus
func QuitGroup(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupID *C.char) {
	openIM.QuitGroup(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupID))
}
//export GetLoginStatus
func GetJoinedGroupList(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char) {
	openIM.GetJoinedGroupList(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID))
}
//export GetLoginStatus
func GetGroupsInfo(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupIDList *C.char) {
	openIM.GetGroupsInfo(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupIDList))
}
//export GetLoginStatus
func SetGroupInfo(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupID *C.char, groupInfo *C.char) {
	openIM.SetGroupInfo(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupID), C.GoString(groupInfo))
}
//export GetLoginStatus
func GetGroupMemberList(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupID *C.char, filter, offset, count int32) {
	openIM.GetGroupMemberList(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupID), filter, offset, count)
}
//export GetLoginStatus
func GetGroupMembersInfo(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupID *C.char, userIDList *C.char) {
	openIM.GetGroupMembersInfo(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupID), C.GoString(userIDList))
}
//export GetLoginStatus
func KickGroupMember(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupID *C.char, reason *C.char, userIDList *C.char) {
	openIM.KickGroupMember(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupID), C.GoString(reason), C.GoString(userIDList))
}
//export GetLoginStatus
func TransferGroupOwner(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupID, newOwnerUserID *C.char) {
	openIM.TransferGroupOwner(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupID), C.GoString(newOwnerUserID))
}
//export GetLoginStatus
func InviteUserToGroup(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupID, reason *C.char, userIDList *C.char) {
	openIM.InviteUserToGroup(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupID), C.GoString(reason), C.GoString(userIDList))
}
//export GetLoginStatus
func GetRecvGroupApplicationList(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char) {
	openIM.GetRecvGroupApplicationList(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID))
}
//export GetLoginStatus
func GetSendGroupApplicationList(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char) {
	openIM.GetSendGroupApplicationList(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID))
}
//export GetLoginStatus
func AcceptGroupApplication(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupID, fromUserID, handleMsg *C.char) {
	openIM.AcceptGroupApplication(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupID), C.GoString(fromUserID), C.GoString(handleMsg))
}
//export GetLoginStatus
func RefuseGroupApplication(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupID, fromUserID, handleMsg *C.char) {
	openIM.RefuseGroupApplication(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupID), C.GoString(fromUserID), C.GoString(handleMsg))
}

////////////////////////////friend/////////////////////////////////////
//export GetLoginStatus
func GetDesignatedFriendsInfo(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, userIDList *C.char) {
	openIM.GetDesignatedFriendsInfo(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(userIDList))
}
//export GetLoginStatus
func GetFriendList(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char) {
	openIM.GetFriendList(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID))
}
//export GetLoginStatus
func CheckFriend(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, userIDList *C.char) {
	openIM.CheckFriend(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(userIDList))
}
//export GetLoginStatus
func AddFriend(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, userIDReqMsg *C.char) {
	openIM.AddFriend(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(userIDReqMsg))
}
//export SetFriendRemark
func SetFriendRemark(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, userIDRemark *C.char) {
	openIM.SetFriendRemark(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(userIDRemark))
}
//export DeleteFriend
func DeleteFriend(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, friendUserID *C.char) {
	openIM.DeleteFriend(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(friendUserID))
}
//export GetRecvFriendApplicationList
func GetRecvFriendApplicationList(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char) {
	openIM.GetRecvFriendApplicationList(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID))
}
//export GetSendFriendApplicationList
func GetSendFriendApplicationList(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char) {
	openIM.GetSendFriendApplicationList(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID))
}
//export AcceptFriendApplication
func AcceptFriendApplication(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, userIDHandleMsg *C.char) {
	openIM.AcceptFriendApplication(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(userIDHandleMsg))
}
//export RefuseFriendApplication
func RefuseFriendApplication(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, userIDHandleMsg *C.char) {
	openIM.RefuseFriendApplication(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(userIDHandleMsg))
}
//export AddBlack
func AddBlack(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, blackUserID *C.char) {
	openIM.AddBlack(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(blackUserID))
}
//export GetBlackList
func GetBlackList(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char) {
	openIM.GetBlackList(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID))
}
//export RemoveBlack
func RemoveBlack(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, removeUserID *C.char) {
	openIM.RemoveBlack(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(removeUserID))
}
//export SetFriendListener
func SetFriendListener(
	onFriendApplicationAdded unsafe.Pointer,
	onFriendApplicationDeleted unsafe.Pointer,
	onFriendApplicationAccepted unsafe.Pointer,
	onFriendApplicationRejected unsafe.Pointer,
	onFriendAdded unsafe.Pointer,
	onFriendDeleted unsafe.Pointer,
	onFriendInfoChanged unsafe.Pointer,
	onBlackAdded unsafe.Pointer,
	onBlackDeleted unsafe.Pointer,
) {
	openIM.SetFriendListener(&OnFriendshipListener{
		onFriendApplicationAdded: func(s string) {
			FunWithString(onFriendApplicationAdded, s)
		},
		onFriendApplicationDeleted: func(s string) {
			FunWithString(onFriendApplicationDeleted, s)
		},
		onFriendApplicationAccepted: func(s string) {
			FunWithString(onFriendApplicationAccepted, s)
		},
		onFriendApplicationRejected: func(s string) {
			FunWithString(onFriendApplicationRejected, s)
		},
		onFriendAdded: func(s string) {
			FunWithString(onFriendAdded, s)
		},
		onFriendDeleted: func(s string) {
			FunWithString(onFriendDeleted, s)
		},
		onFriendInfoChanged: func(s string) {
			FunWithString(onFriendInfoChanged, s)
		},
		onBlackAdded: func(s string) {
			FunWithString(onBlackAdded, s)
		},
		onBlackDeleted: func(s string) {
			FunWithString(onBlackDeleted, s)
		},
	})
}

///////////////////////conversation////////////////////////////////////
//export GetAllConversationList
func GetAllConversationList(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char) {
	openIM.GetAllConversationList(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID))
}
//export GetConversationListSplit
func GetConversationListSplit(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, offset, count int) {
	openIM.GetConversationListSplit(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), offset, count)
}
//export SetConversationRecvMessageOpt
func SetConversationRecvMessageOpt(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, conversationIDList *C.char, opt int) {
	openIM.SetConversationRecvMessageOpt(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(conversationIDList), opt)
}
//export GetConversationRecvMessageOpt
func GetConversationRecvMessageOpt(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, conversationIDList *C.char) {
	openIM.GetConversationRecvMessageOpt(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(conversationIDList))
}
//export GetOneConversation
func GetOneConversation(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, sessionType int, sourceID *C.char) {
	openIM.GetOneConversation(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), sessionType, C.GoString(sourceID))
}
//export GetLoginStatus
func GetMultipleConversation(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, conversationIDList *C.char) {
	openIM.GetMultipleConversation(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(conversationIDList))
}

//export DeleteConversation
func DeleteConversation(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, conversationID *C.char) {
	openIM.DeleteConversation(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(conversationID))
}
//export SetConversationDraft
func SetConversationDraft(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, conversationID, draftText *C.char) {
	openIM.SetConversationDraft(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(conversationID), C.GoString(draftText))
}
//export PinConversation
func PinConversation(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, conversationID *C.char, isPinned bool) {
	openIM.PinConversation(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(conversationID), isPinned)
}
//export GetTotalUnreadMsgCount
func GetTotalUnreadMsgCount(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char) {
	openIM.GetTotalUnreadMsgCount(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID))
}

//export SetConversationListener
func SetConversationListener(
	onSyncServerStart unsafe.Pointer,
	onSyncServerFinish unsafe.Pointer,
	onSyncServerFailed unsafe.Pointer,
	onNewConversation unsafe.Pointer,
	onConversationChanged unsafe.Pointer,
	onTotalUnreadMessageCountChanged unsafe.Pointer,
) {
	openIM.SetConversationListener(&OnConversationListener{
		onSyncServerStart:  func() { FunWithVoid(onSyncServerStart) },
		onSyncServerFinish: func() { FunWithVoid(onSyncServerFinish) },
		onSyncServerFailed: func() {
			FunWithVoid(onSyncServerFailed)
		},
		onNewConversation: func(s string) {
			FunWithString(onNewConversation, s)
		},
		onConversationChanged: func(s string) {
			FunWithString(onConversationChanged, s)
		},
		onTotalUnreadMessageCountChanged: func(i int) {
			FunWithInt(onTotalUnreadMessageCountChanged, i)
		},
	})
}
//export SetAdvancedMsgListener
func SetAdvancedMsgListener(
	onRecvNewMessage unsafe.Pointer,
	onRecvC2CReadReceipt unsafe.Pointer,
	onRecvMessageRevoked unsafe.Pointer,
) {
	openIM.SetAdvancedMsgListener(
		&OnAdvancedMsgListener{
			onRecvNewMessage: func(s string) {
				FunWithString(onRecvNewMessage, s)
			},
			onRecvC2CReadReceipt: func(s string) {
				FunWithString(onRecvC2CReadReceipt, s)
			},
			onRecvMessageRevoked: func(s string) {
				FunWithString(onRecvMessageRevoked, s)
			},
		},
	)
}
//export SetUserListener
func SetUserListener(onSelfInfoUpdated unsafe.Pointer) {
	openIM.SetUserListener(&OnUserListener{
		onSelfInfoUpdated: func(s string) {
			FunWithString(onSelfInfoUpdated, s)
		},
	})
}
//export CreateTextAtMessage
func CreateTextAtMessage(onReturn unsafe.Pointer, operationID *C.char, text, atUserList *C.char) {
	ret := openIM.CreateTextAtMessage(C.GoString(operationID), C.GoString(text), C.GoString(atUserList))
	FunWithString(onReturn, ret)
}

//export CreateTextMessage
func CreateTextMessage(onReturn unsafe.Pointer, operationID *C.char, text *C.char) {
	ret := openIM.CreateTextMessage(C.GoString(operationID), C.GoString(text))
	FunWithString(onReturn, ret)
}
//export CreateLocationMessage
func CreateLocationMessage(onReturn unsafe.Pointer, operationID *C.char, description *C.char, longitude, latitude float64) {
	ret := openIM.CreateLocationMessage(C.GoString(operationID), C.GoString(description), longitude, latitude)
	FunWithString(onReturn, ret)
}
//export CreateCustomMessage
func CreateCustomMessage(onReturn unsafe.Pointer, operationID *C.char, data, extension *C.char, description *C.char) {
	ret := openIM.CreateCustomMessage(C.GoString(operationID), C.GoString(data), C.GoString(extension), C.GoString(description))
	FunWithString(onReturn, ret)
}
//export CreateQuoteMessage
func CreateQuoteMessage(onReturn unsafe.Pointer, operationID *C.char, text *C.char, message *C.char) {
	ret := openIM.CreateQuoteMessage(C.GoString(operationID), C.GoString(text), C.GoString(message))
	FunWithString(onReturn, ret)
}
//export CreateCardMessage
func CreateCardMessage(onReturn unsafe.Pointer, operationID *C.char, cardInfo *C.char) {
	ret := openIM.CreateCardMessage(C.GoString(operationID), C.GoString(cardInfo))
	FunWithString(onReturn, ret)
}
//export CreateVideoMessageFromFullPath
func CreateVideoMessageFromFullPath(onReturn unsafe.Pointer, operationID *C.char, videoFullPath *C.char, videoType *C.char, duration int64, snapshotFullPath *C.char) {
	ret := openIM.CreateVideoMessageFromFullPath(C.GoString(operationID), C.GoString(videoFullPath), C.GoString(videoType), duration, C.GoString(snapshotFullPath))
	FunWithString(onReturn, ret)
}
//export CreateImageMessageFromFullPath
func CreateImageMessageFromFullPath(onReturn unsafe.Pointer, operationID *C.char, imageFullPath *C.char) {
	ret := openIM.CreateImageMessageFromFullPath(C.GoString(operationID), C.GoString(imageFullPath))
	FunWithString(onReturn, ret)
}
//export CreateSoundMessageFromFullPath
func CreateSoundMessageFromFullPath(onReturn unsafe.Pointer, operationID *C.char, soundPath *C.char, duration int64) {
	ret := openIM.CreateSoundMessageFromFullPath(C.GoString(operationID), C.GoString(soundPath), duration)
	FunWithString(onReturn, ret)
}
//export CreateFileMessageFromFullPath
func CreateFileMessageFromFullPath(onReturn unsafe.Pointer, operationID *C.char, fileFullPath, fileName *C.char) {
	ret := openIM.CreateFileMessageFromFullPath(C.GoString(operationID), C.GoString(fileFullPath), C.GoString(fileName))
	FunWithString(onReturn, ret)
}
//export CreateImageMessage
func CreateImageMessage(onReturn unsafe.Pointer, operationID *C.char, imagePath *C.char) {
	ret := openIM.CreateImageMessage(C.GoString(operationID), C.GoString(imagePath))
	FunWithString(onReturn, ret)
}
//export CreateImageMessageByURL
func CreateImageMessageByURL(onReturn unsafe.Pointer, operationID *C.char, sourcePicture, bigPicture, snapshotPicture *C.char) {
	ret := openIM.CreateImageMessageByURL(C.GoString(operationID), C.GoString(sourcePicture), C.GoString(bigPicture), C.GoString(snapshotPicture))
	FunWithString(onReturn, ret)
}
//export CreateSoundMessageByURL
func CreateSoundMessageByURL(onReturn unsafe.Pointer, operationID *C.char, soundBaseInfo *C.char) {
	ret := openIM.CreateSoundMessageByURL(C.GoString(operationID), C.GoString(soundBaseInfo))
	FunWithString(onReturn, ret)
}
//export CreateSoundMessage
func CreateSoundMessage(onReturn unsafe.Pointer, operationID *C.char, soundPath *C.char, duration int64) {
	ret := openIM.CreateSoundMessage(C.GoString(operationID), C.GoString(soundPath), duration)
	FunWithString(onReturn, ret)
}
//export CreateVideoMessageByURL
func CreateVideoMessageByURL(onReturn unsafe.Pointer, operationID *C.char, videoBaseInfo *C.char) {
	ret := openIM.CreateVideoMessageByURL(C.GoString(operationID), C.GoString(videoBaseInfo))
	FunWithString(onReturn, ret)
}
//export CreateVideoMessage
func CreateVideoMessage(onReturn unsafe.Pointer, operationID *C.char, videoPath *C.char, videoType *C.char, duration int64, snapshotPath *C.char) {
	ret := openIM.CreateVideoMessage(C.GoString(operationID), C.GoString(videoPath), C.GoString(videoType), duration, C.GoString(snapshotPath))
	FunWithString(onReturn, ret)
}
//export CreateFileMessageByURL
func CreateFileMessageByURL(onReturn unsafe.Pointer, operationID *C.char, fileBaseInfo *C.char) {
	ret := openIM.CreateFileMessageByURL(C.GoString(operationID), C.GoString(fileBaseInfo))
	FunWithString(onReturn, ret)
}
//export CreateFileMessage
func CreateFileMessage(onReturn unsafe.Pointer, operationID *C.char, filePath *C.char, fileName *C.char) {
	ret := openIM.CreateFileMessage(C.GoString(operationID), C.GoString(filePath), C.GoString(fileName))
	FunWithString(onReturn, ret)
}
//export CreateMergerMessage
func CreateMergerMessage(onReturn unsafe.Pointer, operationID *C.char, messageList, title, summaryList *C.char) {
	ret := openIM.CreateMergerMessage(C.GoString(operationID), C.GoString(messageList), C.GoString(title), C.GoString(summaryList))
	FunWithString(onReturn, ret)
}
//export CreateForwardMessage
func CreateForwardMessage(onReturn unsafe.Pointer, operationID *C.char, m *C.char) {
	ret := openIM.CreateForwardMessage(C.GoString(operationID), C.GoString(m))
	FunWithString(onReturn, ret)
}
//export SendMessage
func SendMessage(onError unsafe.Pointer, onSuccess unsafe.Pointer, onProgress unsafe.Pointer, operationID, message, recvID, groupID, offlinePushInfo *C.char) {
	openIM.SendMessage(&SendMsgCallBack{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
		onProgress: func(i int) {
			FunWithInt(onProgress, i)
		},
	}, C.GoString(operationID), C.GoString(message), C.GoString(recvID), C.GoString(groupID), C.GoString(offlinePushInfo))
}
//export SendMessageNotOss
func SendMessageNotOss(onError unsafe.Pointer, onSuccess unsafe.Pointer, onProgress unsafe.Pointer, operationID *C.char, message, recvID, groupID *C.char, offlinePushInfo *C.char) {
	openIM.SendMessageNotOss(&SendMsgCallBack{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
		onProgress: func(i int) {
			FunWithInt(onProgress, i)
		},
	}, C.GoString(operationID), C.GoString(message), C.GoString(recvID), C.GoString(groupID), C.GoString(offlinePushInfo))
}
//export GetHistoryMessageList
func GetHistoryMessageList(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, getMessageOptions *C.char) {
	openIM.GetHistoryMessageList(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(getMessageOptions))
}
//export RevokeMessage
func RevokeMessage(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, message *C.char) {
	openIM.RevokeMessage(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(message))
}
//export TypingStatusUpdate
func TypingStatusUpdate(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, recvID, msgTip *C.char) {
	openIM.TypingStatusUpdate(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(recvID), C.GoString(msgTip))
}
//export MarkC2CMessageAsRead
func MarkC2CMessageAsRead(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, userID *C.char, msgIDList *C.char) {
	openIM.MarkC2CMessageAsRead(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(userID), C.GoString(msgIDList))
}
//export MarkGroupMessageHasRead
func MarkGroupMessageHasRead(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupID *C.char) {
	openIM.MarkGroupMessageHasRead(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupID))
}
//export DeleteMessageFromLocalStorage
func DeleteMessageFromLocalStorage(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, message *C.char) {
	openIM.DeleteMessageFromLocalStorage(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(message))
}
//export func ClearC2CHistoryMessage(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, userID *C.char) {

func ClearC2CHistoryMessage(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, userID *C.char) {
	openIM.ClearC2CHistoryMessage(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(userID))
}
//export ClearGroupHistoryMessage
func ClearGroupHistoryMessage(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, groupID *C.char) {
	openIM.ClearGroupHistoryMessage(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(groupID))
}
//export InsertSingleMessageToLocalStorage
func InsertSingleMessageToLocalStorage(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, message, recvID, sendID *C.char) {
	openIM.InsertSingleMessageToLocalStorage(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(message), C.GoString(recvID), C.GoString(sendID))
}
//export InsertGroupMessageToLocalStorage
func InsertGroupMessageToLocalStorage(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, message, groupID, sendID *C.char) {
	openIM.InsertGroupMessageToLocalStorage(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(message), C.GoString(groupID), C.GoString(sendID))
}
//export SearchLocalMessages
func SearchLocalMessages(onError unsafe.Pointer, onSuccess unsafe.Pointer, operationID *C.char, searchParam *C.char) {
	openIM.SearchLocalMessages(&Base{
		onError: func(i int, s string) {
			FunWithIntString(onError, i, s)
		},
		onSuccess: func(s string) {
			FunWithString(onSuccess, s)
		},
	}, C.GoString(operationID), C.GoString(searchParam))
}

//func FindMessages(callback common.Base, operationID *C.char, messageIDList *C.char) {
//	userForSDK.Conversation().FindMessages(callback, messageIDList)
//}

func InitOnce(config *sdk_struct.IMConfig) bool {
	return openIM.InitOnce(config)
}

func CheckToken(userID, token *C.char) error {
	return openIM.CheckToken(C.GoString(userID), C.GoString(token))
}

func CheckResourceLoad(uSDK *login.LoginMgr) error {
	return openIM.CheckResourceLoad(uSDK)
}
//export GetConversationIDBySessionType
func GetConversationIDBySessionType(onReturn unsafe.Pointer, sourceID *C.char, sessionType int) {
	ret := openIM.GetConversationIDBySessionType(C.GoString(sourceID), sessionType)
	FunWithString(onReturn, ret)
}
