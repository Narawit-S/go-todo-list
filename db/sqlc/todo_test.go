package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/Narawit-S/go-todo-list/utils"
	"github.com/stretchr/testify/require"
)

func createTestTodo(t *testing.T, arg CreateTodoParams) Todo{
	todo, err := testQuries.CreateTodo(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, todo)

	require.Equal(t, arg.UserID, todo.UserID)
	require.Equal(t, arg.Title, todo.Title)
	require.Equal(t, arg.Finished, todo.Finished)

	require.NotEmpty(t, todo.ID)
	require.NotEmpty(t, todo.CreatedAt)
	require.NotEmpty(t, todo.UpdatedAt)

	return todo
}

func TestCreateDeadlineTodo(t *testing.T)()  {
	user := createTestUser(t)

	arg := CreateTodoParams{
		UserID: user.ID,
		Title: utils.RandomString(10),
		Finished: false,
		Deadline: sql.NullTime{
			Time: utils.DateFromNow(7),
			Valid: true,
		},
	}

	todo := createTestTodo(t, arg)

	require.WithinDuration(t, arg.Deadline.Time, todo.Deadline.Time, time.Second)
}

func TestCreateNoDeadlineTodo(t *testing.T) {
	user := createTestUser(t)

	arg := CreateTodoParams{
		UserID: user.ID,
		Title: utils.RandomString(10),
		Finished: false,
	}

	todo := createTestTodo(t, arg)

	require.Equal(t, false, todo.Deadline.Valid)
}

func TestListTodos(t *testing.T) {
	user := createTestUser(t)

	for i := 0; i < 2; i++ {
		createTestTodo(t, CreateTodoParams{
			UserID: user.ID,
			Title: utils.RandomString(10),
			Finished: false,
			Deadline: sql.NullTime{
				Time: utils.DateFromNow(i + 1),
				Valid: true,
			},
		})
	}

	todo_list, err := testQuries.ListTodos(context.Background(), user.ID)

	require.NoError(t, err)
	require.Len(t, todo_list, 2)

	for _, value := range todo_list {
		require.NotEmpty(t, value)
	}

}

func TestEmptyTodos(t *testing.T) {
	user1 := createTestUser(t)
	user2 := createTestUser(t)

	createTestTodo(t, CreateTodoParams{
		UserID: user1.ID,
		Title: utils.RandomString(10),
		Finished: false,
	})

	todo_list, err := testQuries.ListTodos(context.Background(), user2.ID)

	require.NoError(t, err)
	require.Len(t, todo_list, 0)
}

func TestUpdateTodo(t *testing.T) {
	user := createTestUser(t)

	old_todo := createTestTodo(t, CreateTodoParams{
		UserID: user.ID,
		Title: utils.RandomString(8),
		Finished: false,
	})

	update_arg := UpdateTodoParams{
		ID: old_todo.ID,
		Title: utils.RandomString(10),
		Finished: true,
	}

	update_todo, err := testQuries.UpdateTodo(context.Background(), update_arg)

	require.NoError(t, err)
	require.NotEmpty(t, update_todo)

	require.Equal(t, old_todo.ID, update_todo.ID)
	require.Equal(t, update_arg.Title, update_todo.Title)
	require.Equal(t, update_arg.Finished, update_todo.Finished)
}

func TestDeleteTodo(t *testing.T) {
	user := createTestUser(t)

	todo := createTestTodo(t, CreateTodoParams{
		UserID: user.ID,
		Title: utils.RandomString(4),
		Finished: false,
	})

	err := testQuries.DeleteTodo(context.Background(), todo.ID)

	require.NoError(t, err)
}
