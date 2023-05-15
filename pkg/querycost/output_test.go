package querycost

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	profile := "kurumaebi"
	start := "bom"
	end := ""
	var resultsCosts = [][]string{}
	var output = NewArrayOutput(&resultsCosts)

	NewQuery(profile, start, end, "MONTHLY", "LINKED_ACCOUNT", "", "UnblendedCost", output, "", "custom")
	fmt.Printf("%v\n", output.Array)
}
