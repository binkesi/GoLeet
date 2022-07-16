package dfsbfs

// https://leetcode.cn/problems/operations-on-tree/submissions/

type LockingTree struct {
	nodes    []int
	status   []int
	children [][]int
}

func ConstructorLocking(parent []int) LockingTree {
	l := len(parent)
	children := make([][]int, l)
	for i := 1; i < l; i++ {
		children[parent[i]] = append(children[parent[i]], i)
	}
	return LockingTree{nodes: parent, status: make([]int, l), children: children}
}

func (this *LockingTree) Lock(num int, user int) bool {
	if this.status[num] != 0 {
		return false
	}
	this.status[num] = user
	return true
}

func (this *LockingTree) Unlock(num int, user int) bool {
	if this.status[num] != user {
		return false
	}
	this.status[num] = 0
	return true
}

func (this *LockingTree) Upgrade(num int, user int) bool {
	cntChi := 0
	for i := num; i >= 0; {
		if this.status[i] != 0 {
			return false
		}
		i = this.nodes[i]
	}
	if this.status[num] != 0 {
		return false
	}
	var unlockChildren func(par int)
	unlockChildren = func(par int) {
		if this.status[par] != 0 {
			this.status[par] = 0
			cntChi++
		}
		if len(this.children[par]) == 0 {
			return
		}
		for _, v := range this.children[par] {
			unlockChildren(v)
		}
	}
	unlockChildren(num)
	if cntChi == 0 {
		return false
	}
	this.status[num] = user
	return true
}

/**
 * Your LockingTree object will be instantiated and called as such:
 * obj := Constructor(parent);
 * param_1 := obj.Lock(num,user);
 * param_2 := obj.Unlock(num,user);
 * param_3 := obj.Upgrade(num,user);
 */
