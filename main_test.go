package main

import (
	"testing"
)

/*func BenchmarkCreatePaymentSchedule20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreatePaymentSchedule("Payment1.pdf", 20)
	}
}*/
func BenchmarkCreatePaymentSchedule50(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreatePaymentSchedule("Payment2.pdf", 50)
	}
}

/*func BenchmarkCreatePaymentSchedule100(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CreatePaymentSchedule("Payment3.pdf", 100)
	}
}*/
