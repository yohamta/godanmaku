// +build darwin,!arm,!arm64 freebsd linux windows js
// +build !android
// +build !ios

package touch

// IsTouchPrimaryInput returns if touch is primary input
func IsTouchPrimaryInput() bool {
	return false
}
