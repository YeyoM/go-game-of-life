package helpers

func DetectChanges(prev Size, current Size) bool {
	if prev.Width != current.Width || prev.Height != current.Height {
		return true
	}

	return false
}
