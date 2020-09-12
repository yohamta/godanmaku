// +build android ios darwin,arm darwin,arm64 js

package touch

// IsTouchPrimaryInput returns if touch is primary input
func IsTouchPrimaryInput() bool {
	return true
}
