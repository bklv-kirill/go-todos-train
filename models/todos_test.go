package models

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type testCase struct {
	todo     Todo
	expected string
}
type testCases []testCase

func TestColoredTitle(t *testing.T) {
	var tcs testCases = testCases{
		testCase{
			todo: Todo{
				Title:     "todo",
				CreatedAt: time.Now(),
			},
			expected: "\u001B[32mtodo\u001B[0m",
		},
		testCase{
			todo: Todo{
				Title:     "todo",
				CreatedAt: time.Now().AddDate(0, 0, -1),
			},
			expected: "\u001B[32mtodo\u001B[0m",
		},
		testCase{
			todo: Todo{
				Title:     "todo",
				CreatedAt: time.Now().AddDate(0, 0, -2),
			},
			expected: "\u001B[33mtodo\u001B[0m",
		},
		testCase{
			todo: Todo{
				Title:     "todo",
				CreatedAt: time.Now().AddDate(0, 0, -3),
			},
			expected: "\u001B[33mtodo\u001B[0m",
		},
		testCase{
			todo: Todo{
				Title:     "todo",
				CreatedAt: time.Now().AddDate(0, 0, -4),
			},
			expected: "\u001B[31mtodo\u001B[0m",
		},
		testCase{
			todo: Todo{
				Title:     "todo",
				CreatedAt: time.Now().AddDate(0, 0, -7),
			},
			expected: "\u001B[31mtodo\u001B[0m",
		},
		testCase{
			todo: Todo{
				Title:      "todo",
				CreatedAt:  time.Now().AddDate(0, 0, -7),
				IsComplete: true,
			},
			expected: "\u001B[32mtodo\u001B[0m",
		},
	}

	for _, tc := range tcs {
		require.Equal(t, tc.expected, tc.todo.ColoredTitle())
	}
}

func TestTodo_ColoredStatus(t *testing.T) {
	var tcs testCases = testCases{
		testCase{
			todo: Todo{
				IsComplete: true,
			},
			expected: "\u001B[32mDONE\u001B[0m",
		},
		testCase{
			todo: Todo{
				IsComplete: false,
			},
			expected: "\u001B[31mNOT DONE\u001B[0m",
		},
	}

	for _, tc := range tcs {
		require.Equal(t, tc.expected, tc.todo.ColoredStatus())
	}
}
