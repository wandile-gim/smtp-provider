package test

import (
	"github.com/wandile/smtp-provider/ent"
	"github.com/wandile/smtp-provider/ent/enttest"
	"github.com/wandile/smtp-provider/ent/migrate"
	"testing"
)

func DbConnection(t *testing.T) *ent.Client {
	opts := []enttest.Option{
		enttest.WithOptions(ent.Log(t.Log)),
		enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(false)),
	}
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1", opts...)

	return client
}
