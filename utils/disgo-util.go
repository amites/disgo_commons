package utils

func ZeroBytes(bytes []byte) {
	for i := range bytes {
		bytes[i] = 0
	}
}

