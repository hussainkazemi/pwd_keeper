package utility

import "testing"

type testCase struct {
	input  string
	output string
}

func TestLoadFromEnv(t *testing.T) {
	tCases := []testCase{
		{
			input:  "ENCRYPT_KEY",
			output: "your_encrypt_key",
		},
		{
			input:  "DB_USER",
			output: "your_database_user",
		},
		{
			input:  "DB_PASSWORD",
			output: "your_database_password",
		},
		{
			input:  "DB_NAME",
			output: "your_database_name",
		},
		{
			input:  "vojod_nadarad",
			output: "",
		},
	}

	for _, v := range tCases {
		o, err := LoadFromEnv(v.input)
		if err != nil {
			t.Errorf("can not load from evn error: %s", err.Error())
		}
		if o != v.output {
			t.Errorf("for %s wanted %s but get %s", v.input, v.output, o)
		}
	}

}
