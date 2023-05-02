package token

import (
	"testing"
	"time"

	"github.com/Narawit-S/go-todo-list/utils"
	"github.com/stretchr/testify/require"
)

func TestPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	email := utils.RandomEmail()

	token, err := maker.CreateToken(email, time.Hour)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.ValidateToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, email, payload.Email)
	require.NotZero(t, payload.ID)
	require.WithinDuration(t, time.Now(), payload.IssuedAt, time.Second)
	require.WithinDuration(t, time.Now().Add(time.Hour), payload.ExpiredAt, time.Second)
}

func TestPastoExpiredToken(t *testing.T) {
	maker, err := NewPasetoMaker(utils.RandomString(32))
	require.NoError(t, err)

	email := utils.RandomEmail()

	token, err := maker.CreateToken(email, -time.Hour)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.ValidateToken(token)
	require.Error(t, err)
	require.Nil(t, payload)
}
