package db

import (
	"fmt"
	"github.com/jackc/pgx"
	"sync"
	"time"
	"zhulei.com/stone/global"
)

var once sync.Once
var pool *pgx.ConnPool

func initStone() {
	dbConfig := global.Configs.Db

	conf := new(pgx.ConnPoolConfig)
	conf.Host = dbConfig.Host
	conf.Port = dbConfig.Port
	conf.Database = dbConfig.Database
	conf.User = dbConfig.User
	conf.Password = dbConfig.Password
	conf.MaxConnections = 256
	conf.AfterConnect = afterConnect
	conf.AcquireTimeout = time.Second * 5
	connPool, err := pgx.NewConnPool(*conf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	pool = connPool
}

func Pool() *pgx.ConnPool {
	once.Do(initStone)
	return pool
}

const StmtUserFindByMobile = "StmtUserFindByMobile"
const StmtUserInsert = "StmtUserInsert"

var stmts = map[string]string{
	StmtUserFindByMobile: `
	SELECT COUNT(user_id)
	FROM users
	WHERE mobile=$1;`,
	StmtUserInsert: `
	INSERT INTO users
	(name, mobile, password)
	VALUES ($1, $2, $3);`,
}

func afterConnect(conn *pgx.Conn) error {
	for name, sql := range stmts {
		if _, err := conn.Prepare(name, sql); err != nil {
			fmt.Println(name, sql, err.Error())
			return err
		}
	}
	return nil
}
