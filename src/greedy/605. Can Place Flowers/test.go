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
