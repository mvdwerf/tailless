package tailless

func Parse(srcFilename, destFilename string) error {
	parser := newParser()
	return parser.Parse(srcFilename, destFilename)
}
