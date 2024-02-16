package HashTable

import "fmt"

func RunHashTable() {
	ht := NewHashtable[string, int](16)
	ht.Set("key1", 10)
	ht.Set("key2", 20)

	value, exists := ht.Get("key1")
	if exists {
		fmt.Println("key1:", value)
	} else {
		fmt.Println("key1 does not exist")
	}

	value, exists = ht.Get("key3")
	if exists {
		fmt.Println("key3:", value)
	} else {
		fmt.Println("key3 does not exist")
	}
}
