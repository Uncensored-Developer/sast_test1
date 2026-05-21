package main

import "database/sql"

func userByID(db *sql.DB, id string) {
	db.Query("SELECT * FROM users WHERE id = " + id)
}

const apiKey = "ghp_TEST_SECRET_VALUE_1234567890"
