package bmo

import "time"

type Statistics struct {
	UserID          string    `json:"user_id,omitempty"`
	Username        string    `json:"username,omitempty"`
	ParsedItems     int       `json:"parsed_items,omitempty"`
	Start           time.Time `json:"start,omitempty"`
	End             time.Time `json:"end,omitempty"`
	TotalInRecords  int       `json:"total_in_records,omitempty"`
	TotalInMoney    float64   `json:"total_in_money,omitempty"`
	TotalOutRecords int       `json:"total_out_records,omitempty"`
	TotalOutMoney   float64   `json:"total_out_money,omitempty"`
}

// BMO的账单相当简单，只有4个字段
// Transaction Type,Date Posted, Transaction Amount,Description
type Order struct {
	PayTime                time.Time // 记账日期
	TransactionDescription string    // 交易描述(包括交易对手及摘要)
	Money                  float64   // 记账金额(收入/支出)
	Type                   OrderType // 收/支 (数据中无该列，推测而来)
}

type OrderType string

const (
	OrderTypeSend    OrderType = "支出"
	OrderTypeRecv    OrderType = "收入"
	OrderTypeUnknown OrderType = "Unknown"
)

type TransactionType string

const (
	TransactionTypeCredit  TransactionType = "CREDIT"
	TransactionTypeDebit   TransactionType = "DEBIT"
	TransactionTypeUnknown TransactionType = "UNKNOWN"
)

type CardMode string

const (
	DebitMode  CardMode = "Debit"
	CreditMode CardMode = "Credit"
)

// LocalTimeFmt set time format to utc-7
const LocalTimeFmt = "20060102 -0700"
