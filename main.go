// building a docker container
// taken from this tutorial https://www.youtube.com/watch?v=Utf-A4rODH8
// Sam Duthie 2019

package main

import(
    "fmt"
    "os"
    "os/exec"
    "syscall"
)

// docker run <container> cmd args
// go run main.go run <cmd>
func main() {
    switch os.Args[1] {
    case "run":
        run()
    case "child":
        child()
    default:
        panic("incorrect argument")
    }
}

func run() {
    cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
    cmd.Stdin = os.Stdout
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr

    // Call system commands with new name spaces
    cmd.SysProcAttr = &syscall.SysProcAttr {
        // flag to not call host machine
        Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID,
    }

    must(cmd.Run())


}

func child() {
    fmt.Printf("running %v as PID %d \n", os.Args[2:], os.Getpid())

    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    cmd.Stdin = os.Stdout
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr


    must(cmd.Run())


}

func must(err error) {
    if err != nil {
        panic(err)
    }
}

