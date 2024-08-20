package utils

func Page(currentPage int, perPageSize int) (offset, limit int) {
	return (currentPage - 1) * perPageSize, perPageSize
}
