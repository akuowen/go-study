package web

import (
	"fmt"
	"time"
)

// FilterChainBuilder /*过滤器链构建器
type FilterChainBuilder func(filter Filter) Filter

// Filter /*
type Filter func(ctx *Context)

func MetricFilterBuilder(filter Filter) Filter {

	return func(ctx *Context) {
		startNano := time.Now().Nanosecond()
		filter(ctx)
		endNano := time.Now().Nanosecond()
		useNano := endNano - startNano
		fmt.Printf("run time: %d \n", useNano)
	}

}
