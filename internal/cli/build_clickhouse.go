//go:build clickhouse

package cli

import (
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/henryspk/migrate/v4/database/clickhouse"
)
