package internal

type TrieNode struct {
	children []*TrieNode
	isLeaf   bool
}

type TrieTable struct {
	head *TrieNode
	size int
}

/*
Insert a string with prefix
*/
func (t *TrieTable) Insert(str string) {
	if t.head == nil {
		t.head = &TrieNode{children: make([]*TrieNode, 26)}
	}

	temp := t.head

	for i := 0; i < len(str); i++ {
		key := str[i] - 'a'

		if temp.children[key] == nil {
			temp.children[key] = &TrieNode{children: make([]*TrieNode, 26)}
			temp.isLeaf = false
		}
		temp = temp.children[key]

	}
	temp.isLeaf = true
}

/*

Search if input string has any prefix

*/

func (t *TrieTable) Search(str string) bool {

	temp := t.head
	i := 0
	for i = 0; i < len(str); i++ {
		key := str[i] - 'a'
		if temp.children[key] == nil {
			return false
		}
		temp = temp.children[key]
	}

	if temp.isLeaf != true {
		return false
	}
	return true

}

/*
Search if any word has input prefix
prefix = ab, word = abc
*/
func (t *TrieTable) StartWith(prefix string) bool {
	temp := t.head
	i := 0
	for i = 0; i < len(prefix); i++ {
		key := prefix[i] - 'a'
		if temp.children[key] == nil {
			return false
		}
		temp = temp.children[key]
	}

	return true

}
