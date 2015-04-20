package main

import (
	"container/list"
	"fmt"
)

type RoleInfo struct {
	mRoleID    uint32
	mRoleName  string
	mRoleValue uint32
	mRanking   uint32
}

var mRoleIDMap map[uint32]*RoleInfo
var mRoleNameMap map[string]*RoleInfo
var mRoleList *list.List
var RoleRankSliceList []*list.List
var mRoleId uint32

//添加角色信息
func InsertRoleInfo() {

	var mTempRoleInfo RoleInfo

	mTempRoleInfo.mRoleID = mRoleId
	//fmt.Println("请输入角色ID:")
	//fmt.Scanln(&mRoleInfo.mRoleID)
	fmt.Println("请输入角色名字:")
	fmt.Scanln(&mTempRoleInfo.mRoleName)
	fmt.Println("请输入角色积分:")
	fmt.Scanln(&mTempRoleInfo.mRoleValue)

	if _, ok := mRoleNameMap[mTempRoleInfo.mRoleName]; ok == true {
		fmt.Println("角色名重复,请重新输入")
		return
	}

	mRoleInfo := new(RoleInfo)
	mRoleId = mRoleId + 1
	mRoleInfo.mRoleID = mRoleId
	mRoleInfo.mRoleName = mTempRoleInfo.mRoleName
	mRoleInfo.mRoleValue = mTempRoleInfo.mRoleValue

	fmt.Println("角色ID:", mRoleInfo.mRoleID, ",角色名字:", mRoleInfo.mRoleName, ",角色积分:", mRoleInfo.mRoleValue)
	fmt.Println("-----------------------------------")

	for i := 0; i < len(RoleRankSliceList); i++ {

		if RoleRankSliceList[i].Len() == 0 {
			RoleRankSliceList[i].PushBack(mRoleInfo)
			mRoleIDMap[mRoleInfo.mRoleID] = mRoleInfo
			mRoleNameMap[mRoleInfo.mRoleName] = mRoleInfo
			continue
		}
		if RoleRankSliceList[i].Back().Value.(*RoleInfo).mRoleValue > mRoleInfo.mRoleValue && RoleRankSliceList[i].Len() >= 3 {
			if i+1 == len(RoleRankSliceList) {
				RoleRankSliceList = append(RoleRankSliceList, list.New())
				RoleRankSliceList[i+1].PushBack(mRoleInfo)
				mRoleIDMap[mRoleInfo.mRoleID] = mRoleInfo
				mRoleNameMap[mRoleInfo.mRoleName] = mRoleInfo
				break
			}
			continue
		} else {
			if RoleRankSliceList[i].Front().Value.(*RoleInfo).mRoleValue < mRoleInfo.mRoleValue {
				RoleRankSliceList[i].InsertBefore(mRoleInfo, RoleRankSliceList[i].Front())
				mRoleIDMap[mRoleInfo.mRoleID] = mRoleInfo
				mRoleNameMap[mRoleInfo.mRoleName] = mRoleInfo

			} else {
				for e := RoleRankSliceList[i].Back(); e != nil; e = e.Prev() {
					if mRoleInfo.mRoleValue < e.Value.(*RoleInfo).mRoleValue {
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
}

//初始化容器，方便测试
func InitSliceRole() {
	//RoleInfo{"mRoleID": 1, "mRoleName": kid, "mRoleValue": 100}
	//temp_role := new{RoleInfo{mRoleID: 1, mRoleName: "kid", mRoleValue: 100}}
	temp_role := new(RoleInfo)
	temp_role.mRoleID = 1
	temp_role.mRoleName = "kid"
	temp_role.mRoleValue = 100
	RoleRankSliceList[0].PushBack(temp_role)
	mRoleIDMap[temp_role.mRoleID] = temp_role
	mRoleNameMap[temp_role.mRoleName] = temp_role
}

//更新角色信息
func UpdateRole(role RoleInfo) {

	if _, ok := mRoleIDMap[role.mRoleID]; ok == false {
		fmt.Println("角色名重复没有找到这个人,请重新输入")
		return
	}

	mRoleIDMap[role.mRoleID].mRoleID = role.mRoleID
	mRoleIDMap[role.mRoleID].mRoleName = role.mRoleName
	mRoleIDMap[role.mRoleID].mRoleValue = role.mRoleValue
}

//返回角色当前位置
func GetRoleRanking(mTempRoleId uint32) RoleInfo {
	var mTempRanking uint32
	var result bool
	var mTempRoleInfo RoleInfo
	if mRoleInfo, ok := mRoleIDMap[mTempRoleId]; ok {
		for i := 0; i < len(RoleRankSliceList); i++ {
			mTempRanking = uint32(RoleRankSliceList[i].Len()) + mTempRanking
			for e := RoleRankSliceList[i].Back(); e != nil; e = e.Prev() {
				if mRoleInfo.mRoleValue >= e.Value.(*RoleInfo).mRoleValue {
					mTempRanking = mTempRanking - 1
					if mRoleInfo.mRoleID == e.Value.(*RoleInfo).mRoleID {
						result = true
						mRoleInfo.mRanking = mTempRanking + 1
					}
				} else {
					break
				}
			}
			if result {
				return *mRoleInfo
			}
		}

	} else {
		fmt.Println("not find the key :", mTempRoleId)
	}
	return mTempRoleInfo
}

func DeleteRoleByID(mTempRoleId uint32) {
	var result bool
	if mRoleInfo, ok := mRoleIDMap[mTempRoleId]; ok {
		for i := 0; i < len(RoleRankSliceList); i++ {
			for e := RoleRankSliceList[i].Back(); e != nil; e = e.Prev() {
				if mRoleInfo.mRoleValue >= e.Value.(*RoleInfo).mRoleValue {
					if mRoleInfo.mRoleID == e.Value.(*RoleInfo).mRoleID {
						RoleRankSliceList[i].Remove(e)
						result = true
						break
					}
				} else {
					break
				}
			}
			if result {
				break
			}
		}

		delete(mRoleIDMap, mRoleInfo.mRoleID)
		delete(mRoleNameMap, mRoleInfo.mRoleName)
	} else {
		fmt.Println("not find the key :", mTempRoleId)
	}

}

func main() {
	mRoleIDMap = make(map[uint32]*RoleInfo)
	mRoleNameMap = make(map[string]*RoleInfo)

	mRoleList = list.New()
	RoleRankSliceList = make([]*list.List, 0)
	RoleRankSliceList = append(RoleRankSliceList, mRoleList)
	var num uint32
	for {
		num = num + 1

		if num < 7 {
			InsertRoleInfo()
		}
		if num >= 7 {
			var mTempRoleID uint32

			fmt.Println("请输要查询的角色ID:")
			fmt.Scanln(&mTempRoleID)
			v := GetRoleRanking(mTempRoleID)
			fmt.Println("mRoleID: ", v.mRoleID, ",mRoleName: ", v.mRoleName, ",mRoleValue: ", v.mRoleValue, ",mRanking: ", v.mRanking)
		}
		if num >= 100 {
			for _, v := range mRoleIDMap {
				fmt.Println("IDMap -- mRoleID: ", v.mRoleID, ",mRoleName: ", v.mRoleName, ",mRoleValue: ", v.mRoleValue)
			}
			fmt.Println("-----------------------------------")
			for _, v := range mRoleNameMap {
				fmt.Println("NameMap -- mRoleID: ", v.mRoleID, ",mRoleName: ", v.mRoleName, ",mRoleValue: ", v.mRoleValue)
			}
			fmt.Println("-----------------------------------")

			for i := 0; i < len(RoleRankSliceList); i++ {
				for e := RoleRankSliceList[i].Front(); e != nil; e = e.Next() {
					kv := e.Value.(*RoleInfo)
					fmt.Println("Group:", i, "mRoleID:", kv.mRoleID, "mRoleName:", kv.mRoleName, "mRoleValue:", kv.mRoleValue)
				}
			}
			fmt.Println("-----------------------------------")

			var mTempRoleID uint32

			fmt.Println("请输入删除的角色ID:")
			fmt.Scanln(&mTempRoleID)
			DeleteRoleByID(mTempRoleID)

		}

		for i := 0; i < len(RoleRankSliceList); i++ {
			for e := RoleRankSliceList[i].Front(); e != nil; e = e.Next() {
				kv := e.Value.(*RoleInfo)
				fmt.Println("Group:", i, "mRoleID:", kv.mRoleID, "mRoleName:", kv.mRoleName, "mRoleValue:", kv.mRoleValue)
			}
		}
		fmt.Println("-----------------------------------")
	}
}
