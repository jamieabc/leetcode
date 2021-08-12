package main

// Given an integer array of even length arr, return true if it is possible to reorder arr such that arr[2 * i + 1] = 2 * arr[2 * i] for every 0 <= i < len(arr) / 2, or false otherwise.



// Example 1:

// Input: arr = [3,1,3,6]
// Output: false

// Example 2:

// Input: arr = [2,1,2,6]
// Output: false

// Example 3:

// Input: arr = [4,-2,2,-4]
// Output: true
// Explanation: We can take two groups, [-2,-4] and [2,4] to form [-2,-4,2,4] or [2,4,-2,-4].

// Example 4:

// Input: arr = [1,2,4,16,8,4]
// Output: false



// Constraints:

//     2 <= arr.length <= 3 * 104
//     arr.length is even.
//     -105 <= arr[i] <= 105

func canReorderDoubled(arr []int) bool {
    table := make(map[int]int)
    for _, i := range arr {
        table[i]++
    }

    sort.Slice(arr, func(i, j int) bool {
        return abs(arr[i]) < abs(arr[j])
    })

    var half int
    for i := len(arr)-1; i >= 0; i-- {
        if table[arr[i]] > 0 {
            if arr[i] & 1 > 0 {
                return false
            }

            half = arr[i] / 2

            table[half] -= table[arr[i]]

			// arr[i] could have multiple values, reset it to 0, avoid
			// duplicate calculations
			table[arr[i]] = 0

            if table[half] < 0 {
                return false
            }
        }
    }

    return true
}

func canReorderDoubled3(arr []int) bool {
	table := make(map[int]int)
	for _, i := range arr {
		table[i]++
	}

	nums := make([]int, 0)
	for num := range table {
		nums = append(nums, num)
	}

	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) < abs(nums[j])
	})

	n := len(nums)
    var half int
	for i := n-1; i >= 0; i-- {
		if table[nums[i]] > 0 {

			// cannot be odd number, because it cannot be evenly divided
			if nums[i] & 1 > 0 {
				return false
			}

            half = nums[i]/2
			table[half] -= table[nums[i]]

            if table[half] < 0 {
                return false
            }

		} else if table[nums[i]] < 0 {
			return false
		}
	}

	return true
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func canReorderDoubled2(arr []int) bool {
	table := make(map[int]int)

	for _, i := range arr {
		table[i]++
	}

	nums := make([]int, 0)
	for k := range table {
		nums = append(nums, k)
	}

	sort.Ints(nums)
	n := len(nums)

	var i int
	for i = n-1; i >= 0 && nums[i] > 0; i-- {
		half := nums[i] / 2

		if table[nums[i]] > 0 {
			if nums[i] & 1 > 0 {
				return false
			}

			table[half] -= table[nums[i]]

			if table[half] < 0 {
				return false
			}
		}
	}

	for i = 0; i < n && nums[i] < 0; i++ {
		half := nums[i] / 2

		if table[nums[i]] > 0 {
			if nums[i] & 1 > 0 {
				return false
			}
		}

		table[half] -= table[nums[i]]

		if table[half] < 0 {
			return false
		}
	}

	return true
}

func canReorderDoubled1(arr []int) bool {
    sort.Ints(arr)
    n := len(arr)

    if arr[0] == arr[n-1] && arr[0] == 0 {
        return true
    }

    table := make(map[int]int)

    var i int
    for ; i < n; i++ {
        if arr[i] >= 0 {
            break
        }

        if val, ok := table[arr[i]]; !ok || val == 0 {
            if arr[i] & 1 > 0 {
                return false
            }

            table[arr[i]/2]++
        } else {
            table[arr[i]]--
        }
    }

    for j := n-1; j >= i; j-- {
        if val, ok := table[arr[j]]; !ok || val == 0 {
            if arr[j] & 1 > 0 {
                return false
            }

            table[arr[j]/2]++
        } else {
            table[arr[j]]--
        }
    }

    for _, v := range table {
        if v != 0 {
            return false
        }
    }

    return true
}

//	Notes
//	1.	to find array matches specific rule, current order in array doesn't
//		matter, ony relationships among numbers important, so sort can be used
//
//	2.	after sort, smallest negative number first, so smallest / 2 exist
//		largest number / 2 exist, use map to check, since map operation is
//		expensive, tc is really slow
//
//	3.	inspired from https://leetcode.com/problems/array-of-doubled-pairs/discuss/203183/JavaC%2B%2BPython-Match-from-the-Smallest-or-Biggest-100
//
//		lee has more insight on the array number relationships, the idea is
//		same but much more elegant, count numbers, groups negative & positive
//		numbers together, then start from largest abs value, deduct number with
//		half value because they are paired
//
//		8 -> 4 is a pair, can be removed same count of 8 & 4
//		if there's remaining 4, it means 4 -> 2 pair exists
//
//		also, sort numbers by absolute number, which makes it easier to do
//		calcuation

//	4.	inspired form https://leetcode.com/problems/array-of-doubled-pairs/discuss/1396876/C%2B%2BJavaPython-Greedy-Algorithm-Clear-explanation-O(NlogN)-Clean-and-Concise
//
//		count number, then sort array by absolute number, use sorted number
//		to eliminate pairs, tc O(n log(n))
