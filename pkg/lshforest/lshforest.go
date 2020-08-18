package lshforest

import "github.com/justinfargnoli/lshforest/pkg/lshtree"

// LSHForest is an index of high-dimensional data based on cosine similarity 
type LSHForest struct {
	trees []lshtree.Trie
}

// NewLSHForestDefault constructs an LSHForest struct with sensible defaults
func NewLSHForestDefault(dim uint) *LSHForest {
	return NewLSHForest(5, 20, dim)
}

// NewLSHForest constructs an LSHForest. l := the number of trees in the forest 
// of LSHForest. maxK := the maximum number of hash functions. The larger maxK 
// is, the more accurate LSHForest is and the more space LSHForest takes up.
// dim := the dimension of the input vectors
func NewLSHForest(l, maxK, dim uint) *LSHForest {
	var trees []lshtree.Trie
	for i := uint(0); i < l; i++ {
		trie := lshtree.NewTrie()
		trees = append(trees, trie)
	}
	return &LSHForest{trees: trees}
}

// // Insert puts the vector into the LSHForest
// func (f *LSHForest) Insert(vector []float64, value interface{}) {
	
// }