package models

type Status string

type BankAccount struct {
	ID       string
	BankName string
	Balance  int
	Status   bool
}

func BankAccountFactory(bankName string, initialBalance float64) BankAccount {
	return BankAccount{BankName: bankName, Balance: int(initialBalance * 100), Status: initialBalance > 0}
}

func (acc *BankAccount) Withdrawal(expense Expense) {
	acc.Balance -= int(expense.Value * 100)
	if acc.Balance < 0 {
		acc.Status = false
	}
}

func (acc *BankAccount) Deposit(revenue Revenue) {
	acc.Balance += int(revenue.Value * 100)
	if acc.Balance >= 0 {
		acc.Status = true
	}
}

func (acc *BankAccount) CheckBalance() float64 {
	return float64(acc.Balance / 100)
}
