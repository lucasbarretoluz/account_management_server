package db

import (
	"context"
	"testing"

	"github.com/lucasbarretoluz/accountmanagment/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		UserName:       util.RandomUserName(),
		HashedPassword: "secret",
		FullName:       util.RandomUserName(),
		Email:          util.RandomUserEmail(),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserName, user.UserName)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.NotZero(t, user.CreateAt)

	return user
}

func TestCreateUserProfile(t *testing.T) {
	createRandomUser(t)
}

func TestGetUserProfile(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.UserName)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserName, user2.UserName)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.CreateAt, user2.CreateAt)
}
