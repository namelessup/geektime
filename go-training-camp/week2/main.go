package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

// 问题：我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
// 个人见解：最好wrap住抛给上层。
// 第一：因为dao层基本都是操作数据库或者NoSQL之类的存储中间件的实现层，比如真正的sql语句只有在dao层面才会生成，而sql.ErrNoRows这个错误没有任何业务信息，所以可以将sql语句和非敏感信息包装给上层业务，用于开发排查问题。
// 第二：因为dao层不是业务层面，不能直接判断 sql.ErrNoRows 这是个正常的行为还是确实是数据缺失，根据错误只能处理或者往上抛的原则，所以需要往上抛错误信息。
// 第三：为了保持kit库简洁、高效、清晰的原则，如果dao层是一个封装好的标准库则不建议往上抛，交由调用次dao层方法的调用者Wrap。

// Dao 模拟dao接口，伪代码
type Dao interface {
	Get(target interface{}, sql string, values ...interface{}) error
}

// User 用户信息
type User struct {
	Id   int32
	Name string
}

// GetUserById 根据id获取用户信息
func GetUserById(dao Dao, id int32) (user User, err error) {
	sql := "SELECT id, name FROM t_user WHERE id = ?"
	err = dao.Get(&user, sql, id)
	if err != nil {
		err = fmt.Errorf("%w, sql=%s, id=%v", err, sql, id)
	}
	return
}

func main() {
	// 可以依赖注入进来
	var dao Dao

	user, err := doSomething(dao, 1)
	if err != nil {
		log.Fatalf("doSomething failed. err=%+v", err)
	}

	log.Println(user)
	os.Exit(0)
}

func doSomething(dao Dao, id int32) (*User, error) {
	user, err := GetUserById(dao, 1)
	if err != nil {
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		// 做一些插入或者降级操作
	}

	return &user, nil
}
