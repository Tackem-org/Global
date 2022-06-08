package structs

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"

	pb "github.com/Tackem-org/Global/pb/registration"
)

type Version struct {
	Major  uint8
	Minor  uint8
	Hotfix uint8
}

func StringToVersion(v string) Version {
	vLower := strings.ToLower(v)
	vRemoveExtra := strings.Split(vLower, "-")
	vRemovev := strings.ReplaceAll(vRemoveExtra[0], "v", "")
	splitv := strings.Split(vRemovev, ".")
	major, _ := strconv.Atoi(splitv[0])
	minor, _ := strconv.Atoi(splitv[1])
	hotfix, _ := strconv.Atoi(splitv[2])

	return Version{
		Major:  uint8(major),
		Minor:  uint8(minor),
		Hotfix: uint8(hotfix),
	}
}

func FileToVersion(file string) Version {
	content, _ := ioutil.ReadFile(file)
	text := string(content)
	return StringToVersion(text)
}

func (v Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Hotfix)
}

func (v Version) GreaterThan(c Version) bool {
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
	if v.Major == c.Major && v.Minor == c.Minor && v.Hotfix >= c.Hotfix {
		return true
	}
	return false
}

func (v Version) EqualToHotfix(c Version) bool {
	if v.Major == c.Major && v.Minor == c.Minor && v.Hotfix == c.Hotfix {
		return true
	}
	return false
}

func (v Version) ToProto() *pb.Version {
	return &pb.Version{
		Major:  uint32(v.Major),
		Minor:  uint32(v.Minor),
		Hotfix: uint32(v.Hotfix),
	}
}
