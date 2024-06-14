package main

type DoiCheckerBuilder struct {
	zipFilepath string
}

func aDoiChecker() *DoiCheckerBuilder {
	return &DoiCheckerBuilder{}
}

func (b *DoiCheckerBuilder) withZipFilepath(zfp string) *DoiCheckerBuilder {
	b.zipFilepath = zfp
	return b
}

func (b *DoiCheckerBuilder) build() DoiChecker {
	return DoiChecker{ZipFilepath: b.zipFilepath}
}