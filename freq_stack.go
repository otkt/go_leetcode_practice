// https://leetcode.com/problems/maximum-frequency-stack/description/
package main

type FreqStack struct {
	element_freq map[int]int   // element to freq histogram
	freq_element map[int][]int // freq_count to element list (element list ordered from earliest to latest for next element in the stack for that given freq)
	most_freq    int           // highest freq count
	stack        []int         //actual stack
}

func Constructor() FreqStack {
	return FreqStack{
		element_freq: make(map[int]int),
		freq_element: make(map[int][]int),
		most_freq:    0,
		stack:        []int{},
	}
}

func (f *FreqStack) Push(val int) {
	//element_freq mapping update
	f.element_freq[val]++

	f.freq_element[f.element_freq[val]] = append(f.freq_element[f.element_freq[val]], val)

	//update most_freqq
	if f.element_freq[val] > f.most_freq {
		f.most_freq = f.element_freq[val]
	}
	//finally appeen to stack
	f.stack = append(f.stack, val)
}

func (f *FreqStack) Pop() int {
	//update freq_element
	li := len(f.freq_element[f.most_freq]) - 1
	latest_most_freq_elemt := f.freq_element[f.most_freq][li]
	f.freq_element[f.most_freq] = f.freq_element[f.most_freq][:li]
	if len(f.freq_element[f.most_freq]) == 0 {
		delete(f.freq_element, f.most_freq)
	}

	// update this.most_freq
	new_most_freq := 0
	for i := (f.most_freq); i > 0; i-- {
		if _, ok := f.freq_element[i]; ok {
			new_most_freq = i
			break
		}
	}
	f.most_freq = new_most_freq

	//update element_freq
	f.element_freq[latest_most_freq_elemt]--

	stack_index := 0
	for i := len(f.stack) - 1; i > 0; i-- {
		elem := f.stack[i]
		if elem == latest_most_freq_elemt {
			stack_index = i
			break
		}
	}
	//update stack
	return_elem := f.stack[stack_index]
	f.stack = append(f.stack[0:stack_index], f.stack[(stack_index+1):]...)
	return return_elem
}

// func main() {
// 	obj := Constructor()
// 	obj.Push(5)
// 	obj.Push(7)
// 	obj.Push(5)
// 	obj.Push(7)
// 	obj.Push(4)
// 	obj.Push(5)
// 	p0 := obj.Pop()
// 	p1 := obj.Pop()
// 	p2 := obj.Pop()
// 	p3 := obj.Pop()
// 	print(p0, p1, p2, p3)
// }

/**
 * Your FreqStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * param_2 := obj.Pop();
 */
