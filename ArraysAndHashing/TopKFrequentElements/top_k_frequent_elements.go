package leetcodes

func TopKFrequentElements(nums []int, k int) []int {
	buckets := make([][]int, len(nums)+1)
	hm := make(map[int]int)

	for _, num := range nums {
		_, containsNum := hm[num]
		if !containsNum {
			hm[num] = 0
		}
		hm[num]++
	}

	for key, val := range hm {
		if val == 0 {
			buckets[val] = []int{}
		}

		buckets[val] = append(buckets[val], key)
	}

	res := make([]int, k)
	j := 0
	for i := len(buckets) - 1; i >= 0 && j < k; i-- {
		if buckets[i] == nil {
			continue
		}

		for _, num := range buckets[i] {
			res[j] = num
			j++
		}
	}

	return res
}
