package main

import "fmt"

type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
}

func newWalletFacade(code int) *walletFacade {
	fmt.Println("Starting create account")
	walletFacacde := &walletFacade{
		account:      newAccount(),
		securityCode: newSecurityCode(code),
		wallet:       newWallet(),
		notification: &notification{},
	}
	fmt.Println("Account created")
	return walletFacacde
}

func (w *walletFacade) addMoneyToWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	return nil
}
