package main

func init() {
	DB.AutoMigrate(
		&User{},
		&Token{},
	)
}
