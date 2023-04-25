package fintech

import "testing"

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
	
		got := wallet.Balance()
		expected := Bitcoin(10).String()
	
		if got != expected {
			t.Errorf("Got %s but received %s", got, expected)
		}
	});

	assertBalance := func (t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want.String() {
			t.Errorf("Expected %s but got %s", want, got)
		}
	}

	assertError := func (t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("didn't get an error but wanted one") // Fatal prevents the test to continue, if it fails it stops here
		}

		if got != want {
			t.Errorf("got %q but want %q", got, want)
		}
	}
	
	t.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(10)
		wallet.Withdraw(5)

		got := wallet.Balance()
		want := "5 BTC"

		if got != want {
			t.Errorf("Got %s but expected %s", got, want)
		}
	})

	t.Run("Withdraw with insufficient funds", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		err := wallet.Withdraw(20)

		assertError(t, err, ErrorInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(10))
	})
}