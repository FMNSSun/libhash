# libhash

A library making hashing more conventient. 

```go
h, err := libhash.HashFile("sha1", "/etc/passwd")

if err != nil {
	fmt.Println(h)
}

fmt.Println(libhash.HashStr("md5", "Hello, world!"))
```
