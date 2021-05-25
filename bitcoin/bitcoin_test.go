package bitcoin

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := "10 BTC"

		assert.Equal(t, want, fmt.Sprint(got))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)

		balance := wallet.Balance()

		assert.Nil(t, err)
		assert.Equal(t, Bitcoin(10), balance)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(999)

		assert.Equal(t, startingBalance, wallet.Balance())
		assert.NotNil(t, err)
		assert.True(t, errors.Is(err, ErrInsufficientFunds))
	})
}
