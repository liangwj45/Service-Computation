package entity

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var (
	user   string
	users  map[string]User
	userDB string = "user.json"
)

func init() {
	users = map[string]User{}
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	ReadUserFile()
}

func ReadUserFile() {
	file, err := os.Open(userDB)
	if err != nil {
		log.Println("open file state failed:", err)
		return
	}
	state, _ := file.Stat()
	if state.Size() == 0 {
		log.Println("read file state failed:", err)
		return
	}
	buffer := make([]byte, state.Size())
	_, err = file.Read(buffer)
	if err != nil {
		log.Println("read file failed:", err)
		return
	}
	buffer = []byte(os.ExpandEnv(string(buffer)))
	err = json.Unmarshal(buffer, &users)
	if err != nil {
		log.Fatalln("json unmarshal failed:", err)
	}
}

func WriteUserFile() {
	userRec, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	f, _ := os.Create(userDB)
	defer f.Close()
	f.WriteString(string(userRec))
}

func AddUser(username string, password string, email string, telephone string) {
	user := User{
		Username:  username,
		Password:  password,
		Email:     email,
		Telephone: telephone,
	}
	users[user.Username] = user
}

func ExistUserName(username string) bool {
	for i := range users {
		if users[i].Username == username {
			return true
		}
	}
	return false
}

func CheckPassword(username string, password string) bool {
	for i := range users {
		if username == users[i].Username {
			return password == users[i].Password
		}
	}
	return false
}

func GetAllUser() map[string]User {
	return users
}
