package main

// You have a graph of n nodes labeled from 0 to n - 1. You are given an integer n and a list of edges where edges[i] = [ai, bi] indicates that there is an undirected edge between nodes ai and bi in the graph.

// Return true if the edges of the given graph make up a valid tree, and false otherwise.



// Example 1:

// Input: n = 5, edges = [[0,1],[0,2],[0,3],[1,4]]
// Output: true

// Example 2:

// Input: n = 5, edges = [[0,1],[1,2],[2,3],[1,3],[1,4]]
// Output: false



// Constraints:

//     1 <= 2000 <= n
//     0 <= edges.length <= 5000
//     edges[i].length == 2
//     0 <= ai, bi < n
//     ai != bi
//     There are no self-loops or repeated edges.

func validTree(n int, edges [][]int) bool {
    groups := make([]int, n)
    for i := range groups {
        groups[i] = i
    }

    for _, edge := range edges {
        p1, p2 := find(groups, edge[0]), find(groups, edge[1])

        // an edge connected 2 nodes with same parent, this edge
        // forms a loop
        if p1 == p2 {
            return false
        }

        if p1 < p2 {
            groups[p2] = p1
        } else {
            groups[p1] = p2
        }
    }

    // check there's only one root
    root := find(groups, 0)
    for i := 1; i < n; i++ {
        if find(groups, i) != root {
            return false
        }
    }

    return true
}

func find(groups []int, idx int) int {
    if groups[idx] != idx {
        groups[idx] = find(groups, groups[idx])
    }

    return groups[idx]
}

func validTree2(n int, edges [][]int) bool {
    // n nodes, at most n-1 edges to form a tree
    if n != len(edges)+1 {
        return false
    }

    // avoid edge[0] access
    if n == 1 {
        return true
    }

    adjList := make([][]int, n)
    for i := range adjList {
        adjList[i] = make([]int, 0)
    }

    for _, edge := range edges {
        adjList[edge[0]] = append(adjList[edge[0]], edge[1])
        adjList[edge[1]] = append(adjList[edge[1]], edge[0])
    }

    visited := make([]bool, n)
    // here assumes edges[0] exist, but it's wrong for only one node (no edge)
    visited[edges[0][0]] = true
    queue := adjList[edges[0][0]]

    for len(queue) > 0 {
        size := len(queue)

        for i := 0; i < size; i++ {
            if !visited[queue[i]] {
                visited[queue[i]] = true
                queue = append(queue, adjList[queue[i]]...)
            }
        }

        queue = queue[size:]
    }

    for i := range visited {
        if !visited[i] {
            return false
        }
    }

    return true
}

func validTree1(n int, edges [][]int) bool {
    // only one node, there won't be any edge
    if n == 1 {
        return len(edges) == 0
    }

    counter := make([]int, n)
    adjList := make([][]int, n)
    for i := range adjList {
        adjList[i] = make([]int, 0)
    }
    exist := make([]bool, n)

    for _, edge := range edges {
        adjList[edge[0]] = append(adjList[edge[0]], edge[1])
        adjList[edge[1]] = append(adjList[edge[1]], edge[0])

        if !exist[edge[0]] {
            exist[edge[0]] = true
        }

        if !exist[edge[1]] {
            exist[edge[1]] = true
        }

        counter[edge[0]]++
        counter[edge[1]]++
    }

    queue := make([]int, 0)
    for i := range counter {
        if counter[i] == 1 {
            queue = append(queue, i)
        } else if counter[i] == 0 {
            return false
        }
    }

    var size, node int
    for len(queue) > 0 {
        size = len(queue)

        for i := 0; i < size; i++ {
            node = queue[i]
            exist[node] = false

            for _, j := range adjList[node] {
                if exist[j] {
                    counter[j]--

                    if counter[j] == 1 {
                        queue = append(queue, j)
                    }
                }
            }
        }

        queue = queue[size:]
    }

    // make sure there's only one root
    var root int
    for _, i := range counter {
        if i == 0 {
            root++
        } else if i > 1 {
            // if there a node with edge count > 1, mean's there's loop
            return false
        }
    }

    return root == 1
}

//	Notes
//	1.	at first think about using in-degree to find root, but since edges
//		cannot distingush in or out, so that won't work
//
//	2.	start from leaf node, every time remove leaf node (edge count == 1)
//		and reduce ege count of adjacent nodes connected by leaf node, repeat
//		this process until there are only 2 or less nodes exist
//
//      for a valid tree, it is guaranteed that leaf node has ony one edge,
//      the only excepton is if root node connects to only one node, so take
//      care of that excepton should work
//
//  3.  make sure there's only one root, check only connected == 0 is not enough,
//      consider the condition 1 - 2, 3 - 4
//
//  4.  only exception is one node, so there's no edge to check

//  5.  4 months ago, used union-find to check, the rule is simple, always
//      choose smaller node which makes result normalized

//  6.  inspired from solution, a valid tree satifies following conditions:
//      - full connected
//      - no cycle
//
//      very concise and elegant conclusion...even though solve the problem,
//      i had no insight on this, fragment congnition
//
//      w/o cycle and fully connected means any node can be root, that's the
//      reason dfs could work
