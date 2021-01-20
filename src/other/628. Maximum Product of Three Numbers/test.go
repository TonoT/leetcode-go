package main

func main() {
	maximumProduct([]int{-1, -2, 1})
}
func maximumProduct(nums []int) int {
	min1, min2 := 1>>31, 1>>31
	max1, max2, max3 := -1<<31, -1<<31, -1<<31
	for _, v := range nums {
		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 && v >= min1 {
			min2 = v
		}
		if v > max1 { //是否比最大的数大
			max3 = max2
			max2 = max1
			max1 = v
		} else { //比最大的数小
			if v > max2 { //是否有第二大的数大
				max3 = max2
				max2 = v
			} else {
				if v > max3 { //是否有第三大的数大
					max3 = v
				}
			}
		}

	}
	if min1*min2*max1 >= max1*max2*max3 {
		return min1 * min2 * max1
	} else {
		return max1 * max2 * max3
	}
}
