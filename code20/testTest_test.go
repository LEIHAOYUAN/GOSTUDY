package code20

import "testing"

/*
【功能函数测试】
主要作用检测程序逻辑的正确性
*/
func TestFb1(t *testing.T) {
	if fb1(0) != 0 {
		t.Error(`fb1(0)!=0`)
	}
	if fb1(1) != 1 {
		t.Error(`fb1(1)!=1`)
	}
	if fb1(2) != 1 {
		t.Error(`fb1(2)!=1`)
	}
	if fb1(10) != 55 {
		t.Error(`fb1(10)!=55`)
	}
}

func TestFb2(t *testing.T) {
	if fb2(0) != 0 {
		t.Error(`fb2(0)!=0`)
	}
	if fb2(1) != 1 {
		t.Error(`fb2(1)!=1`)
	}
	if fb2(2) != 1 {
		t.Error(`fb2(2)!=1`)
	}
	if fb2(10) != 55 {
		t.Error(`fb2(10)!=55`)
	}
}
