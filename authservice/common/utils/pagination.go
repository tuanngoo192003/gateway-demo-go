package utils

import (
    "math"
)

const (
    DEFAULT_PAGE int = 1 
    DEFAULT_SIZE int = 15 
)

func GetOffset(page int, perpage *int) int {
    tempPerpage := DEFAULT_SIZE
    if perpage != nil {
        tempPerpage = *perpage
    }
    return (page - 1) * tempPerpage
}

func GetTotalPage(totalRecord float32, perpage *int) int {
    tempPerpage := DEFAULT_SIZE
    if perpage != nil {
        tempPerpage = *perpage
    }
    return int(math.Ceil(float64(totalRecord) / float64(tempPerpage)))
}
