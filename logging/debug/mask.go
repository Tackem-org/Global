package debug

type Mask uint8

const (
	FUNCTIONCALLS Mask = 1 << iota
	HELPERLOCKER
	GPRCSERVER
	GPRCCLIENT

	ALL  = ^Mask(0)
	NONE = Mask(0)
)

func (dm Mask) Set(flag Mask) {
	dm = dm | flag
}

func (dm Mask) Clear(flag Mask) {
	dm = dm &^ flag
}

func (dm Mask) Toggle(flag Mask) {
	dm = dm ^ flag
}

func (dm Mask) Has(flags ...Mask) bool {
	for _, flag := range flags {
		if dm&flag == flag {
			return true
		}
	}
	return false
}

func (dm Mask) HasAny(flag Mask) bool {
	for i := FUNCTIONCALLS; i < ALL; i <<= 1 {
		if dm.Has(i) && flag.Has(i) {
			return true
		}
	}
	return false
}
