package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	UserName string `json:"username"`
	NickName string `json:"nickname"`
	Age      int `json:"age"`
	Birthday string `json:"birthday"`
	Sex      string `json:"sex"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func testMarshal() (ret string, err error) {
	user := &User{
		UserName: "test",
		NickName: "t",
		Age:      18,
		Birthday: "2019/03/03",
		Sex:      "man",
		Email:    "123456@qq.com",
		Phone:    "123456",
	}
  
	jsonData, err := json.Marshal(user)
	if err != nil {
		err = fmt.Errorf("json Marshal error, err ", err)
		return
	}

	ret = string(jsonData)
	fmt.Printf("jsonStr: %s\n", ret)
	// {"UserName":"test","NickName":"t","Age":18,"Birthday":"2019/03/03","Sex":"man","Email":"123456@qq.com","Phone":"123456"}
	return
}

func testUnMarshal() {
	jsonStr, err := testMarshal()
	if err != nil {
		fmt.Println("get Jsonstr error, err: ", err)
		return
	}
	var user User
	err = json.Unmarshal([]byte(jsonStr), &user)
	if err != nil {
		fmt.Println("testUnMarshal Jsonstr error, err: ", err)
		return
	}
	fmt.Println(user)
}

func main()  {
	testMarshal()
	// testUnMarshal()
}
