package context

import (
	"math"
	"sync"
)

func (c *C) Pagify(l []interface{}, j func(i interface{}, tx *Tx) error) *Tx {
	pageSize := appPageSize
	pageCount := int(math.Ceil(float64(len(l)) / float64(pageSize)))

	var wg sync.WaitGroup
	wg.Add(pageCount)

	tx := c.Start("generating pages")

	do := func(pageNumber int, page []interface{}) {
		tx.
			NewThread("processing page %v", pageNumber).
			Run(func(tx *Tx) {
				defer wg.Done()

				for _, x := range page {
					if err := j(x, tx); err != nil {
						tx.Start("Error during operation: %v", err)
					}
				}
			})
	}

	for i := 0; i < pageCount-1; i++ {
		do(i+1, l[i*pageSize:i*pageSize+pageSize])
	}

	do(pageCount, l[pageSize*(pageCount-1):])

	wg.Wait()

	return tx
}
