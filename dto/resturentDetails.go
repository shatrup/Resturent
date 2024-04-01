package dto

// resturent_id
// db:"customer_id"`
type ResturentResponse struct {
	ResturentId    string `db:"resturent_id"`
	ResturentName  string `db:"resturent_name"`
	ResturentImage string `db:"resturent_image"`
	OpenningTime   string `db:"openning_time"`
	ClosingTime    string `db:"closeing_time"`
	Desription     string `db:"description"`
	ContactEmail   string `db:"contact_email"`
	ContactNumber  string `db:"contact_number"`
}
