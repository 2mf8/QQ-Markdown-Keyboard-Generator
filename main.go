package main

import (
	"encoding/json"
	"fmt"

	gkb "github.com/2mf8/QQ-Markdown-Keyboard-Generator/keyboard"
	md "github.com/2mf8/QQ-Markdown-Keyboard-Generator/markdown"
)

func main() {
	// 按钮模板生成
	kb := gkb.Builder(101981675)
	for i := 1; i < 26; i++ {
		is := fmt.Sprintf("%v", i-1)
		kb.Button(is, is, is, i, int(gkb.ActionTypeCallback), int(gkb.PermissionTypAll), false, true, false)
		if i%5 == 0 {
			kb.SetRow()
		}
	}
	// 模板markdown消息生成
	t := md.TempBuilder("3434").TempParamAdd("url", "https://2mf8.cn").TempParamAdd("好家伙", "不错")
	b, _ := json.Marshal(kb)
	tv, _ := json.Marshal(t)
	// markdown模板生成
	_md := md.TempGenerator().
		H1("h1").
		NewLine().
		Image("desc", "url", 786, 343).
		Url("urltext", "urllink").
		Bold("boldtext").
		DividerLine().
		Code("code").
		Italic("italic")
		// 原生markdown消息生成
	m := md.Builder().CmdEnter("测试").CmdInput("测试", "显示", "true")
	// 生成结果打印到控制台
	fmt.Println(string(b), "\n\n\n", string(tv), "\n\n\n", m.Str, "\n\n\n", _md.Str)
}
