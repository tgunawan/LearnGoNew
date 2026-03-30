package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed random agar angkanya selalu berbeda
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1

	percobaan := 0
	maxPercobaan := 7

	fmt.Println("🎯 Game Tebak Angka 1-100")
	fmt.Printf("Kamu punya %d kesempatan.\n\n", maxPercobaan)

	for percobaan < maxPercobaan {
		var tebakan int
		fmt.Printf("[%d/%d] Tebakanmu: ", percobaan+1, maxPercobaan)
		fmt.Scan(&tebakan)
		percobaan++

		switch {
		case tebakan == target:
			fmt.Printf("🎉 BENAR! Jawabannya %d, tebak dalam %d percobaan!\n", target, percobaan)
			return
		case tebakan < target:
			fmt.Println("⬆️  Terlalu kecil!")
		default:
			fmt.Println("⬇️  Terlalu besar!")
		}
	}
	fmt.Printf("😢 Game Over! Jawabannya adalah %d\n", target)
}
