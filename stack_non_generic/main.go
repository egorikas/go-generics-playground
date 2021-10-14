package main

func main() {
	stackString := new(stackGeneric[string])
	si := new(stackGeneric[int])

	stackString.push("Привет :)")
	si.push(123)

	res, _ := stackString.pop()
	println(res)
}

type stackGeneric[T any] []T

func (s *stackGeneric[T]) push(i T){
	*s = append(*s, i)
}

func (s *stackGeneric[T]) pop() (T, bool){
	if len(*s) == 0{
		var zero T
		return zero, false
	}

	t := *s
	return t[len(*s)-1], true
}

type stack []interface{}

func (s *stack) push(item interface{}) {
	*s = append(*s, item)
}

func (s *stack) pop() interface{} {
	t := *s
	return t[len(*s)-1]
}
