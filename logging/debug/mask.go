package debug

type Mask uint8

const (
	FUNCTIONCALLS Mask = 1 << iota
	FUNCTIONARGS
	HELPERLOCKER
	GRPCSERVER
	GRPCCLIENT
	NONE = Mask(0)
)

func (dm *Mask) Set(flag Mask) {
	*dm |= flag
}

func (dm *Mask) Clear(flag Mask) {
	*dm &= ^flag
}

func (dm *Mask) Toggle(flag Mask) {
	*dm ^= flag
}

func (dm *Mask) Has(mask Mask) bool {
	return (mask & *dm) == mask
}

func (dm *Mask) HasAny(mask Mask) bool {
	return (mask & *dm) > 0
}
