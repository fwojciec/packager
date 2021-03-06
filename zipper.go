package packager

type zipper struct{}

func (z *zipper) Zip(buildDir string) (string, error) {
	return "", nil
}

func newZipper() Zipper {
	return &zipper{}
}
