package cache

import "time"

type nowFuncT func() time.Time

// For testing purpose
var nowFunc nowFuncT
