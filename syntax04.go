package main

import (
	"fmt"
	"math/rand"
	"time" // Seed 생성용 패키지
)

// 난수 추출된 수의 소수 판정 프로그램 v0.1
// 소수 : 1과 자기 자신 외에는 나누어 떨어지지 않는 수 (0과 1은 제외)
func main() {
	// Seed 설정
	seed := time.Now().Unix()
	rand.Seed(seed)

	count := 0
	number := rand.Intn(150) + 2 // 0과 1 제외, 2 ~ 151 수
	fmt.Println("임의로 추출된 수 :", number)

	for i := 1; i <= number; i++ {
		if number%i == 0 {
			count++
		}
	}

	if count == 2 {
		fmt.Println(number, "은(는) 소수입니다!")
	} else {
		fmt.Println(number, "은(는) 소수가 아닙니다~")
	}
}
