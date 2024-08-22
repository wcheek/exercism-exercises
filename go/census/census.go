// Package census simulates a system used to collect census data.
package census

// Resident represents a resident in this city.
type Resident struct {
	Name    string
	Age     int
	Address map[string]string
}

// NewResident registers a new resident in this city.
func NewResident(name string, age int, address map[string]string) *Resident {
	return &Resident{Name: name, Age: age, Address: address}
}

// HasRequiredInfo determines if a given resident has all of the required information.
func (r *Resident) HasRequiredInfo() bool {
	nonZeroName := r.Name != ""
	nonZeroAddress := r.Address["street"] != ""
	return nonZeroName && nonZeroAddress
}

// Delete deletes a resident's information.
func (r *Resident) Delete() {
	r.Address = nil
	r.Name = ""
	r.Age = 0
}

// Count counts all residents that have provided the required information.
func Count(residents []*Resident) int {
	totalCount := 0
	for _, resident := range residents {
		if resident.HasRequiredInfo() {
			totalCount++
		}
	}
	return totalCount
}
