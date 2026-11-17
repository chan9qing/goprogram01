package main

import "testing"

	func TestAddAndFind(t *testing.T) {
	var b Book

	if err := b.Add("张三", "13800138000"); err != nil {
		t.Fatalf("添加失败：%v", err)
	}
	if i,_ := b.Find("张三"); i == nil {
		t.Fatal("添加之后应该能查到")
	}
	}

	func TestAddDuplicate(t *testing.T) {
	var b Book
	b.Add("张三", "13800138000")

	if err := b.Add("张三", "13900139000"); err == nil {
		t.Fatal("重名应该报错，但没有")
	}
	}

	func TestRemove(t *testing.T) {
	var b Book
	b.Add("张三", "13800138000")

	if err := b.Remove("张三"); err != nil {
		t.Fatalf("删除失败：%v", err)
	}
	if i,_ := b.Find("张三"); i != nil {
		t.Fatal("删掉之后不应该还能查到")
	}
	if err := b.Remove("王五"); err == nil {
		t.Fatal("删不存在的人应该报错")
	}
	}