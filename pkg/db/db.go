package db

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/NpoolPlatform/fox-proxy/pkg/db/ent"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/go-service-framework/pkg/mysql"
)

func Client() (*ent.Client, error) {
	return client()
}

func client() (*ent.Client, error) {
	conn, err := mysql.GetConn()
	if err != nil {
		return nil, err
	}
	drv := entsql.OpenDB(dialect.MySQL, conn)
	return ent.NewClient(ent.Driver(drv)), nil
}

func Init() error {
	c, err := client()
	if err != nil {
		return err
	}
	return c.
		Schema.Create(context.Background())
}

func WithTx(ctx context.Context, fn func(ctx context.Context, tx *ent.Tx) error) error {
	cli, err := Client()
	if err != nil {
		return err
	}
	tx, err := cli.Tx(ctx)
	if err != nil {
		return fmt.Errorf("fail get client transaction: %v", err)
	}

	succ := false
	defer func() {
		if !succ {
			err := tx.Rollback()
			if err != nil {
				logger.Sugar().Errorf("fail rollback: %v", err)
				return
			}
		}
	}()

	if err := fn(ctx, tx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("committing transaction: %v", err)
	}

	succ = true
	return nil
}

func WithClient(ctx context.Context, fn func(ctx context.Context, cli *ent.Client) error) error {
	cli, err := Client()
	if err != nil {
		return fmt.Errorf("fail get db client: %v", err)
	}

	if err := fn(ctx, cli); err != nil {
		return err
	}
	return nil
}
