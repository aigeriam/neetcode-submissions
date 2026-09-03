

func evalRPN(tokens []string) int {
    stack:=[]int{}
    res:=0

    for _, token:=range tokens{
        num, err:=strconv.Atoi(token)
        if err==nil{
            stack=append(stack, num)
        }
        if len(tokens)==1{
            return num
        }
        if token=="+"{
            res=stack[len(stack)-1]+stack[len(stack)-2]
            stack=stack[:len(stack)-2]
            stack=append(stack, res)
        }
        if token=="/"{
           res=stack[len(stack)-2]/stack[len(stack)-1]
           stack=stack[:len(stack)-2]
           stack=append(stack, res)
        }
        if token=="*"{
            res=stack[len(stack)-1]*stack[len(stack)-2]
            stack=stack[:len(stack)-2]
            stack=append(stack, res)
        }
        if token=="-"{
            res=stack[len(stack)-2]-stack[len(stack)-1]
            stack=stack[:len(stack)-2]
            stack=append(stack, res)
        }




    }
    return res
}



