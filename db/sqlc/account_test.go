package db

import (
	"context"
	"database/sql"
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

	require.Equal(t, acc1.ID, account.ID)
	require.Equal(t, acc1.Balance, account.Balance)
	require.Equal(t, acc1.Currency, account.Currency)

	require.NotZero(t, account.CreatedAt)
}

func TestUpdateAccount(t *testing.T) {
	acc1 := createRamdomAccount(t)

	arg := UpdateAccountBalanceParams{
		ID:      acc1.ID,
		Balance: acc1.Balance,
	}

	account2, err := testQueries.UpdateAccountBalance(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, acc1.ID, account2.ID)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, acc1.Currency, account2.Currency)

	require.NotZero(t, account2.CreatedAt)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRamdomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRamdomAccount(t)
	}

	arg := ListAccountParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccount(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, accounts, 5)
}
