type MinStack struct {
    stack[]int
    minstack[]int
}

func Constructor() MinStack {
    return MinStack{
        stack: []int{},
        minstack: []int{},
    }
}

func (this *MinStack) Push(value int)  {
    if len(this.stack)==0{
        this.minstack=append(this.minstack, value)
    }
    this.stack=append(this.stack, value)
    if value<=this.minstack[len(this.minstack)-1]{
        this.minstack=append(this.minstack, value)
    }
}


func (this *MinStack) Pop()  {
    if this.stack[len(this.stack)-1]==this.minstack[len(this.minstack)-1]{
        this.minstack=this.minstack[:len(this.minstack)-1]
    }
    this.stack=this.stack[:len(this.stack)-1]
}


func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
    return this.minstack[len(this.minstack)-1]
}
