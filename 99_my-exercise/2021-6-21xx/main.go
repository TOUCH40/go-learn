type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
			stack1: list.New(),
			stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int)  {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	// 如果第二个栈为空
	if this.stack2.Len() == 0 {
			for this.stack1.Len() > 0 {
					this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
			}
	}
	if this.stack2.Len() != 0 {
			e := this.stack2.Back()
			this.stack2.Remove(e)
			return e.Value.(int)
	}
	return -1
}

作者：LeetCode-Solution
链接：https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/solution/mian-shi-ti-09-yong-liang-ge-zhan-shi-xian-dui-l-3/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。