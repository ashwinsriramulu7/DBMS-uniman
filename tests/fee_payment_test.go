package tests

import (
    "testing"
    "github.com/ashwinsriramulu7/DBMS-uniman/models"
    "github.com/ashwinsriramulu7/DBMS-uniman/modules"
    "github.com/ashwinsriramulu7/DBMS-uniman/includes"
)
func TestCreateFeePayment(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    fp := models.FeePayment{StudentID: 1, Amount: 5000.50, DatePaid: "2024-04-15", PaymentMode: "CARD"}
    modules.CreateFeePayment(fp)

    var amount float64
    err := db.QueryRow("SELECT amount FROM fee_payment WHERE student_id = 1 AND payment_mode = 'CARD'").Scan(&amount)
    if err != nil || amount != 5000.50 {
        t.Error("FeePayment insert failed")
    }

    db.Exec("DELETE FROM fee_payment WHERE student_id = 1 AND payment_mode = 'CARD'")
}

