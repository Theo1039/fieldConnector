package bibleStudy

import (
    "fieldConnector/stdInput"
    "database/sql"
    "log"
    "bufio"
    "fmt"
    "os"
)

type BibleStudy struct{
    Id int
    Name string
    PhoneNumber string
    HouseNumber string
    StreetName string
    Book_brochure_name string
    Chapter string
    Paragraph string
    Next_appointment string
    QuestionAsked string
}

func CreateBibleStudy(db *sql.DB){
    reader := bufio.NewReader(os.Stdin)
    b := BibleStudy{
        Name: stdInput.GetInput(reader, "Enter Name: "),
        PhoneNumber: stdInput.GetInput(reader, "Enter Phone number: "),
        HouseNumber: stdInput.GetInput(reader, "Enter House number: "),
        StreetName: stdInput.GetInput(reader, "Enter Street Name: "),
        Book_brochure_name: stdInput.GetInput(reader, "Enter Book Or Brochure Name: "),
        Chapter: stdInput.GetInput(reader, "Enter Chapter: "),
        Paragraph: stdInput.GetInput(reader, "Enter Paragraph: "),
        Next_appointment: stdInput.GetInput(reader, "Enter Next Appointment Date: "),
        QuestionAsked: stdInput.GetInput(reader, "Enter Question Asked: "),
    }

    query := `CREATE TABLE IF NOT EXISTS bible_study(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    phoneNumber TEXT, 
    houseNumber TEXT,
    streetName TEXT,
    book_brochure_name TEXT,
    Chapter TEXT,
    paragraph TEXT,
    next_appointment TEXT,
    questionAsked TEXT
    )`

    _, err := db.Exec(query)
    if err != nil{
        log.Fatalf("Creating Database Table Failed %v", err)
    }
    query = `INSERT INTO bible_study(name,phoneNumber,
    houseNumber,streetName,book_brochure_name,chapter,paragraph,
    next_appointment,questionAsked) VALUES(?,?,?,?,?,?,?,?,?)`

    _,err = db.Exec(query,b.Name,b.PhoneNumber,b.HouseNumber,b.StreetName,b.Book_brochure_name,b.Chapter,b.Paragraph,b.Next_appointment,b.QuestionAsked)
    if err != nil {
       log.Fatalf("Database Insertion Failed %v", err)
    }
    fmt.Println("Blble Study saved successfully!")
}

func DisplayBibleStudy(db *sql.DB){

    query := `SELECT
    id,
    COALESCE(name, ""),
    COALESCE(phoneNumber, ""),
    COALESCE(houseNumber,""),
    COALESCE(streetName,""),
    COALESCE(book_brochure_name,""),
    COALESCE(chapter,""),
    COALESCE(paragraph,""),
    COALESCE(next_appointment,""),
    COALESCE(questionAsked,"")
    FROM bible_study WHERE 1=1`

    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    printRows(rows)

    }
    


func printRows(rows *sql.Rows) {
	count := 0

	for rows.Next() {
        var b BibleStudy
		
        err := rows.Scan(&b.Id, &b.Name, &b.PhoneNumber, &b.HouseNumber, &b.StreetName, &b.Book_brochure_name, &b.Chapter, &b.Paragraph, &b.Next_appointment, &b.QuestionAsked)
		
		if err != nil {
			fmt.Println("Error scanning row data:", err)
			return
		}
		count++
        fmt.Printf("ID: %d | Name: %s | Phone No: %s | House No: %s | Street No: %s | Book Or Brochure: %s | Chapter: %s | Paragraph: %s | Next Appointment: %s | Question: %s\n",
        b.Id,b.Name,b.PhoneNumber,b.HouseNumber,b.StreetName,b.Book_brochure_name,b.Chapter,b.Paragraph,b.Next_appointment,b.QuestionAsked)
		
	}

	if count == 0 {
		fmt.Println("No matching return visits found.")
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during row processing:", err)
	}
}