package data

import (
	"encoding/json"
	"errors"
	"log"

	_db "../db"
)

var (
	// ErrInvalidText .
	ErrInvalidText = errors.New("invalid text")
	//ErrInvalidID .
	ErrInvalidID = errors.New("invalid id")
	//ErrInvalidJSON .
	ErrInvalidJSON = errors.New("invalid json")
)

// Datajson .
type Datajson struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Data .
type Data struct {
	ID   int
	Text string
}

func (p *Datajson) String() string {
	b, err := json.Marshal(p)
	if err != nil {
		err = ErrInvalidJSON
		log.Println("ERR:", err)
		return err.Error()
	}

	return string(b)
}

// New .
func New(text string) (err error) {
	// verify text
	if text == "" {
		err = ErrInvalidText
		log.Println("ERR : ", err)
		return
	}
	conn := _db.Conn()
	insert, err := conn.Prepare("INSERT INTO data(text) VALUES(?)")
	insert.Exec(text)

	if err != nil {
		panic(err.Error)
	}
	defer conn.Close()
	log.Println("data saved for ", text)
	return
}

// GetAll .
func GetAll() (data []Data, err error) {
	log.Println("fetch the entire table", nil)
	d := Data{}
	res := []Data{}

	conn := _db.Conn()
	selDB, err := conn.Query("SELECT * FROM data ORDER BY id")
	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		var id int
		var text string
		err = selDB.Scan(&id, &text)
		if err != nil {
			panic(err.Error())
		}
		d.ID = id
		d.Text = text
		res = append(res, d)
	}
	log.Println(len(res), " rows found")

	return res, nil
}

// Get .
func Get(id int) (data Data, err error) {
	log.Println("Selecting data for id : ", id)

	if id == 0 {
		err = ErrInvalidID
		log.Println("ERR : ", err)
		return
	}

	conn := _db.Conn()
	selDB, err := conn.Query("SELECT * FROM data WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}
	for selDB.Next() {
		var id int
		var text string
		err = selDB.Scan(&id, &text)
		if err != nil {
			panic(err.Error())
		}
		data.ID = id
		data.Text = text
	}
	log.Println("details found for ", id)
	defer conn.Close()
	return data, nil
}

// Edit .
func Edit(data Data) (err error) {
	if data.ID == 0 {
		err = ErrInvalidID
	}
	if data.Text == "" {
		err = ErrInvalidText
	}
	if err != nil {
		log.Println("ERR : ", err)
		return
	}
	conn := _db.Conn()
	update, err := conn.Prepare("UPDATE data SET text = ? WHERE id = ?")
	update.Exec(data.Text, data.ID)

	if err != nil {
		panic(err.Error)
	} else {
		log.Println("Updated details of ", data.ID)
	}

	defer conn.Close()
	return
}

// Delete .
func Delete(id int) (err error) {
	if id == 0 {
		err = ErrInvalidID
		log.Println("ERR : ", err)
		return
	}
	conn := _db.Conn()
	delete, err := conn.Prepare("DELETE FROM data WHERE id = ?")
	delete.Exec(id)

	if err != nil {
		panic(err.Error)
	} else {
		log.Println("Deleted details of ", id)
	}

	defer conn.Close()
	return
}
