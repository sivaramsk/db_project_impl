package db0102

import (
	"encoding/binary"
	"io"
)

type Entry struct {
	key []byte
	val []byte
}

func (ent *Entry) Encode() []byte {
	data := make([]byte, 4+4+len(ent.key)+len(ent.val))
	binary.LittleEndian.PutUint32(data[0:4], uint32(len(ent.key)))
	binary.LittleEndian.PutUint32(data[4:8], uint32(len(ent.val)))
	copy(data[8:], ent.key)
	copy(data[8+len(ent.key):], ent.val)
	return data
}

func (ent *Entry) Decode(r io.Reader) error {
	// Allocate 8 bytes for the two uint64 headers
	var header [8]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return err
	}

	// Read using LittleEndian.Uint64 (Requires 8 bytes per slice)
	kLen := int(binary.LittleEndian.Uint32(header[:4]))
	vLen := int(binary.LittleEndian.Uint32(header[4:8]))

	// Read content payloads
	payloadBuf := make([]byte, kLen+vLen)
	_, err := io.ReadFull(r, payloadBuf)
	if err != nil {
		return err
	}

	ent.key = payloadBuf[:kLen]
	ent.val = payloadBuf[kLen:]

	return nil

}

// QzBQWVJJOUhU https://trialofcode.org/
