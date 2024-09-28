package go_todos_train

import (
	"github.com/bklv-kirill/go-todos-train/storage/sqlite3"
	"os"
	"testing"
)

type testCase struct {
	args      []string
	haveError bool
}
type testCases []testCase

var ms *sqlite3.MockStorage = &sqlite3.MockStorage{}
var mbs *sqlite3.MockBrokenStorage = &sqlite3.MockBrokenStorage{}

var listTcs testCases = testCases{
	testCase{
		args:      []string{"command"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "ls"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--ls"},
		haveError: false,
	},
}

func TestList(t *testing.T) {
	var commands *Commands = NewCommands(ms)

	for _, tc := range listTcs {
		os.Args = tc.args

		var err error = commands.Execute()
		if err != nil && !tc.haveError {
			t.Error()
		} else if err == nil && tc.haveError {
			t.Error()
		}
	}
}

func TestListWithBrokenStorage(t *testing.T) {
	var commands *Commands = NewCommands(mbs)

	for _, tc := range listTcs {
		os.Args = tc.args
		if err := commands.Execute(); err == nil {
			t.Error()
		}
	}
}

var addTcs testCases = testCases{
	testCase{
		args:      []string{"command", "--add"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--add", ""},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--add", "error"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--add", "add"},
		haveError: false,
	},
}

func TestAdd(t *testing.T) {
	var commands *Commands = NewCommands(ms)

	for _, tc := range addTcs {
		os.Args = tc.args

		var err error = commands.Execute()
		if err != nil && !tc.haveError {
			t.Error()
		} else if err == nil && tc.haveError {
			t.Error()
		}
	}
}

var rmdTcs testCases = testCases{
	testCase{
		args:      []string{"command", "--rm"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--rm", ""},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--rm", "rm"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--rm", "999"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--rm", "1"},
		haveError: false,
	},
}

func TestRm(t *testing.T) {
	var commands *Commands = NewCommands(ms)

	for _, tc := range rmdTcs {
		os.Args = tc.args

		var err error = commands.Execute()
		if err != nil && !tc.haveError {
			t.Error()
		} else if err == nil && tc.haveError {
			t.Error()
		}
	}
}

var csTcs testCases = testCases{
	testCase{
		args:      []string{"command", "cs"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--cs"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--cs", ""},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--cs", "cs"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--cs", "999"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--cs", "1"},
		haveError: false,
	},
}

func TestCs(t *testing.T) {
	var commands *Commands = NewCommands(ms)

	for _, tc := range csTcs {
		os.Args = tc.args

		var err error = commands.Execute()
		if err != nil && !tc.haveError {
			t.Error()
		} else if err == nil && tc.haveError {
			t.Error()
		}
	}
}

var ctTcs testCases = testCases{
	testCase{
		args:      []string{"command", "ct"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--ct"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--ct", ""},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--ct", "ct", ""},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--ct", "999"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--ct", "1", ""},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--ct", "999", "ct"},
		haveError: true,
	},
	testCase{
		args:      []string{"command", "--ct", "1", "ct"},
		haveError: false,
	},
}

func TestCt(t *testing.T) {
	var commands *Commands = NewCommands(ms)

	for _, tc := range ctTcs {
		os.Args = tc.args

		var err error = commands.Execute()
		if err != nil && !tc.haveError {
			t.Error()
		} else if err == nil && tc.haveError {
			t.Error()
		}
	}
}

var helpTcs testCases = testCases{
	testCase{
		args:      []string{"command", "--help"},
		haveError: false,
	},
}

func TestHelp(t *testing.T) {
	var commands *Commands = NewCommands(ms)

	for _, tc := range helpTcs {
		os.Args = tc.args

		var err error = commands.Execute()
		if err != nil && !tc.haveError {
			t.Error()
		} else if err == nil && tc.haveError {
			t.Error()
		}
	}
}
