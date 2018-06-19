package main

type IntPair [2]int

//ペアの先頭を返すメソッド
func (ip IntPair) First() int {
	return ip[0]
}

//ペあの末尾を返す
func (ip IntPair) Last() (i int) {
	i = ip[1]
	return
}

type Strings []string

func (s Strings) Join(d string) (sum string) {
	for _, v := range s {
		if sum != "" {
			sum += d
		}
		sum += v
	}
	return
}
func main() {
	ip := IntPair{1, 2}
	println(ip.First())
	println(ip.Last())
	println(Strings{"A", "B", "C"}.Join(","))
}
