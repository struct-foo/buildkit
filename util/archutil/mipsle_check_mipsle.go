//go:build !mipsle

package archutil

func mipsleSupported() (string, error) {
	return "", nil
}
