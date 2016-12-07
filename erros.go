package main

import "errors"
import "fmt"

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("cant work with 42")
	}
	return arg + 3, nil
}

type arg_error struct {
	arg  int
	prob string
}

func (e *arg_error) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &arg_error{arg, "cant work with it"}
	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	_, e := f2(42)
	if ae, ok := e.(*arg_error); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}