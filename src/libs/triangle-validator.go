package libs

// ITriangleValidator validates if is a triangle
type ITriangleValidator interface {
	Validate() bool
}

// TriangleValidator implements ITriangleValidator
type TriangleValidator struct {
	UserInput UserInput
}

// Validate the input
func (v *TriangleValidator) Validate() bool {
	if !allSidesArePositive(v.UserInput) {
		return false
	}
	if !twoSidesAreGreaterThanTheThird(v.UserInput) {
		return false
	}
	return true
}

func allSidesArePositive(u UserInput) bool {
	if u.SideA < 0 || u.SideB < 0 || u.SideC < 0 {
		return false
	}
	return true
}

func twoSidesAreGreaterThanTheThird(u UserInput) bool {
	if u.SideA > u.SideB-u.SideC && u.SideB > u.SideA-u.SideC && u.SideB > u.SideC-u.SideA {
		return true
	}
	return false
}
