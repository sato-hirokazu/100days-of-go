func intToRoman(num int) string {
	romanInByte := make([]byte, 0)
	numInSlice := make([]int, 0)
	for num != 0 {
		numInSlice = append(numInSlice, num%10)
		num /= 10
	}
	for i := len(numInSlice)-1; i >= 0; i-- {
		one, five, ten := byte(0), byte(0), byte(0)
		switch i {
		case 0:
			one = 'I'
			five = 'V'
			ten = 'X'
		case 1:
			one = 'X'
			five = 'L'
			ten = 'C'
		case 2:
			one = 'C'
			five = 'D'
			ten = 'M'
		case 3:
			one = 'M'
		}

		switch numInSlice[i] {
		case 0,1,2,3:
			romanInByte = append(romanInByte, bytes.Repeat([]byte{one}, numInSlice[i])...)
		case 4:
			romanInByte = append(romanInByte, one)
			romanInByte = append(romanInByte, five)
		case 5:
			romanInByte = append(romanInByte, five)
		case 6,7,8:
			romanInByte = append(romanInByte, five)
			romanInByte = append(romanInByte, bytes.Repeat([]byte{one}, numInSlice[i]-5)...)
		case 9:
			romanInByte = append(romanInByte, one)
			romanInByte = append(romanInByte, ten)
		}
	}
	return string(romanInByte)
}
