func maxProfit(prices []int) int {
  l:=0
  r:=1
  max:=0
  for r<len(prices){
	if prices[r]<prices[l]{
		l=r
	}
	if diff:=prices[r]-prices[l]; diff>max{
		max=diff
	}
	r++
  }
  return max
}
///l=0 r=1 