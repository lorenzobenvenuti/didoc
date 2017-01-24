package main

import (
	"errors"
	"fmt"
)

type DocType int

const (
	NONE DocType = iota
	TEXT
	MARKDOWN
	HTML
)

func GetDocType(dt string) (DocType, error) {
	if "txt" == dt {
		return TEXT, nil
	} else if "html" == dt {
		return HTML, nil
	} else if "md" == dt {
		return MARKDOWN, nil
	}
	return NONE, errors.New(fmt.Sprintf("Invalid document type %s", dt))
}
