func searchMatrix(matrix [][]int, target int) bool {
	l:=0
	collen:=len(matrix[0])-1
	r:=len(matrix)-1
	mid:=(r+l)/2
	for l<=r{
		if target>=matrix[mid][0] && target<=matrix[mid][collen]{
			nl:=0
			nr:=collen
			m:=(nl+nr)/2
			for nl<=nr{
				if matrix[mid][m]==target{
					return true 
				}else if target>matrix[mid][m]{
					nl=m+1
				}else{
					nr=m-1
				}
				m=(nr+nl)/2
			}
			return false 
		}else if target>matrix[mid][collen]{
			l=mid+1
		}else if target<matrix[mid][0]{
			r=mid-1
		}else{
			return false 
		}
		mid=(l+r)/2
	}
	return false 

}
