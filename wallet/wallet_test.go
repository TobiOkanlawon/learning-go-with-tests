package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposits into the wallet propely", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, &wallet, 10)
	})

	t.Run("withdraws from the wallet properly", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(10)

		assertNoError(t, err)

		assertBalance(t, &wallet, 0)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, &wallet, 20)
		assertError(t, err, ErrorInsufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet *Wallet, amount int) {
	t.Helper()

	got := wallet.Balance()
	want := Bitcoin(amount)

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {

	t.Helper()

	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}

	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	if err != nil {
		t.Error("got an error but didn't expect one")
	}
}
