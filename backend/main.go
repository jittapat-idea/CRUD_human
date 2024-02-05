package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)


// func connect() (*sql.DB, error) {
// 	bin, err := ioutil.ReadFile("/run/secrets/db-password")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return sql.Open("postgres", fmt.Sprintf("postgres://postgres:%s@db:5432/example?sslmode=disable", string(bin)))
// }

// func humanHandler(w http.ResponseWriter, r *http.Request) {
// 	db, err := connect()
// 	if err != nil {
// 		w.WriteHeader(500)
// 		return
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT ID, NAME, location FROM Human")
// 	if err != nil {
// 		w.WriteHeader(500)
// 		return
// 	}

// 	var humans []map[string]interface{}
// 	for rows.Next() {
// 		var id int
// 		var name, location string
// 		err = rows.Scan(&id, &name, &location)

// 		human := map[string]interface{}{
// 			"ID":       id,
// 			"NAME":     name,
// 			"location": location,
// 		}
// 		humans = append(humans, human)
// 	}

// 	json.NewEncoder(w).Encode(humans)
// }

// func gethumans(w http.ResponseWriter, r *http.Request) {
// 	db, err := connect()
// 	if err != nil {
// 		w.WriteHeader(500)
// 		return
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT title FROM blog")
// 	if err != nil {
// 		w.WriteHeader(500)
// 		return
// 	}
// 	var titles []string
// 	for rows.Next() {
// 		var title string
// 		err = rows.Scan(&title)
// 		titles = append(titles, title)
// 	}
// 	json.NewEncoder(w).Encode(titles)
	
// }

// func main() {
// 	log.Print("Prepare db...")
// 	if err := prepare(); err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Print("Listening 8000")
// 	r := mux.NewRouter()
// 	r.HandleFunc("/api/humans", gethumans)
// 	log.Fatal(http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r)))
// }

// func prepare() error {
// 	db, err := connect()
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	for i := 0; i < 60; i++ {
// 		if err := db.Ping(); err == nil {
// 			break
// 		}
// 		time.Sleep(time.Second)
// 	}


// 	// Create "Human" table
// 	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS Human (ID SERIAL, NAME VARCHAR, location VARCHAR)"); err != nil {
// 		return err
// 	}

// 	return nil
// }














//----------------------------------- อันดังเดิม --------------------------------------------------------------
// func connect() (*sql.DB, error) {
// 	bin, err := ioutil.ReadFile("/run/secrets/db-password")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return sql.Open("postgres", fmt.Sprintf("postgres://postgres:%s@db:5432/example?sslmode=disable", string(bin)))
// }

// func humanHandler(w http.ResponseWriter, r *http.Request) {
// 	db, err := connect()
// 	if err != nil {
// 		w.WriteHeader(500)
// 		return
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT ID, NAME, location FROM Human")
// 	if err != nil {
// 		w.WriteHeader(500)
// 		return
// 	}

// 	var humans []map[string]interface{}
// 	for rows.Next() {
// 		var id int
// 		var name, location string
// 		err = rows.Scan(&id, &name, &location)

// 		human := map[string]interface{}{
// 			"ID":       id,
// 			"NAME":     name,
// 			"location": location,
// 		}
// 		humans = append(humans, human)
// 	}

// 	json.NewEncoder(w).Encode(humans)
// }

// func blogHandler(w http.ResponseWriter, r *http.Request) {
// 	db, err := connect()
// 	if err != nil {
// 		w.WriteHeader(500)
// 		return
// 	}
// 	defer db.Close()

// 	rows, err := db.Query("SELECT title FROM blog")
// 	if err != nil {
// 		w.WriteHeader(500)
// 		return
// 	}
// 	var titles []string
// 	for rows.Next() {
// 		var title string
// 		err = rows.Scan(&title)
// 		titles = append(titles, title)
// 	}
// 	json.NewEncoder(w).Encode(titles)
// }

// func main() {
// 	log.Print("Prepare db...")
// 	if err := prepare(); err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Print("Listening 8000")
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", blogHandler)
// 	r.HandleFunc("/humans", humanHandler)
// 	log.Fatal(http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r)))
// }

// func prepare() error {
// 	db, err := connect()
// 	if err != nil {
// 		return err
// 	}
// 	defer db.Close()

// 	for i := 0; i < 60; i++ {
// 		if err := db.Ping(); err == nil {
// 			break
// 		}
// 		time.Sleep(time.Second)
// 	}

// 	// Drop existing "blog" table
// 	if _, err := db.Exec("DROP TABLE IF EXISTS blog"); err != nil {
// 		return err
// 	}

// 	// Create "Human" table
// 	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS Human (ID SERIAL, NAME VARCHAR, location VARCHAR)"); err != nil {
// 		return err
// 	}

// 	// Insert sample data into "Human" table
// 	for i := 0; i < 5; i++ {
// 		if _, err := db.Exec("INSERT INTO Human (NAME, location) VALUES ($1, $2);", fmt.Sprintf("Person #%d", i), fmt.Sprintf("Location #%d", i)); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
//----------------------------------------------------------------------------------------------




























// type Human struct {
// 	ID       int    `json:"id"`
// 	NAME     string `json:"name"`
// 	Location string `json:"location"`
// }
// //ตัวเเปร human
// var humans []Human
// //ตัวแปร global สำหรับเก็บ connection ของฐานข้อมูล
// var db *sql.DB

// //-----------------------------------------------------------------------------------------------------
// func connect() error {
// 	bin, err := ioutil.ReadFile("/run/secrets/db-password")
// 	if err != nil {
// 		return err
// 	}

// 	db, err = sql.Open("postgres", fmt.Sprintf("postgres://postgres:%s@db:5432/example?sslmode=disable", string(bin)))
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func getDB() (*sql.DB, error) {
// 	if db == nil {
// 		if err := connect(); err != nil {
// 			return nil, err
// 		}
// 	}
// 	return db, nil
// }

// func closeDB() {
// 	if db != nil {
// 		db.Close()
// 	}
// }
// //-----------------------------------------------------------------------------------------------------



// // function Read ในการอ่านข้อมูลทังหมด
// func gethumans(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
//     json.NewEncoder(w).Encode(humans)
// }

// // function Read ในการอ่านข้อมูลเเค่ 1 คน
// func gethuman(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) // Get params

// 	// Convert param id to int
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Loop through humans and find with id
// 	for _, person := range humans {
// 		if person.ID == id {
// 			json.NewEncoder(w).Encode(person)
// 			return
// 		}
// 	}

// 	// If no matching ID found
// 	json.NewEncoder(w).Encode(&Human{})
// }

// // function create ในการสร้างคน
// func createHuman(w http.ResponseWriter, r *http.Request){
// 	w.Header().Set("Content-Type", "application/json")
// 	var human Human 

// 	_  = json.NewDecoder(r.Body).Decode(&human)
// 	humans = append(humans, human)
// 	json.NewEncoder(w).Encode(human)

// }

// // Function updateHuman ในการ update ข้อมูล
// func updateHuman(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	// Convert param id to int
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Find the index of the person with the given id
// 	var index int
// 	var found bool
// 	for i, person := range humans {
// 		if person.ID == id {
// 			index = i
// 			found = true
// 			break
// 		}
// 	}

// 	// If the person with the given id is not found, return 404 Not Found
// 	if !found {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	// Decode the request body into a new human
// 	var updatedHuman Human
// 	if err := json.NewDecoder(r.Body).Decode(&updatedHuman); err != nil {
// 		http.Error(w, "Invalid JSON", http.StatusBadRequest)
// 		return
// 	}

// 	// Update only the fields that are provided in the request
// 	// If a field is empty in the request, keep the existing value
// 	humanToUpdate := &humans[index]
// 	if updatedHuman.NAME != "" {
// 		humanToUpdate.NAME = updatedHuman.NAME
// 	}
// 	if updatedHuman.Location != "" {
// 		humanToUpdate.Location = updatedHuman.Location
// 	}

// 	// Send the updated human in the response
// 	json.NewEncoder(w).Encode(humanToUpdate)
// }

// // Function deleteHuman ในการลบข้อมูล
// func deleteHuman(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)

// 	// Convert param id to int
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}

// 	// Find the index of the person with the given id
// 	var index int
// 	var found bool
// 	for i, person := range humans {
// 		if person.ID == id {
// 			index = i
// 			found = true
// 			break
// 		}
// 	}

// 	// If the person with the given id is not found, return 404 Not Found
// 	if !found {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	// Remove the person from the humans slice
// 	humans = append(humans[:index], humans[index+1:]...)

// 	// Send the updated list of humans in the response
// 	json.NewEncoder(w).Encode(humans)
// }


// func main(){
// 	// Init router
// 	r := mux.NewRouter()
//     //test
// 	humans = append(humans, Human{ID: 1,NAME:"jittapat Ausakul",Location:"dusit"})
// 	humans = append(humans, Human{ID: 2,NAME:"pon chon",Location:"japan"})
// 	// Route handles & endpoints
// 	r.HandleFunc("/api/humans", gethumans).Methods("GET")
// 	r.HandleFunc("/api/humans/{id}", gethuman).Methods("GET")
// 	r.HandleFunc("/api/humans", createHuman).Methods("POST")
// 	r.HandleFunc("/api/humans/{id}", updateHuman).Methods("PUT")
// 	r.HandleFunc("/api/humans/{id}", deleteHuman).Methods("DELETE")

// 	log.Fatal(http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r)))


// }





























//-------------------------------------------------------------------------------------------------------
type Human struct {
	ID       int    `json:"id"`
	NAME     string `json:"name"`
	Location string `json:"location"`
}
//ตัวเเปร human
var humans []Human


//-----------------------------------------------------------------------------------------------------
//ฟังก์ชันที่ใช้เชื่อมต่อกับฐานข้อมูล PostgreSQL
func connect() (*sql.DB, error) {

	//ใช้อ่านข้อมูลจากไฟล์ db-password
	bin, err := ioutil.ReadFile("/run/secrets/db-password")
	//ไม่มีข้อผิดพลาดในขั้นตอนการเชื่อมต่อ, ฟังก์ชันจะคืนค่า *sql.DB พร้อม nil ในส่วนของ error.
	if err != nil {
		return nil, err
	}
	//ใช้ sql.Open เพื่อเปิดการเชื่อมต่อกับฐานข้อมูล PostgreSQL โดยใช้ URL connection string ที่ได้จากการฟอร์แมต fmt.Sprintf และรหัสผ่านที่ได้จากไฟล์ db-password ที่ถูกอ่านมา.
	return sql.Open("postgres", fmt.Sprintf("postgres://postgres:%s@db:5432/example?sslmode=disable", string(bin)))
}
//-----------------------------------------------------------------------------------------------------


//http.ResponseWriter ซึ่งใช้ในการสร้าง HTTP response เพื่อส่งกลับไปยัง client.
//http.Request ซึ่งเก็บข้อมูลเกี่ยวกับ HTTP request ที่ถูกส่งมา.

func Home(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "<h1>Hello, world!</h1>")
}

// function Read ดึงข้อมูลทั้งหมดจากฐานข้อมูลและส่งข้อมูลนั้นกลับไปให้กับ client ในรูปแบบ JSON response.
func gethumans(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	db, err := connect()//เรียกใช้ฟังก์ชัน connect เพื่อเชื่อมต่อกับฐานข้อมูล Postgres
	if err != nil {
		w.WriteHeader(500)
	    return
	}
	defer db.Close()//ใช้ defer เพื่อทำให้ฟังก์ชัน db.Close() ถูกเรียกหลังจากที่ฟังก์ชัน gethumans ทำงานเสร็จ.
	rows, err := db.Query("SELECT * FROM humans")//ดึงข้อมูลทั้งหมดจากตาราง humans ในฐานข้อมูล.
	if err != nil {
		//fmt.Fprintf(w,"%s",err)
		w.WriteHeader(500)
	    return
	}

	var humansFromDB []Human//สร้าง slice เพื่อเก็บข้อมูลที่ดึงมาจากฐานข้อมูล.

	for rows.Next() {//วนลูปเพื่อสร้าง struct Human จากข้อมูลที่ดึงมา และนำไปเก็บใน slice humansFromDB.
		var human Human
		err := rows.Scan(&human.ID, &human.NAME, &human.Location)//นำข้อมูลที่ดึงมาเก็บใน struct Human.
		if err != nil {
			log.Fatal(err)
		}
		humansFromDB = append(humansFromDB, human)
	}
	//ใช้ json.NewEncoder เพื่อเขียนข้อมูลในรูปแบบ JSON และส่งกลับไปให้กับ client ผ่าน http.ResponseWriter
	json.NewEncoder(w).Encode(humansFromDB)
}

// function Read ดึงข้อมูลของบุคคลที่มี ID ที่ระบุจากฐานข้อมูลและส่งข้อมูลนั้นกลับไปให้กับ client ในรูปแบบ JSON response.
func gethuman(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)//เพื่อดึงค่า parameter ที่ระบุใน URL จาก request.

	id, err := strconv.Atoi(params["id"])//แปลงค่า ID จาก string เป็น integer โดยใช้ strconv.Atoi
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	db, err := connect()//เชื่อมต่อกับฐานข้อมูล Postgres
	if err != nil {
		w.WriteHeader(500)
	    return
	}
	defer db.Close()//ใช้ defer เพื่อทำให้ฟังก์ชัน db.Close() ถูกเรียกหลังจากที่ฟังก์ชัน gethumans ทำงานเสร็จ.

	///ดึงข้อมูลของบุคคลที่มี ID ที่ระบุจากตาราง humans ในฐานข้อมูล.
	row := db.QueryRow("SELECT * FROM humans WHERE id = $1", id)

	var human Human
	err = row.Scan(&human.ID, &human.NAME, &human.Location)//นำข้อมูลที่ดึงมาเก็บใน struct Human.
	if err != nil {
		http.NotFound(w, r)
		return
	}
    ////ใช้ json.NewEncoder เพื่อเขียนข้อมูลในรูปแบบ JSON และส่งกลับไปให้กับ client ผ่าน http.ResponseWriter
	json.NewEncoder(w).Encode(human)
}

// function create ในการสร้างคนนำข้อมูลนี้ไปเพิ่มในฐานข้อมูล Postgres ด้วยคำสั่ง SQL INSERT.
func createHuman(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")//ตั้งค่า Header ใน response เพื่อระบุว่าข้อมูลที่ส่งกลับจะเป็น JSON.

	var human Human// สร้างตัวแปร human ในรูปแบบ Human(struct) เพื่อเก็บข้อมูลที่รับมา.

	///ใช้ json.NewDecoderเพื่อถอดรหัสข้อมูล JSON ที่ถูกส่งมาจาก request body และเก็บลงในตัวแปร human
	_ = json.NewDecoder(r.Body).Decode(&human)
	db, err := connect()//เชื่อมต่อกับฐานข้อมูล Postgres
	if err != nil {
		w.WriteHeader(500)
	    return
	}
	defer db.Close()//ใช้ defer เพื่อทำให้ฟังก์ชัน db.Close() ถูกเรียกหลังจากที่ฟังก์ชัน gethumans ทำงานเสร็จ.

	//เพิ่มข้อมูลใหม่ลงในตาราง humans ของฐานข้อมูล Postgres โดยใช้คำสั่ง SQL INSERT. 
	//$1 และ $2 คือ placeholder ที่จะถูกแทนที่ด้วยค่าจริงของ human.NAME และ human.Location.
	_, err = db.Exec("INSERT INTO humans (name, location) VALUES ($1, $2)", human.NAME, human.Location)
	if err != nil {//มี error เเจ้งเตือนหยุดการทำงาน
		log.Fatal(err)
	}
    ////ใช้ json.NewEncoder เพื่อเขียนข้อมูลในรูปแบบ JSON และส่งกลับไปให้กับ client ผ่าน http.ResponseWriter
	json.NewEncoder(w).Encode(human)
}

// Function updateHuman ในการ update ข้อมูล
//อัปเดตข้อมูลของบุคคลที่มี ID ที่ระบุใน request URL จาก client โดยใช้ข้อมูลที่รับมาจาก request body ในรูปแบบ JSON
func updateHuman(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")////ตั้งค่า Header ใน response เพื่อระบุว่าข้อมูลที่ส่งกลับจะเป็น JSON.
	params := mux.Vars(r)//ใช้ mux.Vars เพื่อดึงค่า parameter จาก URL ที่มี key เป็น "id" และเก็บไว้ในตัวแปร params.

	id, err := strconv.Atoi(params["id"])////แปลงค่า ID จาก string เป็น integer โดยใช้ strconv.Atoi
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	db, err := connect()//เชื่อมต่อกับฐานข้อมูล Postgres. 
	if err != nil {
		w.WriteHeader(500)
	    return
	}
	defer db.Close()//ใช้ defer เพื่อให้ db.Close() ถูกเรียกหลังจากที่ฟังก์ชัน updateHuman ทำงานเสร็จ.

	//ดึงข้อมูลของบุคคลที่มี ID ที่ระบุจากตาราง humans ในฐานข้อมูล.
	row := db.QueryRow("SELECT * FROM humans WHERE id = $1", id)

	var existingHuman Human// สร้างตัวแปร existingHuman ในรูปแบบ Human(struct) เพื่อเก็บข้อมูลที่รับมา.
	// ดึงข้อมูลบุคคลจาก row ที่ได้มาและเก็บในตัวแปร existingHuman.
	err = row.Scan(&existingHuman.ID, &existingHuman.NAME, &existingHuman.Location)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	var updatedHuman Human// สร้างตัวแปร updatedHuman ในรูปแบบ Human(struct) เพื่อเก็บข้อมูลที่รับมา.
	err = json.NewDecoder(r.Body).Decode(&updatedHuman)//ถอดรหัสข้อมูล JSON จาก request body และเก็บในตัวแปร updatedHuman.
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//ถ้ามีข้อมูล NAME ใน updatedHuman ให้กำหนดค่าใน existingHuman.NAME ใหม่.
	if updatedHuman.NAME != "" {
		existingHuman.NAME = updatedHuman.NAME
	}
	//ถ้ามีข้อมูล location ใน updatedHuman ให้กำหนดค่าใน existingHuman.Location ใหม่.
	if updatedHuman.Location != "" {
		existingHuman.Location = updatedHuman.Location
	}

	// ใช้คำสั่ง SQL UPDATE เพื่ออัปเดตข้อมูลบุคคลในฐานข้อมูล Postgres.
	_, err = db.Exec("UPDATE humans SET name = $1, location = $2 WHERE id = $3", existingHuman.NAME, existingHuman.Location, existingHuman.ID)
	if err != nil {
		log.Fatal(err)
	}
    //ใช้ json.NewEncoder เพื่อเขียนข้อมูลในรูปแบบ JSON และส่งกลับไปให้กับ client ผ่าน http.ResponseWriter
	json.NewEncoder(w).Encode(existingHuman)
}


// Function deleteHuman ลบข้อมูลของบุคคลที่มี ID ที่ระบุใน request URL จาก client ออกจากฐานข้อมูล Postgres:
func deleteHuman(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")//ถอดรหัสข้อมูล JSON จาก request body และเก็บในตัวแปร updatedHuman.
	params := mux.Vars(r)//ใช้ mux.Vars เพื่อดึงค่า parameter จาก URL ที่มี key เป็น "id" และเก็บไว้ในตัวแปร params.

	id, err := strconv.Atoi(params["id"])//แปลงค่า id ที่มีประเภท string เป็น int.
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	db, err := connect() //เชื่อมต่อกับฐานข้อมูล Postgres
	if err != nil {
		w.WriteHeader(500)
	    return
	}
	defer db.Close()//ใช้ defer เพื่อให้ db.Close() ถูกเรียกหลังจากที่ฟังก์ชัน deleteHuman ทำงานเสร็จ.

	//  ค้นหาข้อมูลบุคคลที่มี ID ที่ระบุในฐานข้อมูล Postgres.
	row := db.QueryRow("SELECT * FROM humans WHERE id = $1", id)

	var existingHuman Human// สร้างตัวแปร existingHuman ในรูปแบบ Human(struct) เพื่อเก็บข้อมูลที่รับมา.
	// ดึงข้อมูลบุคคลจาก row ที่ได้มาและเก็บในตัวแปร existingHuman.
	err = row.Scan(&existingHuman.ID, &existingHuman.NAME, &existingHuman.Location)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	//ใช้คำสั่ง SQL DELETE เพื่อลบข้อมูลบุคคลที่มี ID ที่ระบุออกจากฐานข้อมูล Postgres.
	_, err = db.Exec("DELETE FROM humans WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}

	// ส่งข้อมูลบุคคลที่ถูกลบออกจากฐานข้อมูล Postgres กลับไปยัง client ในรูปแบบ JSON.
	json.NewEncoder(w).Encode(existingHuman)
}

//ที่เรียกใช้งานการเชื่อมต่อฐานข้อมูลและเริ่มต้นเซิร์ฟเวอร์ HTTP:
func main() {
	log.Print("Prepare db...")//print log ว่ากำลังเตรียมฐานข้อมูล.
	if err := prepare(); err != nil {//เรียกฟังก์ชัน prepare ซึ่งน่าจะเป็นการเตรียมข้อมูลหรือตารางในฐานข้อมูล 
		log.Fatal(err)
	}

	log.Print("Listening 8000")
	// Init router
	r := mux.NewRouter()//สร้าง Router โดยใช้ mux.NewRouter() จาก package mux เพื่อจัดการ route ของ HTTP requests.

	// Route handles & endpoints
	//กำหนด route สำหรับ HTTP GET,POST,PUT,DELETE ที่ path ต่างๆ ให้เรียกใช้ฟังก์ชัน ต่างๆ
	r.HandleFunc("/",Home).Methods("GET")
	r.HandleFunc("/api/humans", gethumans).Methods("GET")
	r.HandleFunc("/api/humans/{id}", gethuman).Methods("GET")
	r.HandleFunc("/api/humans/create", createHuman).Methods("POST")
	r.HandleFunc("/api/humans/update/{id}", updateHuman).Methods("PUT")
	r.HandleFunc("/api/humans/delete/{id}", deleteHuman).Methods("DELETE")

	// Start the server เริ่มต้นเซิร์ฟเวอร์ HTTP ที่ port 8000 โดยใช้ http.ListenAndServe 
	//และใส่ handlers.LoggingHandler เพื่อบันทึก log ของการเข้าถึง.
	log.Fatal(http.ListenAndServe(":8000", handlers.LoggingHandler(os.Stdout, r)))
}

//ทำหน้าที่เตรียมฐานข้อมูล
func prepare() error {
	// เพื่อเชื่อมต่อกับฐานข้อมูล 
	db, err := connect()
	if err != nil {
		return err
	}
	defer db.Close()// นิยามการปิดการเชื่อมต่อกับฐานข้อมูลเมื่อฟังก์ชันทำงานเสร็จ.
    //ในการรอฐานข้อมูลพร้อมทำงาน, ตรวจสอบการเชื่อมต่อฐานข้อมูลทุกวินาทีเป็นจำนวน 60 ครั้ง.
	for i := 0; i < 60; i++ {
		if err := db.Ping(); err == nil {
			break
		}
		time.Sleep(time.Second)
	}

		// ลบตาราง "humans" ถ้ามีอยู่.
	if _, err := db.Exec("DROP TABLE IF EXISTS humans"); err != nil {
		return err
	}

	// สร้างตาราง "humans" ถ้าไม่มี.
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS humans (ID SERIAL, NAME VARCHAR, location VARCHAR)"); err != nil {
		return err
	}

	// แทรกรายการข้อมูลตัวอย่างลงในตาราง "humans".
	if _, err := db.Exec("INSERT INTO humans (NAME, location) VALUES ($1, $2);", "jittapat ausakul", "bangkok"); err != nil {
		return err
	}
	return nil//ถ้าไม่มีข้อผิดพลาดทุกรายการ, จะส่งค่า nil กลับ.
}


