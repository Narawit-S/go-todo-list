package db

import (
	"context"
	"testing"
	"time"

	"github.com/Narawit-S/go-todo-list/utils"
	"github.com/stretchr/testify/require"
)

func createTestUser(t *testing.T) User {
	arg := CreateUserParams{
		Email: utils.RandomEmail(),
		EncryptedPassword: utils.RandomPassword(),
	}

	user, err := testQuries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.EncryptedPassword, user.EncryptedPassword)

	require.NotEmpty(t, user.ID)
	require.NotEmpty(t, user.CreatedAt)
	require.NotEmpty(t, user.UpdatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createTestUser(t)
}

func TestGetUser(t *testing.T)() {
	user := createTestUser(t)

	query_user, err := testQuries.GetUser(context.Background(), user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, query_user)

	require.Equal(t, user.ID, query_user.ID)
	require.Equal(t, user.Email, query_user.Email)
	require.Equal(t, user.EncryptedPassword, query_user.EncryptedPassword)
	require.WithinDuration(t, user.CreatedAt, query_user.CreatedAt, time.Second)
	require.WithinDuration(t, user.UpdatedAt, query_user.UpdatedAt, time.Second)
}

func TestUpdateUser(t *testing.T) {
	user := createTestUser(t)

	arg := UpdateUserParams{
		ID: user.ID,
		EncryptedPassword: utils.RandomPassword(),
	}

	time.Sleep(time.Second)

	update_user, err := testQuries.UpdateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, update_user)

	require.Equal(t, user.ID, update_user.ID)
	require.Equal(t, arg.EncryptedPassword, update_user.EncryptedPassword)
	require.WithinDuration(t, time.Now(), update_user.UpdatedAt, time.Second)
}
