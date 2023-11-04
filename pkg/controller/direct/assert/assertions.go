package assert

import "k8s.io/klog/v2"

func Fail() {
	klog.Fatalf("assertion failed")
}
