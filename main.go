package main

import (
	"encoding/json"
	"fmt"
	gkb "github.com/2mf8/QQ-Markdown-Keyboard-Generator/keyboard"
)

func main() {
	/* kb := gkb.Builder().
		TextButton("测试", "已测试", "成功", false, true).
		UrlButton("爱魔方吧", "一仝", "https://2mf8.cn", false, true).
		SetRow().
		TextButton("测试", "已测试", "成功", false, true).
		SetRow().
		UrlButton("爱魔方吧", "一仝", "https://2mf8.cn", false, true).
		SetRow() */
	kb := gkb.Builder()
	for i := 1; i < 25; i++{
		is := fmt.Sprintf("%v", i)
		kb.Button(is, is, is, i, int(gkb.ActionTypeCallback), int(gkb.PermissionTypAll), false, true, false)
		if i % 5 == 0 {
			kb.SetRow()
		}
		if i == 24 {
			kb.SetRow()
		}
	}
	b, _ := json.Marshal(kb)
	fmt.Println(string(b))
}
