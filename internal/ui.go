package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var stdin = bufio.NewReader(os.Stdin)

func uiMenu() error {
	fmt.Println("=============简易通讯录============")
	fmt.Println("1.添加联系人")

	fmt.Println("2.查找联系人")

	fmt.Println("3.删除联系人")

	fmt.Println("4.展示所有联系人")

	fmt.Println("5.修改联系人信息")

	fmt.Println("0.保存并退出")
	fmt.Printf("\n")
	return nil 
}

func (b *Book) uiShowAll()  {
	var count int = 0
	for _,c := range b.Contacts{
		fmt.Printf("%s %s\n", c.Name,c.Phone)
		count++
	}
	fmt.Printf("共%d个联系人\n",count)
}

func uiReadLine(prompt string) string {
	fmt.Print(prompt)
	line,err := stdin.ReadString('\n')
	if(err != nil){
		return "出错啦"
	}
	return strings.TrimSpace(line)
}
