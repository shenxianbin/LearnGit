package main

import (
	"Gameserver/logic"
	"fmt"
)

//魔物测试
func main() {
	var userMonsters = logic.UserMonsters{}
	userMonsters.Demo()
	user := &logic.Role{}
	user.SetId(9527)
	userMonsters.Init(user)

	//userMonsters.DemoData(user)

	// userMonsters.GetData(9527, 10022)
	// userMonsters.GetData(9527, 10011)

	fmt.Println(userMonsters.Monsters)
}
