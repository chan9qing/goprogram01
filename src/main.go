package main

import(
	"fmt"
)

func main() {
	var b Book
	err0 := b.Load()
	if err0 == nil {
		fmt.Println("通讯录载入成功！")
	}

	
	for {
		uiMenu()
		chose := uiReadLine("请输入你的选择：")
	
		switch(chose){
		case "1":
			b.Add(uiReadLine("姓名："), uiReadLine("电话号码："))
		case "2":
			c,err2 := b.Find(uiReadLine("输入要寻找的姓名："))
			if err2 != nil{
				fmt.Println(err2)
			}
			if err2 == nil {
				fmt.Printf("%s\n%s", c.Name, c.Phone)
			}
		case "3":
			err3 := b.Remove(uiReadLine("输入要删除的姓名"))
			if err3 != nil{
				fmt.Println(err3)
			}
			if err3 == nil{
				fmt.Println("已删除成功")
			}
		case "4":
			b.uiShowAll()
		case "5":
			err5 := b.Rename(uiReadLine("输入要修改的姓名："),uiReadLine("新姓名："), uiReadLine("新电话："))
			if err5 != nil {
				fmt.Println(err5)
			}
			if err5 == nil {
				fmt.Println("修改成功！")
			}
		case "0":
			b.Save()
			fmt.Println("保存成功，再见！")
			return
		}
	}
}