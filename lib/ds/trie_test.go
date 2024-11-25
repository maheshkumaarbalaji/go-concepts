package ds

import (
	"testing"
)

func Test_RouteParts(t *testing.T) {
	testCases := []struct {
		Name string
		RoutePath string
		ExpArgCount int
	} {
		{ "Normal Route Path", "/abc/nbsdf/123", 3 },
		{ "Route Path with a suffix", "ahf/dggdg/sdfgdfg/", 3 },
		{ "Route Path with two slash in the middle", "/afsf/bfsdf//nfsdnf", 3 }, 
		{ "Route Path with no prefix", "abc/ert/123", 3 },
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(tt *testing.T) {
			routeParts := GetRouteParts(testCase.RoutePath)
			if len(routeParts) != testCase.ExpArgCount {
				tt.Fatalf("Number of route parts (%d) does not match the expected route part count (%d)", len(routeParts), testCase.ExpArgCount)
			} else {
				tt.Logf("Number of route parts (%d) matches the expected route part count (%d)", len(routeParts), testCase.ExpArgCount)
			}
		})
	}
}