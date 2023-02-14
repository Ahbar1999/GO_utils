package go_utils

/*
   Disjoin Sets and Union-Find implementation with 'Path compression' and 'Union by rank' for optimization
   Working example: 
   https://leetcode.com/problems/redundant-connection/solutions/3186545/
*/

type DSU struct {
	// rank[i] denotes rank(max size) of the tree which contains ith node
	// max number of nodes are 1001 for now, it can obv be changed
	rank [1001]int
	// parent of ith node
	parent [1001]int
}

func (this *DSU) Initialize() {
	for i, _ := range this.rank {
		this.rank[i] = 0
		this.parent[i] = i
	}
}

func (this *DSU) find(x int) int {
	// if not available then find
	if this.parent[x] != x {
		y := x
		for this.parent[y] != y {
			y = this.parent[y]
		}
		this.parent[x] = y
	}

	return this.parent[x]
}

func (this *DSU) union(x, y int) bool {
	// find parents of x and y
	x_p, y_p := this.find(x), this.find(y)
	// if they belong to the same parent then return abort!
	if x_p == y_p {
		return false
	}

	// find ranks of the trees representing their parents
	x_r, y_r := this.rank[x_p], this.rank[y_p]

	switch x_r <= y_r {
	case true:
		// x's tree is smaller/equal so we move this under larger/equal tree(y's)
		this.parent[x_p] = y_p
		if x_r == y_r {
			// rank increased because height increased of y's tree
			this.rank[y_r] += 1
		}
	case false:
		// otherwise
		this.parent[y_p] = x_p
	}

	// return successful after joining
	return true
}
