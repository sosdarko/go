package main

import (
	"database/sql"
	"fmt"

	ora "gopkg.in/rana/ora.v4"
)

func test2() {
	db, err := sql.Open("ora", "npfe/pwd4npfe@b22w3x64mm1db1.finbet.com:1521/BGMM1DEV")
	defer db.Close()
	if err != nil {
		panic("shit")
	}
	fmt.Println("connected")
	rows, err := db.Query("SELECT OBJECT_NAME, OBJECT_TYPE FROM user_objects")
	if err != nil {
		panic("shit2")
	}
	cols, _ := rows.Columns()
	fmt.Printf("%v\n", cols)
	i := 0
	for rows.Next() {
		i++
		var OBJECT_NAME string
		var OBJECT_TYPE string
		err = rows.Scan(&OBJECT_NAME, &OBJECT_TYPE)
		fmt.Println(i, OBJECT_NAME, OBJECT_TYPE)
	}
	err = rows.Err() // get any error encountered during iteration
	defer rows.Close()
}

func aMTestDB(base uint64) uint64 {
	// example usage of the ora package driver
	// connect to a server and open a session
	env, err := ora.OpenEnv()
	defer env.Close()
	if err != nil {
		panic(err)
	}
	srvCfg := ora.SrvCfg{Dblink: "BGMM1DEV"}
	srv, err := env.OpenSrv(srvCfg)
	defer srv.Close()
	if err != nil {
		panic(err)
	}
	sesCfg := ora.SesCfg{
		Username: "foconvertor",
		Password: "foconvertor",
	}
	ses, err := srv.OpenSes(sesCfg)
	defer ses.Close()
	if err != nil {
		panic(err)
	}

	// call stored procedure
	// pass *Rset to Exe to receive the results of a sys_refcursor
	stmtProcCall, err := ses.Prep("CALL AMTest(:1,:2)")

	defer stmtProcCall.Close()

	if err != nil {
		panic(err)
	}

	var output uint64

	fmt.Println("executing")
	_, err = stmtProcCall.Exe(&base, &output)
	fmt.Printf("input = %d output=%d\n", base, output)

	if err != nil {
		panic(err)
	}
	return output
}

func main() {
	fmt.Println("start")
	/*
		r := aMTestDB(10)
		fmt.Printf("%d\n", r)
	*/
	test2()
}
