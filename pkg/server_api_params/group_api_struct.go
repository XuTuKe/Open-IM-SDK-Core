package server_api_params

type CommResp struct {
	ErrCode int32  `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}
type CommDataResp struct {
	CommResp
	Data []map[string]interface{} `json:"data"`
}
type CommDataRespOne struct {
	CommResp
	Data map[string]interface{} `json:"data"`
}

type KickGroupMemberReq struct {
	GroupID          string   `json:"groupID" binding:"required"`
	KickedUserIDList []string `json:"kickedUserIDList" binding:"required"`
	Reason           string   `json:"reason"`
	OperationID      string   `json:"operationID" binding:"required"`
}
type KickGroupMemberResp struct {
	CommResp
	UserIDResultList []*UserIDResult `json:"data"`
}

type GetGroupMembersInfoReq struct {
	GroupID     string   `json:"groupID" binding:"required"`
	MemberList  []string `json:"memberList" binding:"required"`
	OperationID string   `json:"operationID" binding:"required"`
}
type GetGroupMembersInfoResp struct {
	CommResp
	MemberList []*GroupMemberFullInfo   `json:"-"`
	Data       []map[string]interface{} `json:"data"`
}

type InviteUserToGroupReq struct {
	GroupID           string   `json:"groupID" binding:"required"`
	InvitedUserIDList []string `json:"invitedUserIDList" binding:"required"`
	Reason            string   `json:"reason"`
	OperationID       string   `json:"operationID" binding:"required"`
}
type InviteUserToGroupResp struct {
	CommResp
	UserIDResultList []*UserIDResult `json:"data"`
}

type GetJoinedGroupListReq struct {
	OperationID string `json:"operationID" binding:"required"`
	FromUserID  string `json:"fromUserID" binding:"required"`
}
type GetJoinedGroupListResp struct {
	CommResp
	GroupInfoList []*GroupInfo
	Data          []map[string]interface{} `json:"data"`
}

type GetGroupMemberListReq struct {
	GroupID     string `json:"groupID"`
	Filter      int32  `json:"filter"`
	NextSeq     int32  `json:"nextSeq"`
	OperationID string `json:"operationID"`
}
type GetGroupMemberListResp struct {
	CommResp
	NextSeq    int32 `json:"nextSeq"`
	MemberList []*GroupMemberFullInfo
	Data       []map[string]interface{} `json:"data"`
}

type GetGroupAllMemberReq struct {
	GroupID     string `json:"groupID" binding:"required"`
	OperationID string `json:"operationID" binding:"required"`
}
type GetGroupAllMemberResp struct {
	CommResp
	MemberList []*GroupMemberFullInfo   `json:"-"`
	Data       []map[string]interface{} `json:"data"`
}

type CreateGroupReq struct {
	MemberList  []*GroupAddMemberInfo `json:"memberList"  binding:"required"`
	OwnerUserID string                `json:"ownerUserID" binding:"required"`
	GroupName   string                `json:"groupName"`
	GroupType   int32                 `json:"groupType"`
	OperationID string                `json:"operationID" binding:"required"`
}
type CreateGroupResp struct {
	CommResp
	GroupInfo GroupInfo
	Data      map[string]interface{} `json:"data"`
}

type GetGroupApplicationListReq struct {
	OperationID string `json:"operationID" binding:"required"`
	FromUserID  string `json:"fromUserID" binding:"required"` //作为管理员或群主收到的 进群申请
}
type GetGroupApplicationListResp struct {
	CommResp
	GroupRequestList []*GroupRequest
	Data             []map[string]interface{} `json:"data"`
}

type GetGroupInfoReq struct {
	GroupIDList []string `json:"groupIDList" binding:"required"`
	OperationID string   `json:"operationID" binding:"required"`
}
type GetGroupInfoResp struct {
	CommResp
	GroupInfoList []*GroupInfo
	Data          []map[string]interface{} `json:"data"`
}

type ApplicationGroupResponseReq struct {
	OperationID  string `json:"operationID" binding:"required"`
	GroupID      string `json:"groupID" binding:"required"`
	FromUserID   string `json:"fromUserID" binding:"required"` //application from FromUserID
	HandledMsg   string `json:"handledMsg"`
	HandleResult int32  `json:"handleResult" binding:"required,oneof=-1 1"`
}
type ApplicationGroupResponseResp struct {
	CommResp
}

type JoinGroupReq struct {
	GroupID     string `json:"groupID" binding:"required"`
	ReqMessage  string `json:"reqMessage"`
	OperationID string `json:"operationID" binding:"required"`
}
type JoinGroupResp struct {
	CommResp
}

type QuitGroupReq struct {
	GroupID     string `json:"groupID" binding:"required"`
	OperationID string `json:"operationID" binding:"required"`
}
type QuitGroupResp struct {
	CommResp
}

type SetGroupInfoReq struct {
	GroupID      string `json:"groupID" binding:"required"`
	GroupName    string `json:"groupName"`
	Notification string `json:"notification"`
	Introduction string `json:"introduction"`
	FaceURL      string `json:"faceURL"`
	Ex           string `json:"ex"`
	OperationID  string `json:"operationID" binding:"required"`
}

type SetGroupInfoResp struct {
	CommResp
}

type TransferGroupOwnerReq struct {
	GroupID        string `json:"groupID" binding:"required"`
	OldOwnerUserID string `json:"oldOwnerUserID" binding:"required"`
	NewOwnerUserID string `json:"newOwnerUserID" binding:"required"`
	OperationID    string `json:"operationID" binding:"required"`
}
type TransferGroupOwnerResp struct {
	CommResp
}