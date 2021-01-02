package main

func main() {

}

func canPlaceFlowers(flowerbed []int, n int) bool {
	num := 1
	cnt := 0
	for i, v := range flowerbed {
		if v == 1 || i == len(flowerbed)-1 {
			if i == len(flowerbed)-1 && v == 0 {
				num += 2
			}
			cnt += (num - 1) / 2
			if cnt >= n {
				return true
			}
			num = 0
		} else {
			num += 1
		}

	}
	return false
}
func canPlaceFlowers2(flowerbed []int, n int) bool {
	length := len(flowerbed)
	i := 0
	for i < length {
		// 1. 种完了
		if n <= 0 {
			break
		}

		// 2. 当前格子已经种了花，则直接跳2格（因为相邻的即使为0也不能放）
		if flowerbed[i] == 1 {
			i += 2
			continue
		}

		if (i == 0 || flowerbed[i-1] == 0) && (i == length-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}

		i++
	}

	if n != 0 {
		return false
	}
	return true
}
