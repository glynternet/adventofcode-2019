package main

import (
	"fmt"
	"math"
)

func main() {
	moduleMasses := masses{
		100725, 63593, 84738, 143809, 108595,
		94419, 91617, 91573, 102728, 143383,
		74613, 80331, 76530, 139884, 104607,
		107171, 107640, 87284, 120827, 85742,
		62474, 97582, 110668, 73426, 57656,
		70819, 89848, 138732, 54386, 116905,
		107954, 131488, 75056, 97660, 55295,
		146265, 58026, 94712, 73636, 138077,
		61480, 148868, 119364, 145430, 103901,
		134202, 106759, 50254, 82440, 117801,
		80263, 97022, 145229, 57702, 57460,
		58401, 145652, 127341, 123585, 65291,
		70219, 147009, 88728, 72059, 83815,
		99635, 80913, 149475, 61798, 110054,
		102505, 148511, 95160, 50208, 129867,
		57079, 138435, 75865, 63185, 142389,
		78370, 108077, 106438, 86267, 100785,
		101165, 68501, 146079, 122420, 121429,
		62608, 115338, 90667, 131391, 50260,
		85343, 76411, 94432, 130126, 80915}

	part01(moduleMasses)
	part02(moduleMasses)
}


type mass float64

func (m mass) fuelRequired() mass {
	return mass(math.Max(math.Floor(float64(m)/3)-2, 0))
}

func (m mass) fuelForMassAndMassesFuel() mass {
	if m == 0 {
		return 0
	}
	massFuel := m.fuelRequired()
	return massFuel + massFuel.fuelForMassAndMassesFuel()
}

type masses []mass

func (ms masses) fuelForMasses() mass {
	var sum mass
	for _, moduleMass := range ms {
		sum += moduleMass.fuelRequired()
	}
	return sum
}

func (ms masses) fuelForMassesAndMassesFuel() mass {
	var sum mass
	for _, moduleMass := range ms {
		sum += moduleMass.fuelForMassAndMassesFuel()
	}
	return sum
}

func part01(ms masses) {
	fmt.Println(int(ms.fuelForMasses()))
}

func part02(masses masses) {
	fmt.Println(int(masses.fuelForMassesAndMassesFuel()))
}
