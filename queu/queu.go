package queu

import "fmt"

type Queu struct {
	size      int
	list      []int
	lastIndex int
}

func NewQueu(size_ int) *Queu {
	return &Queu{
		size:      size_,
		list:      make([]int, size_),
		lastIndex: -1,
	}
}
func (q *Queu) Push(node int) error {
	if q.IsFill(){
		return fmt.Errorf("queu is Fill")
	}
	q.lastIndex++
	q.list[q.lastIndex] = node
	return nil
}
func (q *Queu) Pop() error {
	if q.IsNull(){
		return fmt.Errorf("queu is Empty")
	}
	q.list = q.list[1:]
	q.lastIndex--
	return nil
}

func (q *Queu) IsNull() bool {
	return q.lastIndex == -1
}
func (q *Queu) IsFill() bool {
	return q.lastIndex == q.size - 1
}

func (q *Queu) Front() (int,error){
	if q.IsNull() {
		return -1,fmt.Errorf("queu is null")
	}
	return q.list[0],nil
}
func (q *Queu) Rear()(int,error){
	if q.IsNull() {
		return -1,fmt.Errorf("queu is null")
	}
	return q.list[q.lastIndex],nil
}

func (q *Queu) Show(){
	for i := 0 ; i <= q.lastIndex; i++ {
		fmt.Println(q.list[i])
	}
}