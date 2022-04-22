package function

import (
	"fmt"
	"sort"
)

func Sort() {
	var nums = []int{3, 5, 6, 1, 0, 66}

	//오름차순
	fmt.Println("정렬하기 전 :", nums)
	sort.Ints(nums)
	fmt.Println("정렬하기 후 :", nums)

	var str = []string{"안녕", "지니", "감기", "조심", "꽃"}
	fmt.Println("정렬하기 전 :", str)
	sort.Strings(str)
	fmt.Println("정렬하기 후 :", str)

	//내림차순
	var nums2 = []int{3, 5, 6, 1, 0, 66}
	fmt.Println("정렬하기 전 :", nums2)

	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))
	fmt.Println("정렬하기 후 :", nums2)

	var nums3 = []int{3, 5, 6, 1, 0, 66}
	boolean := sort.IntsAreSorted(nums3) //오름 차순으로 정렬이 되어 있는 확인 하는것
	fmt.Println("결과값 :", boolean)

	var nums4 = []int{5, 8, 9, 6, 2, 1}
	fmt.Println("정렬하기 전 :", nums4)
	sort.IntSlice(nums4).Swap(1, 2) //배열에서 자리 바꾸기
	fmt.Println("정렬하기 후 :", nums4)

	boolean = sort.IntSlice(nums4).Less(3, 4) // x[i] < x[j] 값을 알려주는 것
	fmt.Println("결과값 :", boolean)

}
