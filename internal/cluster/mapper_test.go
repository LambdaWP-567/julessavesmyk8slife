package cluster

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestMapStatus(t *testing.T) {
    tests := []struct {
        status   string
        expected Problem
    }{
        {"CrashLoopBackOff", Problem{Type: "Error", Message: "Pod crashes repeatedly"}},
        {"Pending", Problem{Type: "Warning", Message: "Pod is waiting for resources"}},
        {"Unknown", Problem{Type: "Info", Message: "Status unknown"}},
    }

    for _, tt := range tests {
        t.Run(tt.status, func(t *testing.T) {
            assert.Equal(t, tt.expected, MapStatus(tt.status))
        })
    }
}
