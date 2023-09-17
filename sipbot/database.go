package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var sip_domain string

type SIPAccount struct {
	Username string `json:"username"`
	Domain   string `json:"domain"`
	Password string `json:"password"`
}

func initDB(dsn string, domain string) error {
	sip_domain = domain
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func getAccount(addr string) (*SIPAccount, error) {
	acc := &SIPAccount{Username: addr2user(addr)}

	// check if account exist
	row := db.QueryRow("SELECT domain, password FROM accounts WHERE username = ?", acc.Username)
	if err := row.Scan(&acc.Domain, &acc.Password); err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	} else {
		return acc, nil
	}

	// new account
	acc.Domain = sip_domain
	acc.Password = genPassword()

	stmnt := "INSERT INTO accounts (username, domain, password, algorithm) VALUES (?, ?, ?, ?)"
	_, err := db.Exec(stmnt, acc.Username, acc.Domain, acc.Password, "CLRTXT")
	if err != nil {
		return nil, err
	}

	return acc, nil
}
