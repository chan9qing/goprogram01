package main

import (
	"errors"
	"slices"
)

type Contact struct{
	Name string
	Phone string
}

type Book struct{
	Contacts []*Contact
}

func (b *Book) Add (name string, phone string) error{
	for _,contact := range b.Contacts{
		if contact.Name == name{
			return errors.New("该联系人已添加")
		}
	}
	c := &Contact{
		Name: name,
		Phone: phone,
	}
	b.Contacts = append(b.Contacts, c)
	return nil
}

func (b *Book) Remove (name string) error {
	for i,contact := range b.Contacts{
		if contact.Name == name{
			b.Contacts = slices.Delete(b.Contacts, i, i+1)
			return nil
		}
	}
	return errors.New("该联系人不存在")
}

func (b *Book) Find (name string) (*Contact, error){
	for _,Contact := range b.Contacts{
		if Contact.Name == name{
			return Contact,nil
		}
	}
	return nil,errors.New("未查找到该联系人")
}

func (b *Book) Rename (name string, newname string, newphone string) error {
	for _,Contact := range b.Contacts{
		if Contact.Name == name{
			Contact.Name = newname
			Contact.Phone = newphone 
			return nil
		}
		
	}
	return errors.New("未找到该联系人")
}