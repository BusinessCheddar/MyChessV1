package main

import (
	"fmt"
)

// Stück zeigt Schachstück
type Stück struct {
	Farbe string // Farbe des Stücks: weiß schwarz
	Typ   string // Typ des Stücks: könig dame turm bauer, läufer springer
}

// Schachbrett zeigt Schachbrett
type Schachbrett struct {
	Feld [][]*Stück // 2D-Array von Stücken, schachbrett
}

// NeueInstanz erstellt ein neues Schachbrett
func NeueInstanz() *Schachbrett {
	brett := &Schachbrett{Feld: make([][]*Stück, 8)} // 8 mal 8 Schachfeld
	for i := range brett.Feld {
		brett.Feld[i] = make([]*Stück, 8)
	}
	// Initialisiere die Schachstücke, hab schon mal den Beispiel: Weißer König
	brett.Feld[0][4] = &Stück{Farbe: "Weiß", Typ: "König"}
	// Weitere Stücke initialisieren...hier müsstest man die startformation schreiben!
	return brett
}

// Funktion zeigt Schachbrett an
func (brett *Schachbrett) anzeigen() {
	for _, zeile := range brett.Feld {
		for _, stück := range zeile {
			if stück != nil { // das wollte ich vorhin hinzufügen; es schaut ob ein Element an der jeweiligen Position einen wert hat. Nil ist in Go ein spezieller Wert um zu zeigen dass Zeiger keinen Wert hat
				fmt.Printf("[%s %s] ", stück.Farbe, stück.Typ)
			} else {
				fmt.Print("[  ] ")
			}
		}
		fmt.Println()
	}
}

func main() {
	//  Schachbretts erstellen
	schachbrett := NeueInstanz()

	// Schachbrett anzeigen
	schachbrett.anzeigen()
}
