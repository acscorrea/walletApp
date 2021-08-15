package models

import "testing"

var revenue = Revenue{"1", "Receita teste", 100.00}
var expense = Expense{"1", "Despesa teste", 100.00}

func TestAccount(t *testing.T) {

	t.Run("Factory test, positive balance", func(t *testing.T) {
		acc := BankAccountFactory("MockBank", 1200.00)
		AssertEqualFloat(t, acc.CheckBalance(), 1200.00)
		AssertEqualString(t, acc.BankName, "MockBank")
		AssertEqualBool(t, acc.Status, true)
	})

	t.Run("Factory test, negative balance", func(t *testing.T) {
		acc := BankAccountFactory("MockBank", -1200.00)
		AssertEqualFloat(t, acc.CheckBalance(), -1200.00)
		AssertEqualString(t, acc.BankName, "MockBank")
		AssertEqualBool(t, acc.Status, false)
	})

	t.Run("Deposit test", func(t *testing.T) {
		acc := BankAccount{}
		acc.Deposit(revenue)
		got := acc.CheckBalance()
		AssertEqualFloat(t, got, revenue.Value)
	})

	t.Run("Withdrawal test", func(t *testing.T) {
		acc := BankAccount{}
		acc.Withdrawal(expense)
		got := acc.CheckBalance()
		AssertEqualFloat(t, got, -expense.Value)
	})
}

func AssertEqualFloat(t *testing.T, got, expected float64) {
	if got != expected {
		t.Errorf("Expeceted %.2f, received %.2f", got, expected)
	}
}

func AssertEqualString(t *testing.T, got, expected string) {
	if got != expected {
		t.Errorf("Expeceted %s, received %s", got, expected)
	}
}

func AssertEqualBool(t *testing.T, got, expected bool) {
	if got != expected {
		t.Errorf("Expeceted %t, received %t", got, expected)
	}
}
