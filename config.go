package door

// The multiplexer `Door`
// Cobfig
// Copyright © 2017 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

type typeHash uint64

// Types sections URL
const (
	TYPE_STAT = iota
	TYPE_ARG
)

// Editable parameters
const (
	// The method used by default when you add a Route.
	// Example: GET, POST etc.
	HTTP_METHOD_DEFAULT = "GET"

	// Maximum number of sections in the URL
	// Example: /abc/:par - 2 sections, /a/:b/:c - 3 sections
	HTTP_SECTION_COUNT = 16

	// The maximum length of URL (characters)
	HTTP_PATTERN_COUNT = 512
)

// Non-editable parameters
const (
	DELIMITER_COLON byte     = 58
	DELIMITER_SLASH byte     = 47
	DELIMITER_UINT  typeHash = 47
	SLASH_HASH      typeHash = 1
)
