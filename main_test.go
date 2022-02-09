package main

func TestIsValidNumber(t *testing.T){
	for _, tc := range []struct {
		testnum int
		num string
		base int
		want bool
	}{
		{1, num: "", base: , want: false},
		{2, num: "123", base: , want: false},
		{3, num: "", base: 10, want: false},
		{4, num: "123", base: 10, want: true},
		{5, num: "123", base: 3, want: false},
		{6, num: "ABC", base: 16, want: true},
		{7, num: "AZ", base: 36, want: true},
		{8, num: "1A", base: 1, want: false},
		{9, num: "12xd", base: 36, want: false}, 
	}{
		if got := isValidNumber(tc.num, tc.base); got!= tc.want{
			t.Errorf("Fail: got = %v, want = %v", got, tc.want)
		}
	}
}
