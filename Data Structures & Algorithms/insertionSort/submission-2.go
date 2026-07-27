// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
res := make([][]Pair, len(pairs))
	for i := range pairs {
		tmp := pairs[i]
		j := i - 1
		for j >= 0 && pairs[j].Key > tmp.Key {
			pairs[j+1] = pairs[j]
			j--
		}
		pairs[j+1] = tmp
		res[i] = append([]Pair{}, pairs...)
	}
    return res
}
