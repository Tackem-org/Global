package debug

type Mask uint8

const (
	FUNCTIONCALLS Mask = 1 << iota
	FUNCTIONARGS
	HELPERLOCKER
	GRPCSERVER
	GRPCCLIENT

	ALL  = ^Mask(0)
	NONE = Mask(0)
)

func (dm Mask) Set(flag Mask) {
	dm |= flag
}

func (dm Mask) Clear(flag Mask) {
	dm &= ^flag
}

func (dm Mask) Toggle(flag Mask) {
	dm ^= flag
}

func (dm Mask) Has(mask Mask) bool {
	return (mask & dm) > 0
}

func (dm Mask) HasAny(mask Mask) bool {
	for i := FUNCTIONCALLS; i < ALL; i <<= 1 {
		if dm.Has(i) && mask.Has(i) {
			return true
		}
	}
	return false
}
