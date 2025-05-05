package tests

import (
	"github.com/ashwinsriramulu7/DBMS-uniman/includes"
	"github.com/ashwinsriramulu7/DBMS-uniman/models"
	"github.com/ashwinsriramulu7/DBMS-uniman/modules"
	"testing"
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
func TestGetAndDeleteFeePayment(t *testing.T) {
    db := includes.InitDB()
    defer db.Close()

    fp := models.FeePayment{StudentID: 1, Amount: 1234.56, DatePaid: "2025-05-05", PaymentMode: "ONLINE"}
    modules.CreateFeePayment(fp)

    var id int
    err := db.QueryRow("SELECT id FROM fee_payment WHERE student_id = 1 AND amount = 1234.56 AND payment_mode = 'ONLINE'").Scan(&id)
    if err != nil {
        t.Fatalf("Failed to get inserted fee_payment ID: %v", err)
    }

    got := modules.GetFeePaymentByID(id)
    if got.StudentID != fp.StudentID || got.Amount != fp.Amount || got.DatePaid != fp.DatePaid || got.PaymentMode != fp.PaymentMode {
        t.Errorf("GetFeePaymentByID failed: expected %+v, got %+v", fp, got)
    }

    modules.DeleteFeePaymentByID(id)
    err = db.QueryRow("SELECT id FROM fee_payment WHERE id = ?", id).Scan(&id)
    if err == nil {
        t.Error("DeleteFeePaymentByID failed: record still exists")
    }
}

