package structs

type Version struct {
	Major  uint8
	Minor  uint8
	Hotfix uint8
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
