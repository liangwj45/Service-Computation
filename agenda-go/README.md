# agenda-go

## 实验环境

- [Go Online](www.go-online.org.cn:8080)
- Go1.12

## 项目结构

- `agenda.go` 文件为 `main` 包，提供程序入口。
- `cmd` 包包含已实现的命令行指令。
  - `listUser.go` ：查询所有用户子指令的实现。
  - `login.go` ：账号登录子指令的实现。
  - `register.go` ：账号注册子指令的实现。
  - `root.go` ：agenda 命令行主要指令的实现。
- `entity` 包包含账号实体结构和存储方式的实现。
  - `storage.go` ：用户信息的持久化实现。
  - `user.go` ：账号实体结构。
- `agenda` 为程序可执行文件。
- `user.json` 为用户信息的存储。

## 程序源代码

### entity

#### user.go

```go
package entity

type User struct {
	Username  string
	Password  string
	Email     string
	Telephone string
}
```

#### storage.go

```go
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
```

### cmd

#### root.go

```go
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "agenda",
	Short: "A simple agenda",
	Long:  "Agenda is a application which help a user manage daily meeting schedule.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```

#### register.go

```go
package cmd

import (
	"fmt"
	"log"

	"github.com/liangwj45/Service-Computing/agenda-go/entity"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register -uUsername -pPassword",
	Short: "Register a account.",
	Long:  "Register a account with the username, password, email and telephone.",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		email, _ := cmd.Flags().GetString("email")
		telephone, _ := cmd.Flags().GetString("telephone")
		if !entity.ExistUserName(username) {
			entity.AddUser(username, password, email, telephone)
			log.Println("register succesful:", username, password)
			fmt.Println("----------------------------------------")
			entity.WriteUserFile()
		} else {
			log.Println("register failed:", username)
			fmt.Println("----------------------------------------")
			fmt.Println("username has been used")
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
	registerCmd.Flags().StringP("username", "u", "", "")
	registerCmd.Flags().StringP("password", "p", "", "")
	registerCmd.Flags().StringP("email", "e", "", "")
	registerCmd.Flags().StringP("telephone", "t", "", "")
	registerCmd.MarkFlagRequired("username")
	registerCmd.MarkFlagRequired("password")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
```

login.go

```go
package cmd

import (
	"fmt"
	"log"

	"github.com/liangwj45/Service-Computing/agenda-go/entity"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login -uUsername -pPassword",
	Short: "Log in your account.",
	Long:  "Log in your account with the username and password",
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		if entity.CheckPassword(username, password) {
			log.Println("login succesful:", username)
			fmt.Println("----------------------------------------")
			fmt.Println("welcome")
		} else {
			log.Println("login failed:", username, password)
			fmt.Println("----------------------------------------")
			fmt.Println("username or password incorrect")
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "u", "", "")
	loginCmd.Flags().StringP("password", "p", "", "")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
```

#### listUser.go

```go
package cmd

import (
	"fmt"
	"log"

	"github.com/liangwj45/Service-Computing/agenda-go/entity"
	"github.com/spf13/cobra"
)

var listUserCmd = &cobra.Command{
	Use:   "listUser",
	Short: "List all registered users.",
	Long:  "List the information of all registered users.",
	Run: func(cmd *cobra.Command, args []string) {
		users := entity.GetAllUser()
		log.Println("query all users")
		fmt.Printf("Query all users\n")
		fmt.Printf("Username\t Email\tTelephone\n")
		for i := range users {
			fmt.Printf("%s\t %s\t%s\n", users[i].Username, users[i].Email, users[i].Telephone)
		}
	},
}

func init() {
	rootCmd.AddCommand(listUserCmd)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
```

### agenda.go

```go
package main

import (
	"github.com/liangwj45/Service-Computing/agenda-go/cmd"
)

func main() {
	cmd.Execute()
}
```

## 编译运行结果

### 编译

```bash
go build agenda.go
```

### 运行

查看使用方法。

```bash
./agenda --help
```

![1](.\img\1.png)

尝试注册账号。

```bash
./agenda register
```

![2](.\img\2.png)

注册账号。

```bash
./agenda register -uliang -p123456
```

![3](.\img\3.png)

显示用户名已经被使用，那么用 `listUser` 指令看一下有哪些用户。

```bash
./agenda listUser
```

![4](.\img\4.png)

再注册一个账号。

```bash
./agenda register -uli -p123456 -eli@gmail.com -t13313131414
```

![6](.\img\6.png)

登录账号。

![7](.\img\7.png)

显示登录成功。

