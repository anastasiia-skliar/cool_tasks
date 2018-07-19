package models

import (
	"net/url"
	sq "github.com/Masterminds/squirrel"
	"errors"
)

func Contains(array []string, value string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

func SqlGenerator(findType string, stringArgs []string, numberArgs []string, params url.Values) (string, []interface{}, error) {
	req := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("*").From(findType)
	var (
		request string
		err     error
		and     sq.And = nil
	)
	for key, value := range params {
		switch {
		case Contains(stringArgs, key):
			if len(value) > 1 {
				var or sq.Or = nil
				for _, v := range value {
					or = append(or, sq.Eq{key: v})
				}
				and = append(and, or)
			} else {
				and = append(and, sq.Eq{key: value[0]})
			}
		case Contains(numberArgs, key):
			if len(value) == 2 {
				and = append(and, sq.And{sq.GtOrEq{key: value[1]}, sq.LtOrEq{key: value[0]}})
			} else {
				and = append(and, sq.Eq{key: value[0]})

			}
		case key == "id":
			and = append(and, sq.Eq{key: value[0]})
		default:
			return "", nil, errors.New("ERROR: Bad request")
		}
	}
	request, args, err := req.Where(and).ToSql()
	if err != nil {
		return "", nil, err
	}
	return request, args, err
}
