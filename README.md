# hash partial brute

example go program to brute force part of the hash.

Note that performance is really low now. It can brute force first 3 bytes of sha1 in seconds, 4 bytes is a challange.

I would like to GPU accelerate project one day



### output
```
1
2
3
4
fba810b15a966c533fb0007adcd9283889c44ed7 {"user":"admin","msg":"+pq`"} 
[123 34 117 115 101 114 34 58 34 97 100 109 105 110 34 44 34 109 115 103 34 58 34 43 112 113 96 34 125]
```

### check

```
echo -n '{"user":"admin","msg":"+pq`"}' | sha1sum
fba810b15a966c533fb0007adcd9283889c44ed7  -
```