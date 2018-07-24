package loyalty_test

import (
	"testing"
	. "loyalty"
)
func Test_ConnectDatabase_Should_Be_Not_Nil(t *testing.T) {
	actual := ConnectDatabase()
	if actual == nil {
		t.Errorf("expect 'Not nil' but got 'Nil'")
	}

}
