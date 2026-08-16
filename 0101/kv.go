package db0101

// package main

import (
	"fmt"
)

type KV struct {
	mem map[string][]byte
}

func (kv *KV) Open() error {
	kv.mem = map[string][]byte{} // empty
	return nil
}

func (kv *KV) Close() error { return nil }

func (kv *KV) Get(key []byte) (val []byte, ok bool, err error) {

	keyStr := string(key)
	if keyStr == "" {
		return nil, false, nil
	}

	sourceValue, exists := kv.mem[keyStr]
	if !exists {
		return nil, false, nil
	}

	// Allocate exact destination memory
	dst := make([]byte, len(sourceValue))

	// copy() returns total elements transferred
	bytesCopied := copy(dst, sourceValue)

	if bytesCopied != len(sourceValue) {
		return nil, false, fmt.Errorf("copy mismatch: expected %d bytes, copied %d", len(sourceValue), bytesCopied)
	}

	fmt.Printf("kv.mem[%s]: %s\n", keyStr, dst)

	return dst, true, nil
}

func (kv *KV) Set(key []byte, val []byte) (updated bool, err error) {
	keyStr := string(key)
	if keyStr == "" {
		return false, nil
	}

	// Allocate exact destination memory
	dst := make([]byte, len(val))
	// copy() returns total elements transferred
	bytesCopied := copy(dst, val)

	if bytesCopied != len(val) {
		return false, fmt.Errorf("copy mismatch: expected %d bytes, copied %d", len(val), bytesCopied)
	}

	kv.mem[keyStr] = dst

	fmt.Printf("kv.mem[%s]: %s\n", keyStr, kv.mem[keyStr])

	return true, nil

}

func (kv *KV) Del(key []byte) (deleted bool, err error) {
	keyStr := string(key)
	if keyStr == "" {
		return false, nil
	}

	_, exists := kv.mem[keyStr]
	if !exists {
		return false, nil
	}

	// 2. Remove the key completely from the map structure
	delete(kv.mem, keyStr)

	return true, nil

}

func main() {
	kv := KV{}
	err := kv.Open()

	if err != nil {
		fmt.Errorf("Error opening kv")
	}
	defer kv.Close()

	updated, err := kv.Set([]byte("k1"), []byte("v1"))
	if err != nil {
		fmt.Errorf("Error setting k1 to kv")
	}

	fmt.Println("Value has been set. Is updated:", updated)

	_, ok, err := kv.Get([]byte("xxx"))
	fmt.Println("Get Value:", ok, "Error: ", err)

}
