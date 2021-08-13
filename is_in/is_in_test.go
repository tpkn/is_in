package is_in

import (
	"fmt"
	"testing"
	
	"is_in/models"
)

type options struct {
	string1    string
	string2    string
	prepare    bool
	ignorecase bool
	result     string
}

var cli = models.CLI{
	Delimiter: ",",
	IgnoreCase: true,
	Prepare:   false,
}

var checklist = []options{
	{ "A", "'a' ,b", true, true, "true" },
	{ "A,B,C", "A,B,C", false, false, "true" },
	{ "A,B,C", "C,A,B", false, false, "true" },
	{ "A", "C,A,B", false, false, "true" },
	{ "A", "C, A ,B", false, false, "false" },
	{ "A", "C, 'A' ,B", true, false, "true" },
	{ "A", `"""  A ''''" , B`, true, false, "true" },
	{ "A", "", false, false, "false" },
	{ "", "A", false, false, "false" },
	{ "", "", false, false, "false" },
}

func Test_List(t *testing.T) {
	for _, test := range checklist {
		result, _ := Run(test.string1, test.string2, models.CLI{
			Delimiter:  ",",
			Prepare:    test.prepare,
			IgnoreCase: test.ignorecase,
			Verbose:    false,
		})
		
		if result != test.result {
			if result == "true" {
				t.Error(fmt.Sprintf("%q is not in %q", test.string1, test.string2))
			} else {
				t.Error(fmt.Sprintf("%q is in %q", test.string1, test.string2))
			}
		} else {
			prepared := ""
			if test.prepare {
				prepared = "(prepared)"
			}
			ignorecase := ""
			if test.ignorecase {
				ignorecase = "(ignore case)"
			}
			if result == "true" {
				fmt.Println(fmt.Sprintf("✓ %q is in %q %v %v", test.string1, test.string2, prepared, ignorecase))
			} else {
				fmt.Println(fmt.Sprintf("✓ %q is not in %q %v %v", test.string1, test.string2, prepared, ignorecase))
			}
		}
	}
}

func Benchmark_IsIn(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Run("a,g, h,t, u, ", "a, b, c, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u, v, w, x, y, z,", cli)
	}
}