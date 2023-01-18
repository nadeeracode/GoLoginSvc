package db

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/nadeeracode/GoLoginSvc/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	arg := CreateUserParams{
		ID:                uuid.New().String(),
		Username:          util.RandomString(6),
		FirstName:         util.RandomString(6),
		LastName:          util.RandomString(6),
		HashedPassword:    hashedPassword,
		Email:             util.RandomEmail(),
		UserType:          1,
		Active:            1,
		CreatedAt:         time.Now(),
		LastAccessedAt:    time.Now(),
		PasswordChangedAt: time.Now(),
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	return user
}

func TestCreateUser(t *testing.T) {
	fmt.Println(">> before:", "TTTT")
	createRandomUser(t)

}
