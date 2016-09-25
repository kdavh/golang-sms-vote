package person

type Model struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	// CreatedAt int32  `db:"created_at"`
	Email string
}
