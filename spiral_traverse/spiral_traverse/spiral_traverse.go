package spiral_traverse

func SpiralTraverse(array [][]int) []int {
	fc := 0
	lc := len(array[0]) - 1
	fr := 0
	lr := len(array) - 1
	itemCount := len(array) * len(array[0])
	spiralFlattened := []int{}

	for fc <= lc || fr <= lr {

		for i := fc; i <= lc; i++ {
			spiralFlattened = append(spiralFlattened, array[fr][i])
		}
		if len(spiralFlattened) >= itemCount {
			break
		}
		fr++

		for i := fr; i <= lr; i++ {
			spiralFlattened = append(spiralFlattened, array[i][lc])
		}
		if len(spiralFlattened) >= itemCount {
			break
		}
		lc--

		for i := lc; i >= fc; i-- {
			spiralFlattened = append(spiralFlattened, array[lr][i])
		}
		if len(spiralFlattened) >= itemCount {
			break
		}
		lr--

		for i := lr; i >= fr; i-- {
			spiralFlattened = append(spiralFlattened, array[i][fc])
		}
		if len(spiralFlattened) >= itemCount {
			break
		}
		fc++
	}
	return spiralFlattened
}
