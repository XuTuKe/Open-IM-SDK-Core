package common

import (
	"open_im_sdk/pkg/db"
	api "open_im_sdk/pkg/server_api_params"
)

func MergeBlackFriendResult(base []*db.LocalBlack, add []*db.LocalFriend) []api.FullUserInfo {
	m := make(map[string]api.FullUserInfo, 0)
	for _, v := range base {
		node := api.FullUserInfo{}
		node.BlackInfo = v
		m[v.BlockUserID] = node
	}

	for _, v := range add {
		if t, ok := m[v.FriendUserID]; ok {
			t.FriendInfo = v
			m[v.FriendUserID] = t
		}
	}

	r := make([]api.FullUserInfo, 0)
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func MergeFriendBlackResult(base []*db.LocalFriend, add []*db.LocalBlack) []api.FullUserInfo {
	m := make(map[string]api.FullUserInfo, 0)
	for _, v := range base {
		node := api.FullUserInfo{}
		node.FriendInfo = v
		m[v.FriendUserID] = node
	}

	for _, v := range add {
		if t, ok := m[v.BlockUserID]; ok {

			t.BlackInfo = v
			m[v.BlockUserID] = t
		}
	}

	r := make([]api.FullUserInfo, 0)
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func MergeUserResult(publicList []*api.PublicUserInfo, friendList []*db.LocalFriend, blackList []*db.LocalBlack) []api.FullUserInfo {
	m := make(map[string]api.FullUserInfo, 0)
	for _, v := range publicList {
		node := api.FullUserInfo{}
		node.PublicInfo = v
		m[v.UserID] = node
	}

	for _, v := range friendList {
		if t, ok := m[v.FriendUserID]; ok {
			t.FriendInfo = v
			m[v.FriendUserID] = t
		} else {
			node := api.FullUserInfo{PublicInfo: &api.PublicUserInfo{}}
			node.FriendInfo = v
			node.PublicInfo.UserID = v.FriendUserID
			node.PublicInfo.FaceURL = v.FaceURL
			node.PublicInfo.Nickname = v.Nickname
			node.PublicInfo.Gender = v.Gender
			m[v.FriendUserID] = node
		}
	}

	for _, v := range blackList {
		if t, ok := m[v.BlockUserID]; ok {
			t.BlackInfo = v
			m[v.BlockUserID] = t
		} else {
			node := api.FullUserInfo{PublicInfo: &api.PublicUserInfo{}}
			node.BlackInfo = v
			node.PublicInfo.UserID = v.BlockUserID
			node.PublicInfo.FaceURL = v.FaceURL
			node.PublicInfo.Nickname = v.Nickname
			node.PublicInfo.Gender = v.Gender
			m[v.BlockUserID] = node
		}
	}

	r := make([]api.FullUserInfo, 0)
	for _, v := range m {
		r = append(r, v)
	}
	return r
}
