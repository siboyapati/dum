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

//type MapSum struct {
//	value  int
//	isWord bool
//	child  []*MapSum
//}
//
//func Constructor() MapSum {
//	return MapSum{
//		child: make([]*MapSum, 26),
//	}
//}
//
//func (this *MapSum) Insert(key string, val int) {
//	cur := this
//	for _, v := range key {
//		index := v - 'a'
//		if cur.child[index] == nil {
//			cur.child[index] = &MapSum{
//				child: make([]*MapSum, 26),
//			}
//		}
//		cur = cur.child[index]
//	}
//	cur.value = val
//	cur.isWord = true
//}
//
//func (this *MapSum) Sum(prefix string) int {
//	cur := this
//	var result int
//	for _, v := range prefix {
//		index := v - 'a'
//		if cur.child[index] == nil {
//			return 0
//		}
//		cur = cur.child[index]
//	}
//	q := make([]*MapSum, 0)
//	q = append(q, cur)
//
//	for len(q) != 0 {
//		curNode := q[0]
//		q = q[1:]
//
//		if curNode.isWord {
//			result += curNode.value
//		}
//		for i := 0; i < 26; i++ {
//			if curNode.child[i] != nil {
//				q = append(q, curNode.child[i])
//			}
//		}
//
//	}
//	return result
//}
