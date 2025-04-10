//go:build ignore
// +build ignore

package gopty

/*
#include <sys/time.h>
#include <stdlib.h>
#include <sys/tty.h>
*/
import "C"

type ptmget C.struct_ptmget

var ioctl_PTMGET = C.PTMGET
