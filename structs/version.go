package structs

import (
	"fmt"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	pb "github.com/Tackem-org/Proto/pb/registration"
)

type Version struct {
	Major  uint8
	Minor  uint8
	Hotfix uint8
}

func (v Version) String() string {
	logging.Debugf(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.structs.Version{%d.%d.%d}.String", v.Major, v.Minor, v.Hotfix)
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Hotfix)
}
func (v Version) GreaterThan(c Version) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.structs.Version{%s}.GreaterThan", v.String())
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] version=%v", c.String())
	if v.Major > c.Major {
		return true
	} else if v.Major < c.Major {
		return false
	}

	if v.Minor > c.Minor {
		return true
	} else if v.Minor < c.Minor {
		return false
	}

	if v.Hotfix > c.Hotfix {
		return true
	} else if v.Hotfix < c.Hotfix {
		return false
	}

	return false
}

func (v Version) GreaterThanOrEqualTo(c Version) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.structs.Version{%s}.GreaterThanOrEqualTo", v.String())
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] version=%v", c.String())
	if v.Major > c.Major {
		return true
	} else if v.Major < c.Major {
		return false
	}

	if v.Minor > c.Minor {
		return true
	} else if v.Minor < c.Minor {
		return false
	}

	if v.Hotfix > c.Hotfix {
		return true
	} else if v.Hotfix < c.Hotfix {
		return false
	}

	return true
}

func (v Version) LessThan(c Version) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.structs.Version{%s}.LessThan", v.String())
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] version=%v", c.String())
	if v.Major < c.Major {
		return true
	} else if v.Major > c.Major {
		return false
	}

	if v.Minor < c.Minor {
		return true
	} else if v.Minor > c.Minor {
		return false
	}

	if v.Hotfix < c.Hotfix {
		return true
	} else if v.Hotfix > c.Hotfix {
		return false
	}

	return false
}

func (v Version) LessThanOrEqualTo(c Version) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.structs.Version{%s}.LessThanOrEqual", v.String())
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] version=%v", c.String())
	if v.Major < c.Major {
		return true
	} else if v.Major > c.Major {
		return false
	}

	if v.Minor < c.Minor {
		return true
	} else if v.Minor > c.Minor {
		return false
	}

	if v.Hotfix < c.Hotfix {
		return true
	} else if v.Hotfix > c.Hotfix {
		return false
	}
	return true
}

func (v Version) EqualTo(c Version) bool {
	logging.Debugf(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.structs.Version{%s}.EqualTo", v.String())
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] version=%v", c.String())
	if v.Major == c.Major && v.Minor == c.Minor && v.Hotfix >= c.Hotfix {
		return true
	}
	return false
}

func (v Version) ToProto() *pb.Version {
	logging.Debugf(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Global.structs.Version{%s}.ToProto", v.String())
	return &pb.Version{
		Major:  uint32(v.Major),
		Minor:  uint32(v.Minor),
		Hotfix: uint32(v.Hotfix),
	}
}
