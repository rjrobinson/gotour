package trie

type (
	Trie     struct{} // building a trie struct
	Iterator struct{} // building an iterator struct
)

// a new trie
func New() *Trie {
	return new(Trie)
}

func (t *Trie) Get(key []byte) interface{} {
	return nil
}

func (t *Trie) Set(key []byte, value interface{}) bool {
	return false
}

func (t *Trie) Delete(key []byte) bool {
	return false
}
