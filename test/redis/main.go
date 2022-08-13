package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/labstack/gommon/log"
)

const (
	// Redis commands.
	cmdDel     = "DEL"
	cmdHDel    = "HDEL"
	cmdHGetAll = "HGETALL"
	cmdHGet    = "HGET"
	cmdHMGet   = "HMGET"
	cmdHSet    = "HSET"
	cmdSet     = "SET"
	cmdGet     = "GET"
	cmdSetNX   = "SETNX"
	cmdHLen    = "HLEN"
	cmdLPush   = "LPUSH"
	cmdLIndex  = "LINDEX"
	cmdLTrim   = "LTRIM"

	// Redis command options.
	optMillisecond = "PX"
	optSecond      = "EX"
	optNotExist    = "NX"

	// Redis OK status.
	kvStatusOK = "OK"

	// Redis Keys
	strKey1 = "test.str.first"
	strKey2 = "test.str.second"
)

func main() {
	var intVal int
	var strVal string

	kvp := redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.DialURL("redis://localhost:6379/0")
		},
	}
	conn := kvp.Get()

	defer func() {
		err := kvp.Close()
		if err != nil {
			log.Warn(err)
		}
	}()

	repl, err := conn.Do(cmdSet, strKey1, "here is first text value")
	if err != nil {
		log.Warn(err)
	}
	fmt.Println(repl)

	repl, err = conn.Do(cmdSet, strKey2, 123)
	if err != nil {
		log.Warn(err)
	}
	fmt.Println(repl)

	repl, err = conn.Do(cmdGet, strKey1)
	if err != nil {
		log.Warn(err)
	}
	fmt.Println(string(repl.([]byte)))
	strVal, err = redis.String(conn.Do(cmdGet, strKey1))
	fmt.Println(strVal)

	intVal, err = redis.Int(conn.Do(cmdGet, strKey2))
	if err != nil {
		log.Warn(err)
	}
	fmt.Printf("type: %T, val: %v\n", intVal, intVal)

	reply, err := redis.Values(conn.Do("MGET", strKey1, strKey2))
	strVal = ""
	intVal = 0
	repl, err = redis.Scan(reply, &strVal, &intVal)
	if err != nil {
		log.Warn(err)
	}
	fmt.Printf("StringVal: %s, Intval: %v\n", strVal, intVal)
}
