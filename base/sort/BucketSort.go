package sort

/* 桶排序 */

// Asc：升序
func BucketSortAsc(arr []int) []int {
	num := len(arr)
	max := getMaxValue(arr)
	buckets := make([][]int, num)
	index := 0

	for i := 0; i < num; i++ {
		index = arr[i] * (num - 1) / max
		buckets[index] = append(buckets[index], arr[i])
	}

	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			sortInBucket(buckets[i])
			copy(arr[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}
	return arr
}

func sortInBucket(bucket []int) {
	length := len(bucket)
	if length == 1 {
		return
	}
	for i := 1; i < length; i++ {
		backup := bucket[i]
		j := i - 1
		for j >= 0 && backup < bucket[j] {
			bucket[j+1] = bucket[j]
			j--
		}
		bucket[j+1] = backup
	}
}
