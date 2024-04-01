package domain

type Resturents struct {
  ResturentId          string `json:"resturent_id"`
  ResturentName        string `json:"resturent_name"`
  ResturentImage        string `json:"resturent_image"`
  OpenningTime     string `json:"openning_time"`
  ClosingTime string `json:"closeing_time"`
  Desription        string `json:"description"`
  ContactEmail      string `json:"contact_email"`
  ContactNumber        string `json:"contact_number"`
}
