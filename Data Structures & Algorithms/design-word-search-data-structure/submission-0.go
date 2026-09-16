type WordDictionary struct {
    root *Trie
}

type Trie struct {
    children  [26]*Trie
    endofword bool
}

func Constructor() WordDictionary {
    return WordDictionary{root: NewTrie()}
}

func NewTrie() *Trie {
    return &Trie{}
}

func (this *WordDictionary) AddWord(word string) {
    cur := this.root
    for _, c := range word {
        i := c - 'a'
        if cur.children[i] == nil {
            cur.children[i] = NewTrie()
        }
        cur = cur.children[i]
    }
    cur.endofword = true
}

func (this *WordDictionary) Search(word string) bool {
    return this.dfs(word, 0, this.root)
}

func (this *WordDictionary) dfs(word string, j int, cur *Trie) bool {
    for i:=j; i<len(word); i++{
        c:=word[i]
        if c=='.'{
            for _, child:=range cur.children{
                if child!=nil && this.dfs(word, i+1, child){
                    return true 
                }
            }
            return false
        }else{
            ind:=c-'a'
            if cur.children[ind]==nil{
                return false
            }
            cur=cur.children[ind]
        }
    }
    return cur.endofword
}