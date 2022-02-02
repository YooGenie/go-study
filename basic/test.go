package basic

import "fmt"

func Solution()  {
	//var grade = []int{2,1,1,3,5,2}
	//result :=0
	//
	//for i:=len(grade); i>1; i-- {
	//	if grade[i-1] < grade[i-2] {
	//
	//		fmt.Println("zz : ",grade[i-2])
	//		result+= grade[i-2] - grade[i-1]
	//		grade[i-2] = grade[i-1]
	//	}
	//}
	//fmt.Println(result)
	//return result

	n := 4
	horizontal := true
	var result = make([][]int, n,n)
	count :=1

	for i:=0 ; i<n; i++ {
		if i==0 {
			result[i] = append(result[i], count)
			count ++
		}
		if  i!=0 && horizontal {
			fmt.Println("i 값 : ", i)
			for j:=0 ; j<i ;j++ {
				fmt.Println("j 값 : ", j)
				result[j] = append(result[j], count)
				count++
			}
			for j:=i ; j>0; j--{
				result[i][j] = count
				count++
			}
			horizontal=false
		}else {

		}

	}
	fmt.Println(result)
}
