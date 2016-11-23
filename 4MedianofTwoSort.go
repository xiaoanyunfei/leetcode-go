package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length := len(nums1) + len(nums2)
	tmp := make([]int, length/2+1)
	for i, j := 0, 0; i+j < length/2+1; {
		if i < len(nums1) && j < len(nums2) {
			if nums1[i] < nums2[j] {
				tmp[i+j] = nums1[i]
				i += 1
			} else {
				tmp[i+j] = nums2[j]
				j += 1
			}
		} else if i < len(nums1) {
			tmp[i+j] = nums1[i]
			i += 1
		} else {
			tmp[i+j] = nums2[j]
			j += 1
		}
	}
	if length%2 == 1 {
		return float64(tmp[length/2])
	} else {
		return float64(tmp[length/2-1]+tmp[length/2]) / 2
	}

}

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
	nums1 = []int{1, 2}
	nums2 = []int{3, 4}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
