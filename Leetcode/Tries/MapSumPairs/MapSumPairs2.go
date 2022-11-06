package tries2

//Input
//["MapSum", "insert", "sum", "insert", "sum"]
//[[], ["apple", 3], ["ap"], ["app", 2], ["ap"]]
//Output
//[null, null, 3, null, 5]
//
//Explanation
//MapSum mapSum = new MapSum();
//mapSum.insert("apple", 3);
//mapSum.sum("ap");           // return 3 (apple = 3)
//mapSum.insert("app", 2);
//mapSum.sum("ap");           // return 5 (apple + app = 3 + 2 = 5)

//https://leetcode.com/problems/map-sum-pairs/

type MapSum struct {
	value  int
	isWord bool
	child  map[rune]*MapSum
}

func Constructor() MapSum {
	return MapSum{
		child: map[rune]*MapSum{},
	}
}

func (this *MapSum) Insert(key string, val int) {
	cur := this
	for _, v := range key {
		if cur.child[v] == nil {
			cur.child[v] = &MapSum{
				child: map[rune]*MapSum{},
			}
		}
		cur = cur.child[v]
	}
	cur.value = val
	cur.isWord = true
}

func (this *MapSum) Sum(prefix string) int {
	cur := this
	var result int
	for _, v := range prefix {
		if cur.child[v] == nil {
			return 0
		}
		cur = cur.child[v]
	}
	q := make([]*MapSum, 0)
	q = append(q, cur)

	for len(q) != 0 {
		curNode := q[0]
		q = q[1:]

		if curNode.isWord {
			result += curNode.value
		}
		for _, v := range curNode.child {
			q = append(q, v)
		}
	}
	return result
}
