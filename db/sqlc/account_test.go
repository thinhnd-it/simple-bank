package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thinhnd-it/simplebank/utils"
)

func createRamdomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Balance, arg.Balance)
	require.Equal(t, account.Owner, arg.Owner)
	require.Equal(t, account.Currency, arg.Currency)

	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.ID)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRamdomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc1 := createRamdomAccount(t)

	account, err := testQueries.GetAccount(context.Background(), acc1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account)
}
