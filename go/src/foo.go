package main

// typedef int (*intFunc) ();
//
// extern void *foo(char *);
//
// int
// bridge_int_func(intFunc f)
// {
//		return f();
// }
//
// int fortytwo()
// {
//	    return 42;
// }
import "C"
import "fmt"

func main() {
	var i C.int = C.fortytwo()



	fmt.Println("Hello", i, C.foo("hllo"))
}
