package HashTable

import (
	"fmt"
	"hash/fnv"
)

// KeyValuePair 用於存儲哈希表中的鍵值對
type KeyValuePair[K comparable, V any] struct {
	Key   K
	Value V
}

// Hashtable 結構定義，使用泛型支持任意類型的鍵和值
type HashTable[K comparable, V any] struct {
	buckets [][]KeyValuePair[K, V]
	size    int
}

// NewHashtable 創建一個新的哈希表實例，預設大小為
func NewHashtable[K comparable, V any](size int) *HashTable[K, V] {
	return &HashTable[K, V]{
		buckets: make([][]KeyValuePair[K, V], size),
		size:    size,
	}
}

// 哈希函数计算：hashFunction(key) 根据传入的键 key 计算出一个 32位的哈希值。
// 哈希函数的目的是将任意长度的输入（在这里是键 key）映射到一个固定大小的整数值（哈希值）。
// FNV 算法是实现这一映射的方式之一。
func hashFunction[K comparable](key K) uint32 {
	h := fnv.New32a()
	h.Write([]byte(fmt.Sprintf("%v", key)))
	return h.Sum32()
}

func (ht *HashTable[K, V]) Set(key K, value V) {
	//取模运算：计算出的哈希值然后通过取模运算 hashFunction(key) % uint32(ht.size) 与哈希表的桶数量进行运算。这个过程的目的是将哈希值转换成一个在哈希表桶索引范围内的数值。取模运算确保了这个结果是一个介于 0 到 ht.size - 1（包含）之间的整数，这样就可以用它作为索引来访问哈希表的特定桶了。
	//例如，如果哈希值是 1024，并且哈希表的大小（ht.size）是 16，那么 1024 % 16 = 0。这意味着键 key 对应的值应该存储在索引为 0 的桶中。
	//目的：这种方法的主要目的是均匀分布所有可能的键到哈希表的不同桶中。理想情况下，这可以最小化哈希碰撞（不同的键产生相同的哈希值或索引），从而优化哈希表的性能，使得数据的插入、查找和删除操作尽可能接近常数时间复杂度（O(1)）。
	index := hashFunction(key) % uint32(ht.size)
	for i, pair := range ht.buckets[index] {
		if pair.Key == key {
			ht.buckets[index][i].Value = value
			return
		}
	}
	ht.buckets[index] = append(ht.buckets[index], KeyValuePair[K, V]{Key: key, Value: value})
}

func (ht *HashTable[K, V]) Get(key K) (V, bool) {
	index := hashFunction(key) % uint32(ht.size)
	for _, pair := range ht.buckets[index] {
		if pair.Key == key {
			return pair.Value, true
		}
	}
	var zeroValue V
	return zeroValue, false
}
