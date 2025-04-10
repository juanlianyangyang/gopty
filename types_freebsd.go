//go:build ignore
// +build ignore

package gopty

/*
#include <sys/param.h>
#include <sys/filio.h>
*/
import "C"

const (
	_C_SPECNAMELEN = C.SPECNAMELEN /* max length of devicename */
)

type fiodgnameArg C.struct_fiodgname_arg
