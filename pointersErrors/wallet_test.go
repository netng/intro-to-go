package pointersErrors

import "testing"

/*
Make test:
- check balance
- deposit
- withdraw, and handle error if insufficent funds

*/

func TestWallet(t *testing.T) {
	t.Run("balance", func(t *testing.T) {
		wallet := Wallet{50}

		want := Bitcoin(50)
		assertBalance(t, wallet, want)
	})

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(20)

		want := Bitcoin(20)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 50}
		err := wallet.Withdraw(20)

		want := Bitcoin(30)

		assertNoError(t, err)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(50))

		assertError(t, err, ErrInsuficientBalance)
		assertBalance(t, wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got: %s, want: %s", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("want some error message but didn't one")
	}

	if got != want {
		t.Errorf("got: %q, want: %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("got an error but didn't one")
	}
}
