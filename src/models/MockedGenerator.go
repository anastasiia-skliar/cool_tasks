package models

import "net/url"

func MockedGenerator(req string, vars []interface{}, err error) {
	SQLGenerator = func(findType string, stringArgs []string, numberArgs []string, params url.Values) (string, []interface{}, error) {
		return req, vars, err
	}
}
