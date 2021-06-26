var romans = map[int]string{
	1: "I",
	5: "V",
	10:"X",
	50:"L",
	100:"C",
	500:"D",
	1000:"M",
}
func intToRoman(num int) string {
    var res string
	d:=1
	for sub:=num; sub>0; sub/=10 {
		remainder:=sub%10
		n:=remainder*d
		if n>=1{
			half:=d*10/2
			if n==half-d{
				res=romans[d]+romans[half]+res
			} else if n==d*10-d{
				res=romans[d]+romans[d*10]+res
			} else if n<half{
				res=strings.Repeat(romans[d],remainder)+res
			} else {
				res=romans[half]+strings.Repeat(romans[d],remainder-5)+res
			}
		}
		d*=10
		if d>1000{
			break
		}
	}
	return res
}
