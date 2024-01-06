package conserve

import (
    "bytes"
    "fmt"
    "os/exec"
)

func execute(executable string, args ...string) (string, error) {
    cmd := exec.Command(executable, args...)
    if cmd.Stdout != nil {
        return "", fmt.Errorf("exec: Stdout already set")
    }
    if cmd.Stderr != nil {
        return "", fmt.Errorf("exec: Stderr already set")
    }
    var stdoutBuf bytes.Buffer
    var stderrBuf bytes.Buffer
    cmd.Stdout = &stdoutBuf
    cmd.Stderr = &stderrBuf
    err := cmd.Run()
    if err != nil {
        return "", fmt.Errorf("Error on conserve %s execution\n%s\n%s", executable, err, stderrBuf.String())
    }
    return stdoutBuf.String(), nil
}
