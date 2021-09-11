package command_test

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"testing"
)

func Test_shell(t *testing.T) {
	cmd := exec.Command("date")

	out, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	fmt.Printf("out: %s\n", out)

	// cmd = exec.Command("grep", "hello")

	// in, _ := cmd.StdinPipe()
	// fmt.Println("date:")
	// output, _ := cmd.StdoutPipe()
	// cmd.Start()
	// in.Write([]byte("hello grep\ngoodbyte grep"))
	// defer in.Close()
	// content, _ := io.ReadAll(output)
	// cmd.Wait()

	// fmt.Printf("content: %s\n", content)

	cmd = exec.Command("bash", "-c", "ls -a -l -h")
	out, err = cmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("out: %s\n", out)
}

func Test_command(t *testing.T) {
	binary, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}
	args := []string{"ls", "-a", "-l", "-h"}

	env := os.Environ()

	err = syscall.Exec(binary, args, env)
	if err != nil {
		panic(err)
	}

}
