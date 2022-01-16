package structs

type Version struct {
	Major  uint8
	Minor  uint8
	Hotfix uint8
}

func (v Version) HigherThan(c Version) bool {
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

func (v Version) HigherOrEqualThan(c Version) bool {
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

func (v Version) LowerThan(c Version) bool {
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

func (v Version) LowerOrEqualThan(c Version) bool {
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
