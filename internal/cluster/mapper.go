package cluster

type Problem struct {
    Type    string
    Message string
}

func MapStatus(status string) Problem {
    switch status {
    case "CrashLoopBackOff":
        return Problem{Type: "Error", Message: "Pod crashes repeatedly"}
    case "Pending":
        return Problem{Type: "Warning", Message: "Pod is waiting for resources"}
    default:
        return Problem{Type: "Info", Message: "Status unknown"}
    }
}
