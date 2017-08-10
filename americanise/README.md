
<p>
british-american.txt 里边是英式英语到美式英语的映射，每一行两个值由空格隔开<br>
本程序将in.txt中的英式英语替换为美式英语
</p>

<pre><code>
/usr/lib/golang/bin/go run /root/github/go/americanise/main.go --help
exit status 1
usage: main [<]infile.txt [>] outfile.txt
</pre></code>

<pre><code>
british-american.txt:
ha la

in.txt:
ha

/usr/lib/golang/bin/go run /root/github/go/americanise/main.go ./in.txt
la

-------------
/usr/lib/golang/bin/go run /root/github/go/americanise/main.go ./in.txt ./out.txt

生成out.txt:
la

</pre></code>

test
-----------------------------------
<pre><code>
package main

import "fmt"

func main()  {
	usForBritish := make(map[string]string)
	usForBritish["a"] = "b"
	word := "a"
	usWord, found := usForBritish[word]

	fmt.Println(usWord)
	fmt.Println(found)
}
---------------------------------
/usr/lib/golang/bin/go run /root/GoglandProjects/test4/main.go
b
true

如果 word="sdfsdf"
/usr/lib/golang/bin/go run /root/GoglandProjects/test4/main.go

false

</pre></code>