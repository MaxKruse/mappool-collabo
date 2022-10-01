package modenum

const (
	None              = 0
	NoFail            = 1
	Easy              = 2
	TouchDevice       = 4
	Hidden            = 8
	HardRock          = 16
	SuddenDeath       = 32
	DoubleTime        = 64
	Relax             = 128
	HalfTime          = 256
	Nightcore         = 512 // Only set along with DoubleTime. i.e: NC only gives 576
	Flashlight        = 1024
	Autoplay          = 2048
	SpunOut           = 4096
	Relax2            = 8192  // Autopilot
	Perfect           = 16384 // Only set along with SuddenDeath. i.e: PF only gives 16416
	Key4              = 32768
	Key5              = 65536
	Key6              = 131072
	Key7              = 262144
	Key8              = 524288
	FadeIn            = 1048576
	Random            = 2097152
	Cinema            = 4194304
	Target            = 8388608
	Key9              = 16777216
	KeyCoop           = 33554432
	Key1              = 67108864
	Key3              = 134217728
	Key2              = 268435456
	ScoreV2           = 536870912
	Mirror            = 1073741824
	KeyMod            = Key1 | Key2 | Key3 | Key4 | Key5 | Key6 | Key7 | Key8 | Key9 | KeyCoop
	FreeModAllowed    = NoFail | Easy | Hidden | HardRock | SuddenDeath | Flashlight | FadeIn | Relax | Relax2 | SpunOut | KeyMod
	ScoreIncreaseMods = Hidden | HardRock | DoubleTime | Flashlight | FadeIn
)

func ModIntsToStringArray(mods int64) []string {
	var modStrings []string

	if mods&NoFail != 0 {
		modStrings = append(modStrings, "NF")
	}

	if mods&Easy != 0 {
		modStrings = append(modStrings, "EZ")
	}

	if mods&TouchDevice != 0 {
		modStrings = append(modStrings, "TD")
	}

	if mods&Hidden != 0 {
		modStrings = append(modStrings, "HD")
	}

	if mods&HardRock != 0 {
		modStrings = append(modStrings, "HR")
	}

	if mods&SuddenDeath != 0 {
		modStrings = append(modStrings, "SD")
	}

	if mods&DoubleTime != 0 {
		modStrings = append(modStrings, "DT")
	}

	if mods&Relax != 0 {
		modStrings = append(modStrings, "RX")
	}

	if mods&HalfTime != 0 {
		modStrings = append(modStrings, "HT")
	}

	if mods&Nightcore != 0 {
		modStrings = append(modStrings, "NC")
	}

	if mods&Flashlight != 0 {
		modStrings = append(modStrings, "FL")
	}

	if mods&Autoplay != 0 {
		modStrings = append(modStrings, "AT")
	}

	if mods&SpunOut != 0 {
		modStrings = append(modStrings, "SO")
	}

	if mods&Relax2 != 0 {
		modStrings = append(modStrings, "AP")
	}

	if mods&Perfect != 0 {
		modStrings = append(modStrings, "PF")
	}

	if mods&Key4 != 0 {
		modStrings = append(modStrings, "4K")
	}

	if mods&Key5 != 0 {
		modStrings = append(modStrings, "5K")
	}

	if mods&Key6 != 0 {
		modStrings = append(modStrings, "6K")
	}

	if mods&Key7 != 0 {
		modStrings = append(modStrings, "7K")
	}

	if mods&Key8 != 0 {
		modStrings = append(modStrings, "8K")
	}

	if mods&FadeIn != 0 {
		modStrings = append(modStrings, "FI")
	}

	if mods&Random != 0 {
		modStrings = append(modStrings, "RD")
	}

	return modStrings
}

func ModStringsToInt64(modStrings []string) int64 {
	var mods int64

	for _, mod := range modStrings {
		switch mod {
		case "NF":
			mods |= NoFail
		case "EZ":
			mods |= Easy
		case "TD":
			mods |= TouchDevice
		case "HD":
			mods |= Hidden
		case "HR":
			mods |= HardRock
		case "SD":
			mods |= SuddenDeath
		case "DT":
			mods |= DoubleTime
		case "RX":
			mods |= Relax
		case "HT":
			mods |= HalfTime
		case "NC":
			mods |= Nightcore
		case "FL":
			mods |= Flashlight
		case "AT":
			mods |= Autoplay
		case "SO":
			mods |= SpunOut
		case "AP":
			mods |= Relax2
		case "PF":
			mods |= Perfect
		case "4K":
			mods |= Key4
		case "5K":
			mods |= Key5
		case "6K":
			mods |= Key6
		case "7K":
			mods |= Key7
		case "8K":
			mods |= Key8
		case "FI":
			mods |= FadeIn
		case "RD":
			mods |= Random
		}
	}

	return mods
}
