package main

import (
	"testing"
	lib "local/pkg"
)

func TestGroupStrCase1(t *testing.T) {
	actual := lib.GroupStr("aaabbbccccc")
	expect := "a3b3c5"
	if actual != expect {
		t.Errorf("%s != %s", actual, expect)
	}
}

func TestGroupStrCase2(t *testing.T) {
	actual := lib.GroupStr("zzzzcccUUUuu")
	expect := "U3c3u2z4"
	if actual != expect {
		t.Errorf("%s != %s", actual, expect)
	}
}

func TestGroupStrCase3(t *testing.T) {
	actual := lib.GroupStr("ЯЯЯБББддд")
	expect := "Б3Я3д3"
	if actual != expect {
		t.Errorf("%s != %s", actual, expect)
	}
}

func TestGroupStrCase4(t *testing.T) {
	actual := lib.GroupStr("aaabbbcccccaaaaa")
	expect := "a8b3c5"
	if actual != expect {
		t.Errorf("%s != %s", actual, expect)
	}
}

func TestGroupStrCase5(t *testing.T) {
	actual := lib.GroupStr("aaabbbccccc")
	expect := "a3b3c5"
	if actual != expect {
		t.Errorf("%s != %s", actual, expect)
	}
}

func TestGroupStrCase6(t *testing.T) {
	actual := lib.GroupStr("abbbccccc")
	expect := "ab3c5"
	if actual != expect {
		t.Errorf("%s != %s", actual, expect)
	}
}

func TestGroupStrCase7(t *testing.T) {
	actual := lib.GroupStr("a")
	expect := "a"
	if actual != expect {
		t.Errorf("%s != %s", actual, expect)
	}
}

func TestGroupStrCase8(t *testing.T) {
	actual := lib.GroupStr("aa")
	expect := "a2"
	if actual != expect {
		t.Errorf("%s != %s", actual, expect)
	}
}


func TestGroupStrCase9(t *testing.T) {
	actual := lib.GroupStr("")
	expect := ""
	if actual != expect {
		t.Errorf("%s != %s", actual, expect)
	}
}