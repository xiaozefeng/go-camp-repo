package main

import (
	"fmt"
	"os/exec"
	"sync"
)

func main() {
	var args = []string{
		"server",
		"merchant",
		"admin",
		"server",
		"mall-product",
		"mall-server",
		"mall-booking",
		"pay-center",
		"member",
		"balance",
		"distribution",
		"coupon",
		"data-process-ability",
		"otmgroup-merchant-external-ability",
		"goods",
	}

	var wg sync.WaitGroup

	for _, v := range args {
		wg.Add(1)
		v := v
		go func() {
			defer wg.Done()
			err := run(v)
			if err != nil {
				fmt.Printf("err = %+v\n", err)
			}
		}()
	}

	wg.Wait()
}

func run(app string) error {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("make app=%s", app))
	out, err := cmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(out)
	return err
}
