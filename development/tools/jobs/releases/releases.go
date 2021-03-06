// Code generated by rendertemplates. DO NOT EDIT. 

package releases

// List of currently supported releases
var (
	Release113 = mustParse("1.13")
	Release112 = mustParse("1.12")
	Release111 = mustParse("1.11")
	Release110 = mustParse("1.10")
)

// GetAllKymaReleaseBranches returns all supported kyma release branches
func GetAllKymaReleases() []*SupportedRelease {
	return []*SupportedRelease{
		Release112,
		Release111,
		Release110,
	}
}

// GetNextKymaRelease returns the version of kyma currently under development
func GetNextKymaRelease() *SupportedRelease {
	return Release113
}


