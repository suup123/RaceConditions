package main

import (
	"fmt"
	"sync"
	"time"
)

var bankoBalansas int // Banko balanso kintamasis
var wg sync.WaitGroup // Naudojant gijas "Go" programavimo kalboje taip pat naudojame ir "WaitGroup"
// "WaitGroup" funkcija nelendant į detales - palaukti kol visos gijos bus įvykdytos.
var mu sync.Mutex

func Pirkimas(pirkimoSuma int) {
	defer wg.Done() // Kai įvykdoma funkcija, tik tada "WaitGroup" pereina į "Done" būseną.

	mu.Lock()
	suma := bankoBalansas            // Sumą prilyginam banko balansui
	time.Sleep(1 * time.Microsecond) // Imituojam kažkokį apdorojimą
	suma = suma - pirkimoSuma        // Iš visos sumos atimam pirkinio sumą
	bankoBalansas = suma             // Banko balansui priskiriame gautą sumą
	mu.Unlock()
}

func main() {
	//Tarkime, kad turime 1000 eur banke
	bankoBalansas = 1000
	//Nustatytas periodinis investavimas į 7 skirtingus fondus
	//Bankas nuskaito pinigus tuo pačiu momentu.
	wg.Add(7) // Į "WaitGroup" pridedame 7 gijas
	// Naudojant “go” prieš funkciją – vykdomos funkcijos lygiagrečiai
	go Pirkimas(159) // Vykdo pirkimą už 159
	go Pirkimas(5)   // Vykdo pirkimą už 5
	go Pirkimas(157) // Vykdo pirkimą už 157
	go Pirkimas(4)   // Vykdo pirkimą už 4
	go Pirkimas(158) // Vykdo pirkimą už 158
	go Pirkimas(6)   // Vykdo pirkimą už 6
	go Pirkimas(1)   // Vykdo pirkimą už 1

	wg.Wait() // Laukiama, kol visos gijos bus įvykdytos.
	//Galutinė suma banko sąskaitoje:
	fmt.Println("Galutine suma : ", bankoBalansas)
}
