package cli_user_test

import (
	_ "pvectl/cli/command/commands"
	cliTest "pvectl/test/cli"
	"testing"
)

func Test_User_List(t *testing.T) {
	Test := cliTest.Test{
		Expected: `"userid":"root@pam"`,
		ReqErr:   false,
		Contains: true,
		Args:     []string{"-i", "list", "users"},
	}
	Test.StandardTest(t)
}
