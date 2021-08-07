package main

// There are n houses in a village. We want to supply water for all the houses by building wells and laying pipes.

// For each house i, we can either build a well inside it directly with cost wells[i - 1] (note the -1 due to 0-indexing), or pipe in water from another well to it. The costs to lay pipes between houses are given by the array pipes, where each pipes[j] = [house1j, house2j, costj] represents the cost to connect house1j and house2j together using a pipe. Connections are bidirectional.

// Return the minimum total cost to supply water to all houses.



// Example 1:

// Input: n = 3, wells = [1,2,2], pipes = [[1,2,1],[2,3,1]]
// Output: 3
// Explanation:
// The image shows the costs of connecting houses using pipes.
// The best strategy is to build a well in the first house with cost 1 and connect the other houses to it with cost 2 so the total cost is 3.



// Constraints:

//     1 <= n <= 104
//     wells.length == n
//     0 <= wells[i] <= 105
//     1 <= pipes.length <= 104
//     pipes[j].length == 3
//     1 <= house1j, house2j <= n
//     0 <= costj <= 105
//     house1j != house2j

func minCostToSupplyWater(n int, wells []int, pipes [][]int) int {
    remain := n

    for i, well := range wells {
        pipes = append(pipes, []int{i+1, i+1, well})
    }

    var cost int
    groups := make([]int, n+1)
    ranks := make([]int, n+1)
    for i := range groups {
        groups[i] = i
        ranks[i] = 1
    }
    ready := make([]bool, n+1)

    sort.Slice(pipes, func(i, j int) bool {
        if pipes[i][2] != pipes[j][2] {
            return pipes[i][2] < pipes[j][2]
        }

        if pipes[j][0] == pipes[j][1] {
            return false
        }

        return true
    })

    for i := 0; remain > 0 && i < len(pipes); i++ {
        pipe := pipes[i]

        if pipe[0] == pipe[1] {
            p1 := parent(groups, pipe[0])

            if !ready[p1] {
                cost += pipe[2]
                remain -= ranks[p1]
                ready[p1] = true
            }

            continue
        }

        p1, p2 := parent(groups, pipe[0]), parent(groups, pipe[1])

        if p1 == p2 {
            continue
        }

        if ready[p1] && ready[p2] {
            continue
        }

        if ready[p1] || ready[p2] {
            cost += pipe[2]

            if ready[p1] {
                remain -= ranks[p2]
            } else {
                remain -= ranks[p1]
            }

            ready[p1] = true
            ready[p2] = true

            groups[p2] = p1
            ranks[p1] += ranks[p2]

            continue
        }

        cost += pipe[2]

        if ranks[p1] >= ranks[p2] {
            groups[p2] = p1
            ranks[p1] += ranks[p2]
        } else {
            groups[p1] = p2
            ranks[p2] += ranks[p1]
        }
    }

    return cost
}

func parent(groups []int, idx int) int {
    if groups[idx] != idx {
        groups[idx] = parent(groups, groups[idx])
    }

    return groups[idx]
}

//	Notes
//	1.	take some time to think about the problem, at first glance it's about
//		minimum spanning tree, but with some slightly different because two
//		types of costs: node it self and edge weight
//
//	2.	while thinking about decision for a node to make, following conditions
//		should be considerd:
//		- do iself already okay
//		- any one of node in all already connected gruops okay
//		- what are other nodes reachable

//		the second and third combined indicates disjoint set (union-find)
//		in order to get minimum cost, combine itself and edge into single
//		array and sort, then select by greedy

//	3.	inspired from solution, there's another way to solve it (not implement)
