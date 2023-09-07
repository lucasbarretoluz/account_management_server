package db

import (
	"context"
	"testing"

	"github.com/lucasbarretoluz/accountmanagment/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransaction(t *testing.T) Transaction {
	user := createRandomUser(t)
	agr := CreateTransactionParams{
		IDUser:      user.IDUser,
		TotalValue:  util.RandomMoney(),
		Category:    util.RandomCategory(),
		Description: util.RandomDescription(),
		IsExpense:   util.RandomIsExpense(),
	}

	transaction, err := testQueries.CreateTransaction(context.Background(), agr)

	require.NoError(t, err)
	require.NotEmpty(t, transaction)

	require.Equal(t, agr.TotalValue, transaction.TotalValue)
	require.Equal(t, agr.Category, transaction.Category)
	require.Equal(t, agr.Description, transaction.Description)
	require.Equal(t, agr.IsExpense, transaction.IsExpense)

	require.NotZero(t, transaction.IDUser)
	require.NotZero(t, transaction.TransactionsAt)

	return transaction
}

func TestCreateTransaction(t *testing.T) {
	createRandomTransaction(t)
}

func TestDeleteTransaction(t *testing.T) {
	transaction1 := createRandomTransaction(t)
	err := testQueries.DeleteTransaction(context.Background(), transaction1.IDTransaction)
	require.NoError(t, err)

	transaction2, err := testQueries.GetTransaction(context.Background(), transaction1.IDTransaction)
	require.Error(t, err)
	require.Empty(t, transaction2)
}

func TestGetTransaction(t *testing.T) {
	transaction1 := createRandomTransaction(t)
	transaction2, err := testQueries.GetTransaction(context.Background(), transaction1.IDTransaction)
	require.NoError(t, err)
	require.NotEmpty(t, transaction2)

	require.Equal(t, transaction1.IDTransaction, transaction2.IDTransaction)
	require.Equal(t, transaction1.TotalValue, transaction2.TotalValue)
	require.Equal(t, transaction1.Category, transaction2.Category)
	require.Equal(t, transaction1.Description, transaction2.Description)
	require.Equal(t, transaction1.IsExpense, transaction2.IsExpense)
}

func TestListTransactions(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomTransaction(t)
	}

	arg := ListTransactionsParams{
		Limit:  5,
		Offset: 5,
	}

	transactions, err := testQueries.ListTransactions(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, transactions, 5)

	for _, transaction := range transactions {
		require.NotEmpty(t, transaction)
	}
}

func TestUpdateTransaction(t *testing.T) {
	transaction1 := createRandomTransaction(t)
	arg := UpdateTransactionParams{
		IDTransaction: transaction1.IDTransaction,
		TotalValue:    util.RandomMoney(),
		Category:      util.RandomCategory(),
		Description:   util.RandomDescription(),
		IsExpense:     util.RandomIsExpense(),
	}
	transaction2, err := testQueries.UpdateTransaction(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transaction2)

	require.Equal(t, transaction1.IDTransaction, transaction2.IDTransaction)
	require.Equal(t, arg.TotalValue, transaction2.TotalValue)
	require.Equal(t, arg.Category, transaction2.Category)
	require.Equal(t, arg.Description, transaction2.Description)
	require.Equal(t, arg.IsExpense, transaction2.IsExpense)
}
