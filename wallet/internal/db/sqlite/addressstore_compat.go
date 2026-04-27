package sqlite

import (
	"context"

	db "github.com/btcsuite/btcwallet/wallet/internal/db"
)

// ImportAccount is not yet implemented for the SQLite store.
func (s *Store) ImportAccount(_ context.Context,
	_ db.ImportAccountParams) (*db.AccountProperties, error) {

	return nil, db.AccountManagerCompatNotImplemented("ImportAccount")
}
