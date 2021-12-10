//go:build clickhouse
// +build clickhouse

package cli

import (
	_ "github.com/mailru/go-clickhouse"
	//_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/golang-migrate/migrate/v4/database/clickhouse"
)
