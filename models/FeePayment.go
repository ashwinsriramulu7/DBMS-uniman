package models

type FeePayment struct {
    ID          int
    StudentID   int
    Amount      float64
    DatePaid    string
    PaymentMode string
}

