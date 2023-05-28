package safety

import (
	"fmt"
)

func SafeBootstrap(bootstrap func()) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("recovered from panic : %v\n", r)
				}
			}()
			bootstrap()
		}()
	}
}
