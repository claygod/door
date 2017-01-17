package door

// The multiplexer `Door`
// Section
// Copyright © 2017 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

// Each route includes a section
type section struct {
	id      []byte
	typeSec int // 0 - TYPE_STAT, 1 - TYPE_ARG
}

func newSection(sec []byte, tps int) *section {
	return &section{id: sec, typeSec: tps}
}
