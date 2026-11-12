package main

import (
	"errors"
	"fmt"
	"os"
	"bufio"
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

}