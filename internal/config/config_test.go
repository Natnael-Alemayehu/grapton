package config

import (
	"testing"
)

func TestRead(t *testing.T) {
	usecases := []struct {
		name          string
		expectedDBURL string
	}{
		{
			name:          "read default",
			expectedDBURL: "postgres://example",
		},
	}

	for i, tc := range usecases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := Read()
			if err != nil {
				t.Errorf("TEST: %d - NAME: %v - ERROR: %v", i, tc.name, err)
				return
			}

			if actual.DBURL != tc.expectedDBURL {
				t.Errorf("TEST: %d - NAME: %v - GOT: %v , EXPECTED: %v", i, tc.name, actual.DBURL, tc.expectedDBURL)
			}
		})
	}

}

func TestSetUser(t *testing.T) {
	usecases := []struct {
		name             string
		expectedDBURL    string
		expectedUserName string
	}{
		{
			name:             "read default",
			expectedDBURL:    "postgres://example",
			expectedUserName: "changed nate",
		},
	}

	for i, tc := range usecases {
		t.Run(tc.name, func(t *testing.T) {
			cfg := Config{
				DBURL:           "postgres://example",
				CurrentUserName: "nate",
			}
			actual, err := cfg.SetUser("changed nate")
			if err != nil {
				t.Errorf("TEST: %d - NAME: %v - ERROR: %v", i, tc.name, err)
				return
			}

			if actual.DBURL != tc.expectedDBURL {
				t.Errorf("TEST: %d - NAME: %v - GOT: %v , EXPECTED: %v", i, tc.name, actual.DBURL, tc.expectedDBURL)
			}

			if actual.CurrentUserName != tc.expectedUserName {
				t.Errorf("TEST: %d - NAME: %v - GOT: %v , EXPECTED: %v", i, tc.name, actual.CurrentUserName, tc.expectedUserName)
			}

		})
	}

}

func TestGetConfingFilePath(t *testing.T) {
	usecases := []struct {
		name     string
		filepath string
	}{
		{
			name:     "read default",
			filepath: "/home/nate/.gatorconfig.json",
		},
	}

	for i, tc := range usecases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getConfigFilePath()
			if err != nil {
				t.Errorf("TEST: %d - NAME: %v - ERROR: %v", i, tc.name, err)
				return
			}

			if tc.filepath != actual {
				t.Errorf("TEST: %d - NAME: %v - GOT: %v , EXPECTED: %v", i, tc.name, actual, tc.filepath)
				return
			}
		})
	}

}
