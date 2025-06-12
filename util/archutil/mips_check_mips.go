//go:build !mips

package archutil

func mipsSupported() (string, error) {
	return "", nil
}
