package main

import (
	"testing"
)

func TestTopN(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		n       int
		wantLen int
	}{
		{"empty input", []int{}, 0, 0},
		{"negative input", []int{3, 4}, -2, 0},
		{"normal case ", []int{30, 40, 50}, 2, 2},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TopN(tc.input, tc.n)
			if len(got) != tc.wantLen {
				t.Errorf("got len %d, want %d", len(got), tc.wantLen)
			}
		})
	}
}

func TestCorrectOutput(t *testing.T) {
	tc := struct {
		name string
		inp  []int
		n    int
		out  []int
	}{
		"validates correct TopN", []int{29, 1, 3, 40}, 2, []int{40, 29},
	}

	t.Run(tc.name, func(t *testing.T) {
		got := TopN(tc.inp, tc.n)
		for i, v := range tc.out {
			if got[i] != v {
				t.Errorf("Output arrays do not match expect %d but got %d | []exp %v | []got %v ", v, got[i], tc.out, got)
			}
		}
	})
}

func TestAccount(t *testing.T) {

	test_account := Account{
		balance: 500,
		account_owner: Customer{
			name:     "Michail",
			lastname: "Bredichin",
			phone:    812384884,
		},
		account_number: "ABNA134124566132",
	}

	tests := []struct {
		name   string
		amount float64
	}{
		{"valid deposit", 100},
		{"invalid deposit", -100},
		{"withdraw", 100},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := test_account.Deposit(tc.amount)
			if err == nil {
				if test_account.balance != tc.amount {
					t.Errorf("balance is not correct, want %f but got %f", tc.amount, test_account.balance)
				}
			} else {
				t.Errorf("error is not nil, want %v but got %v", nil, err)
			}
		})
	}
}
