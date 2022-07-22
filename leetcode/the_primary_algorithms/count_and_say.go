package the_primary_algorithms

import (
	"bytes"
)

var buffer bytes.Buffer

func CountAndSay(n int) string {
	component := "1"

	for i := 1; i < n; i++ {
		// var ans string
		var l, r int
		for r < len(component) {
			if component[r] != component[l] {

				buffer.WriteByte(byte(r-l) + '0')
				buffer.WriteByte(component[l])

				l = r
			}
			r++
		}
		buffer.WriteByte(byte(r-l) + '0')
		buffer.WriteByte(component[l])


		component = buffer.String()

		buffer.Reset()
	}
	

	return component
}
