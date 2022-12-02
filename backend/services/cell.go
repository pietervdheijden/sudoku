package services

// TODO: consider renaming options to candidates
// TODO: consider creating type for options
type Cell struct {
	Id int
	Row int
	Column int
	Square int
	Number int
	Options []int
}

func (cell *Cell) removeOption(value int) bool {
	for i, option := range (*cell).Options {
		if option == value {
			(*cell).Options = append((*cell).Options[:i], (*cell).Options[i+1:]...)
			return true
		}
	}
    return false
}

func (cell *Cell) removeOptions(options []int) bool {
	removed := false
	for _, option := range options {
		r := cell.removeOption(option)
		if r {
			removed = true
		}
	}
	return removed
}