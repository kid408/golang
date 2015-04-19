package main

import (
	"container/list"
	"fmt"
)

type RoleInfo struct {
	mRoleID    uint32
	mRoleName  string
	mRoleValue uint32
}

func main() {
	mRoleIDMap := make(map[uint32]RoleInfo)
	mRoleNameMap := make(map[string]RoleInfo)
	var mRoleId uint32
	var mRoleInfo RoleInfo

	mRoleList := list.New()
	RoleRankSliceList := make([]*list.List, 0)
	RoleRankSliceList = append(RoleRankSliceList, mRoleList)

	for {

		mRoleInfo.mRoleID = mRoleId
		//fmt.Println("请输入角色ID:")
		//fmt.Scanln(&mRoleInfo.mRoleID)
		fmt.Println("请输入角色名字:")
		fmt.Scanln(&mRoleInfo.mRoleName)
		fmt.Println("请输入角色积分:")
		fmt.Scanln(&mRoleInfo.mRoleValue)

		if _, ok := mRoleNameMap[mRoleInfo.mRoleName]; ok == true {
			fmt.Println("角色名重复,请重新输入")
			continue
		}
		mRoleId = mRoleId + 1
		fmt.Println("角色ID:", mRoleInfo.mRoleID, ",角色名字:", mRoleInfo.mRoleName, ",角色积分:", mRoleInfo.mRoleValue)
		fmt.Println("-----------------------------------")

		for i := 0; i < len(RoleRankSliceList); i++ {

			if RoleRankSliceList[i].Len() == 0 {
				RoleRankSliceList[i].PushBack(mRoleInfo)
				mRoleIDMap[mRoleInfo.mRoleID] = mRoleInfo
				mRoleNameMap[mRoleInfo.mRoleName] = mRoleInfo
				continue
			}

			if RoleRankSliceList[i].Back().Value.(RoleInfo).mRoleValue > mRoleInfo.mRoleValue && RoleRankSliceList[i].Len() >= 3 {
				if i+1 == len(RoleRankSliceList) {
					RoleRankSliceList = append(RoleRankSliceList, list.New())
					RoleRankSliceList[i+1].PushBack(mRoleInfo)
					mRoleIDMap[mRoleInfo.mRoleID] = mRoleInfo
					mRoleNameMap[mRoleInfo.mRoleName] = mRoleInfo
					break
				}
				continue
			} else {
				if RoleRankSliceList[i].Front().Value.(RoleInfo).mRoleValue < mRoleInfo.mRoleValue {
					RoleRankSliceList[i].InsertBefore(mRoleInfo, RoleRankSliceList[i].Front())
					mRoleIDMap[mRoleInfo.mRoleID] = mRoleInfo
					mRoleNameMap[mRoleInfo.mRoleName] = mRoleInfo

				} else {
					for e := RoleRankSliceList[i].Back(); e != nil; e = e.Prev() {
						if mRoleInfo.mRoleValue < e.Value.(RoleInfo).mRoleValue {
							RoleRankSliceList[i].InsertAfter(mRoleInfo, e)
							mRoleIDMap[mRoleInfo.mRoleID] = mRoleInfo
							mRoleNameMap[mRoleInfo.mRoleName] = mRoleInfo
							break
						}
					}
				}
				break
			}
		}

		for i := 0; i < len(RoleRankSliceList); i++ {
			for e := RoleRankSliceList[i].Front(); e != nil; e = e.Next() {
				kv := e.Value.(RoleInfo)
				fmt.Println("Group:", i, "mRoleID:", kv.mRoleID, "mRoleName:", kv.mRoleName, "mRoleValue:", kv.mRoleValue)
			}
		}
		fmt.Println("-----------------------------------")
	}
}
