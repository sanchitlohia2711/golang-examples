package main

import "fmt"

type account struct {
	id string
}

func newAccount() *account {
	return &account{
		id: "abc",
	}
}

func (a *account) checkAccount(accountID string) error {
	if a.id != accountID {
		return fmt.Errorf("AccountId is incorrect")
	}
	fmt.Println("Account Verified")
	return nil
}
