package mailmanmembers

import (
	"fmt"

	"github.com/anaskhan96/soup"
)

// MembersFromString parses a string containing an html document from the
// mailman member list and returns a list of strings containing the mail
// addresses of the member list
func MembersFromString(html string) ([]string, error) {
	result := make([]string, 0)
	root := soup.HTMLParse(html)

	tables := root.FindAll("table")

	if len(tables) < 5 {
		return nil, fmt.Errorf("too few tables! only %d found", len(tables))
	}

	if err := tables[4].Error; err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("error parsing table: %v", err)
	}

	trs := tables[4].FindAll("tr")

	if len(trs) < 3 {
		return nil, fmt.Errorf("too few rows in table! only %d found", len(trs))
	}

	for _, tr := range trs[2:] {
		if err := tr.Error; err != nil {
			return nil, fmt.Errorf("error parsing row: %v", err)
		}

		a := tr.Find("a")
		if err := a.Error; err != nil {
			return nil, fmt.Errorf("error parsing anchor: %v", err)
		}

		result = append(result, a.Text())
	}

	return result, nil
}
