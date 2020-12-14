package ptrs

import "testing"

func TestWallet(t *testing.T) {
	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didnt want one")
		}
	}
	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but did not get one")
		}
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))

	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(Bitcoin(5))
		assertBalance(t, wallet, Bitcoin(5))
		assertNoError(t, err)
	})
	t.Run("insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(Bitcoin(15))
		assertBalance(t, wallet, Bitcoin(10))
		assertError(t, err, ErrInsufficientFunds)
	})

}
