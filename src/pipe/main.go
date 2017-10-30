package main

import "fmt"
import "os/exec"
import "bufio"
import (
	"bytes"
	"strings"
	"os"
	"syscall"
	"io/ioutil"
)

func t1() {
	//create cmd
	cmd_go_env := exec.Command("go", "env")
	cmd_grep := exec.Command("grep", "GOROOT")

	stdout_env, env_error := cmd_go_env.StdoutPipe()
	if env_error != nil {
		fmt.Println("Error happened about standard output pipe ", env_error)
		return
	}

	//env_error := cmd_go_env.Start()
	if env_error := cmd_go_env.Start(); env_error != nil {
		fmt.Println("Error happened in execution ", env_error)
		return
	}

	//get the output of go env
	stdout_buf_grep := bufio.NewReader(stdout_env)

	//create input pipe for grep command
	stdin_grep, grep_error := cmd_grep.StdinPipe()
	if grep_error != nil {
		fmt.Println("Error happened about standard input pipe ", grep_error)
		return
	}

	//connect the two pipes together
	stdout_buf_grep.WriteTo(stdin_grep)

	//set buffer for reading
	var buf_result bytes.Buffer
	cmd_grep.Stdout = &buf_result

	//grep_error := cmd_grep.Start()
	if grep_error := cmd_grep.Start(); grep_error != nil {
		fmt.Println("Error happened in execution ", grep_error)
		return
	}

	err := stdin_grep.Close()
	if err != nil {
		fmt.Println("Error happened in closing pipe", err)
		return
	}

	//make sure all the infor in the buffer could be read
	if err := cmd_grep.Wait(); err != nil {
		fmt.Println("Error happened in Wait process")
		return
	}
	fmt.Println(buf_result.String())
}

func t2() {
	//create cmd
	cmd_grep := exec.Command("grep", "GOROOT")

	text := strings.NewReader("ABCDEFG\nGOROOT=123\nsadfsdfgdsfgfgdg\nhshshshss")
	stdout_buf_grep := bufio.NewReader(text)

	//create input pipe for grep command
	stdin_grep, grep_error := cmd_grep.StdinPipe()
	if grep_error != nil {
		fmt.Println("Error happened about standard input pipe ", grep_error)
		return
	}

	//connect the two pipes together
	stdout_buf_grep.WriteTo(stdin_grep)

	//set buffer for reading
	var buf_result bytes.Buffer
	cmd_grep.Stdout = &buf_result

	//grep_error := cmd_grep.Start()
	if grep_error := cmd_grep.Start(); grep_error != nil {
		fmt.Println("Error happened in execution ", grep_error)
		return
	}

	err := stdin_grep.Close()
	if err != nil {
		fmt.Println("Error happened in closing pipe", err)
		return
	}

	//make sure all the infor in the buffer could be read
	if err := cmd_grep.Wait(); err != nil {
		fmt.Println("Error happened in Wait process")
		return
	}
	fmt.Println(buf_result.String())
}

func t3() {
	//create cmd
	cmd_grep := exec.Command("less")

	text := strings.NewReader("ABCDEFG\nGOROOT=123\nsadfsdfgdsfgfgdg\nhshshshss")
	stdout_buf_grep := bufio.NewReader(text)

	//create input pipe for grep command
	stdin_grep, grep_error := cmd_grep.StdinPipe()
	if grep_error != nil {
		fmt.Println("Error happened about standard input pipe ", grep_error)
		return
	}

	//connect the two pipes together
	stdout_buf_grep.WriteTo(stdin_grep)

	cmd_grep.Stdout = os.Stdout

	//grep_error := cmd_grep.Start()
	if grep_error := cmd_grep.Start(); grep_error != nil {
		fmt.Println("Error happened in execution ", grep_error)
		return
	}

	if err := cmd_grep.Wait(); err != nil {
		fmt.Println("Error happened in Wait process")
		return
	}

}

func t4() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	//Exec 需要的参数是切片的形式的（不是放在一起的一个大字符串）。我们给 ls 一些基本的参数。注意，第一个参数需要是程序名。
	args := []string{"ls", "-a", "-l", "-h"}
	//Exec 同样需要使用环境变量。这里我们仅提供当前的环境变量。
	env := os.Environ()
	//这里是 os.Exec 调用。如果这个调用成功，那么我们的进程将在这里被替换成 /bin/ls -a -l -h 进程。如果存在错误，那么我们将会得到一个返回值。
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func t5() {
	binary, lookErr := exec.LookPath("less")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"less", "/tmp/a.txt"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

var text =
`options:
        -name: sdfdfsdf
        - sdfsdfsdfsdfffffffffff
`

func t6() {
	ioutil.WriteFile("/tmp/dat1", []byte(text), 0644)
}

func t7() {

	ioutil.WriteFile("/tmp/dat1", []byte(text), 0644)

	binary, lookErr := exec.LookPath("less")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"less", "/tmp/dat1"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}

func main() {
	t7()
}