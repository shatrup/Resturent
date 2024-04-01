
package app

import (
  "fmt"
  "net/http"
  "encoding/json"
  "Resturent/config"
  "Resturent/dto"
  "Resturent/domain"
  "github.com/gorilla/mux"
)


func Start() {
  fmt.Println("Start Resturent Apps")
  router := mux.NewRouter()
	// define routes
	router.
		HandleFunc("/resturents", getAllResturent).
		Methods(http.MethodGet)

  router.
		HandleFunc("/resturents", insertResturentDetails).
		Methods(http.MethodPost)
    router.
  		HandleFunc("/resturents/{id}", updateResturentDetails).
  		Methods(http.MethodPut)

    router.
  		HandleFunc("/resturents/{id}", deleteResturentDetails).
  		Methods(http.MethodDelete)
    fmt.Println(http.ListenAndServe(":10000",router))
}

func getAllResturent(w http.ResponseWriter, r *http.Request) {
  var err error
	resturents := make([]dto.ResturentResponse, 0)
  client := config.DbConnection()
	getAllSql := "select * from resturentDetails"
	err = client.Select(&resturents, getAllSql)
	if err != nil {
		fmt.Println("Error while querying resturents table " + err.Error())
	}
  fmt.Println("All resturnet name is ", resturents)
}

func insertResturentDetails(w http.ResponseWriter, r *http.Request) {
  var err error
  resturent := domain.Resturents{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&resturent); err != nil {
    w.WriteHeader(http.StatusBadRequest)
   w.Write([]byte("internal server error"))
		return
	}
	defer r.Body.Close()
  client := config.DbConnection()
  sqlInsert := "INSERT INTO resturentDetails (resturent_id, resturent_name, resturent_image, openning_time, closeing_time,description,contact_email,contact_number) values (?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := client.Exec(sqlInsert, resturent.ResturentId, resturent.ResturentName, resturent.ResturentImage, resturent.OpenningTime,
    resturent.ClosingTime, resturent.Desription,resturent.ContactEmail, resturent.ContactNumber)
	if err != nil {
		fmt.Println("Error while creating new account: " + err.Error())
	}
  fmt.Println("Result is ", result)
}

func updateResturentDetails(w http.ResponseWriter, r *http.Request) {
  var err error
  resturent := domain.Resturents{}
  vars := mux.Vars(r)
	id := vars["id"]
  decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&resturent); err != nil {
    w.WriteHeader(http.StatusBadRequest)
    w.Write([]byte("internal server error"))
		return
	}
  fmt.Println("resturent Name is ", resturent)
  client := config.DbConnection()
	sqlUpdate := "UPDATE resturentDetails SET resturent_name = ? WHERE resturent_id = ?"
	_, err = client.Exec(sqlUpdate, resturent.ResturentName, id)
	if err != nil {
		fmt.Println("Error while querying resturents table " + err.Error())
	}
  fmt.Println("Restrrent table are updated")
}

func deleteResturentDetails(w http.ResponseWriter, r *http.Request) {
  var err error
  vars := mux.Vars(r)
	id := vars["id"]
  client := config.DbConnection()
	deleteQuery := "DELETE FROM resturentDetails WHERE resturent_id = ?"
	_, err = client.Exec(deleteQuery, id)
	if err != nil {
		fmt.Println("Error while querying resturents table " + err.Error())
	}
  fmt.Println("deleted resturent details ", id)
}
