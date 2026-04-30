package kvdb

import (
	"context"
	"errors"
	"fmt"
	"iter"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcwallet/wallet/internal/db"
	"github.com/btcsuite/btcwallet/wallet/internal/db/page"
)

var errNoAddressInPkScript = errors.New("pkScript has no address")

// NewDerivedAddress is not yet implemented for kvdb.
func (s *Store) NewDerivedAddress(ctx context.Context,
	_ db.NewDerivedAddressParams,
	_ db.AddressDerivationFunc) (*db.AddressInfo, error) {

	return nil, notImplemented(ctx, "NewDerivedAddress")
}

// NewImportedAddress is not yet implemented for kvdb.
func (s *Store) NewImportedAddress(ctx context.Context,
	_ db.NewImportedAddressParams) (*db.AddressInfo, error) {

	return nil, notImplemented(ctx, "NewImportedAddress")
}

// GetAddress is not yet implemented for kvdb.
func (s *Store) GetAddress(ctx context.Context,
	_ db.GetAddressQuery) (*db.AddressInfo, error) {

	return nil, notImplemented(ctx, "GetAddress")
}

// ListAddresses is not yet implemented for kvdb.
func (s *Store) ListAddresses(ctx context.Context,
	_ db.ListAddressesQuery) (page.Result[db.AddressInfo, uint32], error) {

	return page.Result[db.AddressInfo, uint32]{},
		notImplemented(ctx, "ListAddresses")
}

// IterAddresses is not yet implemented for kvdb.
func (s *Store) IterAddresses(ctx context.Context,
	_ db.ListAddressesQuery) iter.Seq2[db.AddressInfo, error] {

	return func(yield func(db.AddressInfo, error) bool) {
		var zero db.AddressInfo

		yield(zero, notImplemented(ctx, "IterAddresses"))
	}
}

// GetAddressSecret is not yet implemented for kvdb.
func (s *Store) GetAddressSecret(ctx context.Context,
	_ db.GetAddressSecretQuery) (*db.AddressSecret, error) {

	return nil, notImplemented(ctx, "GetAddressSecret")
}

// ListAddressTypes is not yet implemented for kvdb.
func (s *Store) ListAddressTypes(ctx context.Context) ([]db.AddressTypeInfo,
	error) {

	return nil, notImplemented(ctx, "ListAddressTypes")
}

// GetAddressType is not yet implemented for kvdb.
func (s *Store) GetAddressType(ctx context.Context,
	_ db.AddressType) (db.AddressTypeInfo, error) {

	return db.AddressTypeInfo{}, notImplemented(ctx, "GetAddressType")
}

// kvdbAddressFromPkScript extracts the first standard address from one script.
//
// This lets the legacy address manager resolve wallet metadata from a
// script-pubkey-only store lookup.
//
//nolint:ireturn
func kvdbAddressFromPkScript(pkScript []byte,
	chainParams *chaincfg.Params) (btcutil.Address, error) {

	_, addrs, _, err := txscript.ExtractPkScriptAddrs(pkScript, chainParams)
	if err != nil {
		return nil, fmt.Errorf("extract address from pkScript: %w", err)
	}

	if len(addrs) == 0 {
		return nil, fmt.Errorf(
			"extract address from pkScript: %w", errNoAddressInPkScript,
		)
	}

	return addrs[0], nil
}
