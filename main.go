// building a docker container
// taken from this tutorial https://www.youtube.com/watch?v=Utf-A4rODH8
// Sam Duthie 2019

package main

import(
    "fmt"
    "os"
    "os/exec"
)

// docker run <container> cmd args
// go run main.go run <cmd>
func main() {
    switch os.Args[1] {
    case "run":
        run()
    default:
        panic("incorrect argument")
    }
}

func run() {
    fmt.Printf("running %v\n", os.Args[2:])

    cmd := exec.Command(os.Args[2], os.Args[3:]...)
    cmd.Stdin = os.Stdout
    cmd.Stderr = os.Stderr

    must(cmd.Run())


}

func must(err error) {
    if err != nil {
        panic(err)
    }
}

