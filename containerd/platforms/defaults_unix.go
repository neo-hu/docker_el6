// +build !windows

package platforms


// Default returns the default matcher for the platform.
func Default() string {
	return Format(DefaultSpec())
}

