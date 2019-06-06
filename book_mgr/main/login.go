package main

import (
	"fmt"
	"../model"
	"strings"
)

func inputUser() (username, passwd string) {
	fmt.Printf("username:")
	fmt.Scan(&username)
	fmt.Printf("passwd:")
	fmt.Scan(&passwd)

	username = strings.TrimSpace(username)
	passwd = strings.TrimSpace(passwd)
	return
}

func adminLogin() (user *model.User, err error) {

	username, passwd := inputUser()
	user, err = mgr.AdminLogin(username, passwd)
	if err != nil {
		return
	}

	return
}

func userLogin() (user *model.User, err error) {

	username, passwd := inputUser()
	user, err = mgr.UserLogin(username, passwd)
	if err != nil {
		return
	}

	return
}
