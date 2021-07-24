package main

import "math"

// There is a new alien language which uses the latin alphabet. However, the order among letters are unknown to you. You receive a list of non-empty words from the dictionary, where words are sorted lexicographically by the rules of this new language. Derive the order of letters in this language.
//
// Example 1:
//
// Input:
// [
//   "wrt",
//   "wrf",
//   "er",
//   "ett",
//   "rftt"
// ]
//
// Output: "wertf"
//
// Example 2:
//
// Input:
// [
//   "z",
//   "x"
// ]
//
// Output: "zx"
//
// Example 3:
//
// Input:
// [
//   "z",
//   "x",
//   "z"
// ]
//
// Output: ""
//
// Explanation: The order is invalid, so return "".
//
// Note:
//
//     You may assume all letters are in lowercase.
//     If the order is invalid, return an empty string.
//     There may be multiple valid order of letters, return any one of them is fine.

func alienOrder(words []string) string {
    inDegree := make([]int, 26)
    unique := make([]bool, 26)
    to := make([][]bool, 26)
    for i := range to {
        to[i] = make([]bool, 26)
    }

    var j, idx, cur int
    prev := -1
    for i, word := range words {
        cur = int(word[0]-'a')
        unique[cur] = true

        j = 0
        if prev != -1 {
            if cur != prev {
                // first char different provide order information
                if !to[prev][cur] {
                    inDegree[cur]++
                    to[prev][cur] = true
                }
                j = 1
            } else {
                // first char same, two words should provide order information
                for j = range word {
                    unique[word[j]-'a'] = true

                    if j == len(words[i-1]) {
						// make sure latter checking won't get out of boundary
						// two words with same prefix, shorter one shold be
						// former than the other
						// e.g. [abc, ab] is invalid, where as [ab, abc] is valid
                        j--
                        break
                    }

                    if word[j] != words[i-1][j] {
                        former, latter := int(words[i-1][j]-'a'), int(word[j]-'a')

                        if !to[former][latter] {
                            inDegree[latter]++
                            to[former][latter] = true
                        }

                        break
                    }
                }

                // all previous characters are same, shorter word shoud be in front
                if words[i-1][j] == word[j] && len(words[i-1]) > len(word) {
                    return ""
                }
                j++
            }
        }
        // process remain words
        for ; j < len(word); j++ {
            unique[word[j]-'a'] = true
        }

        prev = cur
    }

    queue := make([]int, 0)
    var distinct int
    for i := range unique {
        if unique[i] {
            distinct++

            if inDegree[i] == 0 {
                queue = append(queue, i)
            }
        }
    }

    // no char with 0 in degree
    if len(queue) == 0 {
        return ""
    }

    // topological sort
    order := make([]byte, 0)

    for len(queue) > 0 {
        idx = queue[0]
        queue = queue[1:]
        order = append(order, byte('a'+idx))

        for i, val := range to[idx] {
            if val {
                inDegree[i]--

                if inDegree[i] == 0 {
                    queue = append(queue, i)
                }
            }
        }
    }

    if len(order) != distinct {
        return ""
    }
    return string(order)
}

func alienOrder1(words []string) string {
	graph, inDegree := buildGraph1(words)

	var wordCount int
	for i := range inDegree {
		if inDegree[i] != -1 {
			wordCount++
		}
	}

	order := topologicalSort1(graph, inDegree)

	if len(order) == wordCount {
		return string(order)
	}
	return ""
}

func topologicalSort1(graph map[byte][]byte, inDegree []int) []byte {
	sources := make([]byte, 0)
	for i := range inDegree {
		if inDegree[i] == 0 {
			sources = append(sources, byte(i+'a'))
		}
	}

	result := make([]byte, 0)

	for len(sources) > 0 {
		next := sources[0]
		sources = sources[1:]

		result = append(result, next)

		for _, n := range graph[next] {
			inDegree[n-'a']--
			if inDegree[n-'a'] == 0 {
				sources = append(sources, n)
			}
		}
	}

	return result
}

func buildGraph1(words []string) (map[byte][]byte, []int) {
	graph := make(map[byte][]byte)
	inDegree := make([]int, 26)
	for i := range inDegree {
		inDegree[i] = -1
	}

	// reset every existing char in-degree to 0
	for _, word := range words {
		for i := range word {
			inDegree[word[i]-'a'] = 0
		}
	}

	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if words[i][0] == words[j][0] {
				// compare to find first difference
				var k int
				for k = 1; k < len(words[i]) && k < len(words[j]); k++ {
					if words[i][k] != words[j][k] {
						graph[words[i][k]] = append(graph[words[i][k]], words[j][k])
						addDegree(inDegree, words[j][k])
						k++
						break
					}
				}
				// words[i] longer than words[j] and part of them are same,
				// since it's lexical order, this should not happen
				// e.g. ["abc", "ab"]
				if k == len(words[j]) && len(words[i]) > len(words[j]) {
					inDegree[words[i][0]-'a'] = math.MaxInt32
				}
			} else {
				// change word, add another relationship since list in lexical
				// order
				graph[words[i][0]] = append(graph[words[i][0]], words[j][0])
				addDegree(inDegree, words[j][0])

				i = j - 1
				break
			}
		}
	}

	return graph, inDegree
}

func addDegree(inDegree []int, b byte) {
	idx := b - 'a'
	if inDegree[idx] == -1 {
		inDegree[idx] = 1
	} else {
		inDegree[idx]++
	}
}

//	Notes
//	1.	for every word first char, need to put it into graph

//	2.	for char after first difference, also need to add them into graph

//	3.	inspired from https://leetcode.com/problems/alien-dictionary/discuss/545020/%22abc%22%22ab%22-expected-%22%22

//		word is in lexical order, so if all char are same and one reaches end,
//		shorter one should be earlier

//	4.	initialize all existing char to 0

//	5.	clarification https://leetcode.com/problems/alien-dictionary/discuss/70111/The-description-is-wrong
