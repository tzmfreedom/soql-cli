package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestParser(t *testing.T) {
	testCases := []struct {
		Code     string
		Expected *Statement
	}{
		{
			`INSERT Account(Name, Field1__c, Field2__c, Field3__c) VALUES ('foo', 1, 1.23, true)`,
			&Statement{
				Type:    "INSERT",
				Sobject: "Account",
				Values: map[string]string{
					"Name":      "foo",
					"Field1__c": "1",
					"Field2__c": "1.23",
					"Field3__c": "true",
				},
			},
		},
		{
			`UPDATE Custom1__c SET Name='foo', Field1__c=1, Field2__c = 1.23, Field3__c=true WHERE id = 1 OR Field1__c != NULL`,
			&Statement{
				Type:    "UPDATE",
				Sobject: "Custom1__c",
				Values: map[string]string{
					"Name":      "foo",
					"Field1__c": "1",
					"Field2__c": "1.23",
					"Field3__c": "true",
				},
				Where: "WHERE id = 1 OR Field1__c != NULL",
			},
		},
		{
			`DELETE FROM Contact WHERE id = 1 AND Field1__c < 2`,
			&Statement{
				Type:    "DELETE",
				Sobject: "Contact",
				Where:   "WHERE id = 1 AND Field1__c < 2",
			},
		},
	}

	for _, testCase := range testCases {
		actual := ParseString(testCase.Code)
		if !cmp.Equal(testCase.Expected, actual) {
			fmt.Println(cmp.Diff(testCase.Expected, actual))
			t.Fatalf("expected %s, actual %s", testCase.Expected, actual)
		}
	}
}
