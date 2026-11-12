package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

const FilePath = "contact.txt"

func (b *Book) Save() error {
	fp,err := os.Create(FilePath)
	if err != nil {
		return errors.New("取文件指针出错")
	}
	defer fp.Close()
	for _,c := range b.Contacts{ 
		_,_ = fmt.Fprintf(fp, "%s %s\n", c.Name, c.Phone)
	}
	return nil
}

func (b *Book) Load() error {
	fp,err := os.Open(FilePath)
	if err != nil {
		return err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	parts := strings.Fields(scanner.Text())
	if len(parts) == 2 {
		c := &Contact{
			Name: parts[0],
			Phone: parts[1],
		}
		b.Contacts = append(b.Contacts, c)
	}
	if erraa := scanner.Err(); erraa != nil{
		return erraa
	}
	return nil
}