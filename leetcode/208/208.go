type Trie struct {
    nodes   []*Trie
    end     bool
}

/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{
        nodes: make([]*Trie, 26),
        end: false,
    }
}

func (this *Trie) InsertWord(word string, index int) *Trie {
    if this == nil {
        this = & Trie{
            nodes: make([]*Trie, 26),
            end: false,
        }
    }

    if index < len(word) {
        c := word[index] - 'a'
        this.nodes[c] = this.nodes[c].InsertWord(word, index + 1)
    } else {
        this.end = true
    }

    return this
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    this.InsertWord(word, 0)
}

func (this *Trie) SearchWord(word string, index int) *Trie {
    if this == nil {
        return nil
    }
    if index < len(word) {
        c := word[index] - 'a'
        return this.nodes[c].SearchWord(word, index + 1)
    } else {
        return this
    }
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    node := this.SearchWord(word, 0)
    if node == nil {
        return false
    } 
    return node.end
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    node := this.SearchWord(prefix, 0)
    if node == nil {
        return false
    } 
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
