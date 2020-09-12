// +build darwin,!arm,!arm64 freebsd linux windows
// +build !android
// +build !ios
// +build !js

package touch

// IsTouchPrimaryInput returns if touch is primary input
func IsTouchPrimaryInput() bool {
	return false
}
