package main

type FreqStack struct {
	element_freq  map[int]int
	element_index map[int][]int
	freq_element  map[int][]int
	most_freq     int
	stack         []int
}

func Constructor() FreqStack {
	return FreqStack{
		element_freq:  make(map[int]int),   // Initialize element_freq map
		element_index: make(map[int][]int), // Initialize element_index map
		freq_element:  make(map[int][]int), // Initialize freq_element map
		most_freq:     0,
		stack:         []int{},
	}
}

func (this *FreqStack) Push(val int) {
	//element_freq mapping update
	this.element_freq[val]++
	// if freq, ok := this.element_freq[val]; ok {
	// 	freq++
	// } else {
	// 	this.element_freq[val] = 1
	// }
	//element_index update
	//this.element_index[val] = append(this.element_index[val], len(this.stack))

	//freq_element update
	//first remove the element from freq
	freq_to_remove := this.element_freq[val] - 1
	if freq_to_remove > 0 {
		for i, elem := range this.freq_element[freq_to_remove] {
			if elem == val {
				this.freq_element[freq_to_remove] = append(this.freq_element[freq_to_remove][0:i], this.freq_element[freq_to_remove][i+1:]...)
			}
		}
	}
	//second add to new freq
	freq := freq_to_remove + 1
	for i, elem := range this.freq_element[freq] {
		if elem == val {
			this.freq_element[freq] = append(this.freq_element[freq][0:i], this.freq_element[freq][i+1:]...)
		}
	}

	this.freq_element[freq] = append(this.freq_element[freq], val)

	//update most_freqq
	if this.element_freq[val] > this.most_freq {
		this.most_freq = this.element_freq[val]
	}
	//finally appeen to stack
	this.stack = append(this.stack, val)
}

func (this *FreqStack) Pop() int {
	//update
	li := len(this.freq_element[this.most_freq]) - 1
	latest_most_freq_elemt := this.freq_element[this.most_freq][li]
	this.freq_element[this.most_freq] = this.freq_element[this.most_freq][:li]
	if len(this.freq_element[this.most_freq]) == 0 {
		delete(this.freq_element, this.most_freq)
	}
	// update this.most_freq
	new_most_freq := 0
	for i := (this.most_freq - 1); i > 0; i-- {
		if _, ok := this.freq_element[i]; ok {
			new_most_freq = i
			break
		}
	}
	this.most_freq = new_most_freq

	//update element_freq
	this.element_freq[latest_most_freq_elemt]--
	//update element index
	// li = len(this.element_index[latest_most_freq_elemt]) - 1
	// stack_index := this.element_index[latest_most_freq_elemt][li]
	// this.element_index[latest_most_freq_elemt] = this.element_index[latest_most_freq_elemt][:li]
	stack_index := 0
	for i := len(this.stack) - 1; i > 0; i-- {
		elem := this.stack[i]
		if elem == latest_most_freq_elemt {
			stack_index = i
		}
	}
	//update stack
	return_elem := this.stack[stack_index]
	this.stack = append(this.stack[0:stack_index], this.stack[(stack_index+1):]...)
	return return_elem
}

func main() {
	obj := Constructor()
	obj.Push(5)
	param_2 := obj.Pop()
	print(param_2)
}

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
