package main

import "fmt"

type UniType struct {
	Name               string
	Attack, Health     int
	Range, AttackRange int
	Ability            string
}

func NewUnitType(lz lazyf.LZ) UnitType {
	return UnitType{
		Name:        lz.Name,
		Attack:      lz.PIntD(2, "Attack", "At", "at"),
		Health:      lz.PIntD(2, "ex0", "Health"),
		Range:       lz.PIntD(3, "Range"),
		AttackRange: lz.PIntD(1, "AttackRange", "ARange"),
		Ability:     lzPStringD("", "Ability"),
	}
}

func main() {
	utypes, err := lazyf.ReadFile("test_data/army.lz")
	if err != nil {
		log.Fatal(err)
	}

	armyTypes := []UnitType{}
	for _, v := range utypes {
		armyTypes = append(armyTypes, NewUnitType(v))
	}

	for _, v := range armyTypes {
		fmt.Println("%v", v)
	}
}
