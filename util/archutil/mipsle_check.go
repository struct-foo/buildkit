//go:build !mipsle

package archutil

func mipsleSupported() (string, error) {
	return check("mipsle", Binarymipsle)
}
