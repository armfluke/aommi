package loyalty_test

import (
	"testing"
	. "loyalty"
)
func Test_GetHistoryFromDatabase_Should_ListHistory(t *testing.T) {
	Account := "1100400758552"
	actual := GetHistoryFromDatabase(Account)
	expected := `[{rewardcode:"xxxyyy",date:"2018-07-24",promotionname:"แลกตั๋วฟรี",pomotionpoint:100},{rewardcode:"xxxyyy",date:"2018-07-24",promotionname:"แลกตั๋วฟรี",pomotionpoint:100},{rewardcode:"xxxyyy",date:"2018-07-24",promotionname:"แลกตั๋วฟรี",pomotionpoint:100}]`
	
	if actual != expected {
		t.Errorf("expect '%s' but got '%s'",expected,actual)
	}
}
