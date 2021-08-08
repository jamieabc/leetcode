package main

// Given a string s, partition s such that every substring of the partition is a palindrome.

// Return the minimum cuts needed for a palindrome partitioning of s.



// Example 1:

// Input: s = "aab"
// Output: 1
// Explanation: The palindrome partitioning ["aa","b"] could be produced using 1 cut.

// Example 2:

// Input: s = "a"
// Output: 0

// Example 3:

// Input: s = "ab"
// Output: 1



// Constraints:

//     1 <= s.length <= 2000
//     s consists of lower-case English letters only.

func minCut(s string) int {
    n := len(s)
    table := make([][]bool, n)
    for i := range table {
        table[i] = make([]bool, n)
        table[i][i] = true
    }

    for i := range s {
        // center at i
        for d := 1; i+d < n && i-d >= 0; d++ {
            if s[i-d] == s[i+d] {
                table[i-d][i+d] = true
            } else {
                break
            }
        }

        // center at i & i+1
        if i+1 < n && s[i] == s[i+1] {
            table[i][i+1] = true

            for d := 1; i-d >= 0 && i+1+d < n; d++ {
                if s[i-d] == s[i+1+d] {
                    table[i-d][i+1+d] = true
                } else {
                    break
                }
            }
        }
    }

    dp := make([]int, n+1)
    dp[n-1] = 1

    for i := n-2; i >= 0; i-- {
        shortest := n-i

        for j := i; j < n; j++ {
            if table[i][j] {
                shortest = min(shortest, 1+dp[j+1])
            }
        }

        dp[i] = shortest
    }

    return dp[0]-1
}

func min(i, j int) int {
    if i <= j {
        return i
    }
    return j
}

//	Notes
//	1.	creat a table to check from index i ~ j is palindrome or not, tc O(n^2)
//		after creating this table, for index i, try all j (j >= i) to find
//		shortest palindrome count, tc O(n!)
//
//	2.	when creating palndrome table, could use a better way of construction:
//		for every index i or (i & i+1), try to see if this index can form a
//		palindrom center at i or (i & i+1)
//
//
//		the good of this is if longer length is palindrome, then shorter length
//		is also palindrome, which means if at some length not palindrome, later
//		length doesn't need to be checked

//	3.	while first try is TLE, suddently figure out that recurring problem
//		is what's the shortest palindrome start from index j
//
//		if index i ~ j-1 is palindrome, and also knows palindrome count from
//		j ~ end, then i ~ end = 1 + palindrome_count(j~end)
//
//		this is similar to LIS O(n^2)
//
//		apply this technique, overall tc O(n^2 + n^2)
