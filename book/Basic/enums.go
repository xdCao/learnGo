package Basic

import "fmt"

func main() {

	type Weapon int

	const (
		Arrow Weapon = iota
		Shuriken
		SniperRifle
		Rifle
		Blower
	)

	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)

	var weapon Weapon = Blower

	fmt.Println(weapon)

}
