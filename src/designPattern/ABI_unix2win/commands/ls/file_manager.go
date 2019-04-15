package ls

func ls(path string) string {
	return "file1\nfile2"
}

func ls_l(path string) string {
	str := "-rw-r--r-- 1 root root        69 Dec 25 21:44 ali.env\n"
	str += "-rw-r--r-- 1 root root       115 Dec 25 22:02 do.sh"
	return str
}

func ls_a(path string) string {
	return ".\n..\nfile1\nfile2"
}