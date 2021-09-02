package closure

import "fmt"

func Delay(){
	var fns []func()
	for i:=0 ; i<10; i ++ {
		i:=i
		fns = append(fns, func ()  {
			fmt.Printf("hello,this is %d\n", i)
		})
	}
	for _, fn := range fns {
		fn()
	}

}