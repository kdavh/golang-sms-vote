package main

import (
	"database/sql"
)

// Up is executed when this migration is applied
func Up_20160922212245(txn *sql.Tx) {
	txn.Exec(`
        CREATE TABLE person (
            first_name text,
            last_name text,
            email text,
            id SERIAL PRIMARY KEY
        );

        CREATE TABLE place (
            country text,
            city text NULL,
            telcode integer,
            id SERIAL PRIMARY KEY
        )
    `)
	txn.Exec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	txn.Exec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "John", "Doe", "johndoeDNE@gmail.net")
	txn.Exec("INSERT INTO place (country, city, telcode) VALUES ($1, $2, $3)", "United States", "New York", "1")
	txn.Exec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Hong Kong", "852")
	txn.Exec("INSERT INTO place (country, telcode) VALUES ($1, $2)", "Singapore", "65")
}

// Down is executed when this migration is rolled back
func Down_20160922212245(txn *sql.Tx) {
	txn.Exec(`
        DROP TABLE person;
        DROP TABLE place;
    `)
}
