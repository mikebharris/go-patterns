package main

type DoiChecker struct {
	ZipFilepath string
}

type Results struct {
	ThePath string
}

func (c *DoiChecker) GetDoiResults() (Results, error) {
	return Results{ThePath: c.ZipFilepath}, nil
}
