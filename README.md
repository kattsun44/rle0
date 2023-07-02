# rle0
Run-length Encoding/Decoding CLI


## Example
```shell
# rle0 aaaaaaaaaa
a10
aaaaaaaaaa
```
```shell
# rle0 "Hello, World\!"
H1,e1,l2,o1,,1, 1,W1,o1,r1,l1,d1,!1
Hello, World!
```
```shell
# rle0 AAAAABBBBCCCDDE
A5,B4,C3,D2,E1
AAAAABBBBCCCDDE
```