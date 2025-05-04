package modules

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"log"
)

func CreateFeePayment(f models.FeePayment) {
	db := includes.InitDB()
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO fee_payment(student_id, amount, date_paid, payment_mode) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(f.StudentID, f.Amount, f.DatePaid, f.PaymentMode)
	if err != nil {
		log.Fatal(err)
	}
}
