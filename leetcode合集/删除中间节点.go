package main

/**
若链表中的某个节点，既不是链表头节点，也不是链表尾节点，则称其为该链表的「中间节点」。

假定已知链表的某一个中间节点，请实现一种算法，将该节点从链表中删除。

例如，传入节点c（位于单向链表a->b->c->d->e->f中），将其删除后，剩余链表为a->b->d->e->f



示例：

输入：节点5（位于单向链表4->5->1->9中）
输出：不返回任何数据，从链表中删除传入的节点 5，使链表变为4->1->9

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/delete-middle-node-lcci
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {

}

func main() {

}
