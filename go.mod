go 1.23.5

module github.com/depinnetwork/depin-sdk

require (
	cosmossdk.io/api v0.8.2
	cosmossdk.io/collections v1.1.0
	cosmossdk.io/core v1.0.0
	cosmossdk.io/core/testing v0.0.2
	cosmossdk.io/depinject v1.1.0
	cosmossdk.io/errors v1.0.1
	cosmossdk.io/log v1.5.0
	cosmossdk.io/math v1.5.0
	cosmossdk.io/schema v1.0.0
	cosmossdk.io/x/bank v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/staking v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/tx v1.1.0
	filippo.io/edwards25519 v1.1.0
	github.com/99designs/keyring v1.2.2
	github.com/bgentry/speakeasy v0.2.0
	github.com/cometbft/cometbft/api v1.0.0
	github.com/cosmos/cosmos-db v1.1.1
	github.com/cosmos/cosmos-proto v1.0.0-beta.5
	github.com/cosmos/gogoproto v1.7.0
	github.com/cosmos/ledger-cosmos-go v0.14.0
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.0
	github.com/depinnetwork/bip39 v0.0.0-00010101000000-000000000000
	github.com/depinnetwork/btcutil v0.0.0-00010101000000-000000000000
	github.com/depinnetwork/depin-sdk/store v0.0.0-rc.1
	github.com/depinnetwork/por-consensus v0.0.0-20250322023815-76d0ebe3db59-20250322020112-8b685c2363ae-20250321114042-ad1fc897e07e-20250321113431-d202622238c6-20250321112806-4b08bd2e2b2b-20250321112100-653b8b4dd1bc-20250321110407-c5e8a59ea541-20250321110407-c5e8a59ea541
	github.com/depinnetwork/por-consensus/api/cometbft/abci v0.0.0-00010101000000-000000000000
	github.com/depinnetwork/por-consensus/api/cometbft/crypto v0.0.0-00010101000000-000000000000
	github.com/depinnetwork/por-consensus/api/cometbft/p2p v0.0.0-00010101000000-000000000000
	github.com/depinnetwork/por-consensus/api/cometbft/types v0.0.0-00010101000000-000000000000
	github.com/depinnetwork/por-consensus/api/cometbft/version v0.0.0-00010101000000-000000000000
	github.com/golang/protobuf v1.5.4
	github.com/google/go-cmp v0.7.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hashicorp/go-metrics v0.5.4
	github.com/hashicorp/golang-lru v1.0.2
	github.com/hdevalence/ed25519consensus v0.2.0
	github.com/huandu/skiplist v1.2.1
	github.com/magiconair/properties v1.8.9
	github.com/mattn/go-isatty v0.0.20
	github.com/mdp/qrterminal/v3 v3.2.0
	github.com/muesli/termenv v0.15.2
	github.com/prometheus/client_golang v1.21.1
	github.com/prometheus/common v0.62.0
	github.com/spf13/cast v1.7.1
	github.com/spf13/cobra v1.9.1
	github.com/spf13/pflag v1.0.6
	github.com/spf13/viper v1.19.0
	github.com/stretchr/testify v1.10.0
	github.com/tendermint/go-amino v0.16.0
	gitlab.com/yawning/secp256k1-voi v0.0.0-20230925100816-f2616030848b
	go.uber.org/mock v0.5.0
	golang.org/x/crypto v0.36.0
	google.golang.org/genproto/googleapis/api v0.0.0-20250106144421-5f5ef82da422
	google.golang.org/grpc v1.71.0
	google.golang.org/protobuf v1.36.5
	gotest.tools/v3 v3.5.1
	pgregory.net/rapid v1.1.0
	sigs.k8s.io/yaml v1.4.0
)

require (
	buf.build/gen/go/cometbft/cometbft/protocolbuffers/go v1.36.4-20241120201313-68e42a58b301.1 // indirect
	buf.build/gen/go/cosmos/gogo-proto/protocolbuffers/go v1.36.4-20240130113600-88ef6483f90f.1 // indirect
	github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4 // indirect
	github.com/DataDog/datadog-go v4.8.3+incompatible // indirect
	github.com/DataDog/zstd v1.5.6 // indirect
	github.com/Microsoft/go-winio v0.6.2 // indirect
	github.com/aymanbagabas/go-osc52/v2 v2.0.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bytedance/sonic v1.12.8 // indirect
	github.com/bytedance/sonic/loader v0.2.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/cloudwego/base64x v0.1.5 // indirect
	github.com/cockroachdb/apd/v3 v3.2.1 // indirect
	github.com/cockroachdb/errors v1.11.3 // indirect
	github.com/cockroachdb/fifo v0.0.0-20240816210425-c5d0cb0b6fc0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20241215232642-bb51bb14a506 // indirect
	github.com/cockroachdb/pebble v1.1.4 // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/cockroachdb/tokenbucket v0.0.0-20230807174530-cc333fc44b06 // indirect
	github.com/cosmos/ics23/go v0.11.0 // indirect
	github.com/danieljoos/wincred v1.2.1 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/depinnetwork/iavl v0.0.0-00010101000000-000000000000 // indirect
	github.com/depinnetwork/por-consensus/api/cometbft/consensus v0.0.0-00010101000000-000000000000 // indirect
	github.com/depinnetwork/por-consensus/api/cometbft/libs/bits v0.0.0-00010101000000-000000000000 // indirect
	github.com/depinnetwork/por-consensus/api/cometbft/mempool v0.0.0-00010101000000-000000000000 // indirect
	github.com/depinnetwork/por-consensus/api/cometbft/privval v0.0.0-00010101000000-000000000000 // indirect
	github.com/depinnetwork/por-consensus/api/cometbft/state v0.0.0-00010101000000-000000000000 // indirect
	github.com/depinnetwork/por-consensus/api/cometbft/store v0.0.0-00010101000000-000000000000 // indirect
	github.com/dvsekhvalnov/jose2go v1.6.0 // indirect
	github.com/emicklei/dot v1.6.4 // indirect
	github.com/ethereum/go-ethereum v1.15.5 // indirect
	github.com/fatih/color v1.18.0 // indirect
	github.com/fsnotify/fsnotify v1.8.0 // indirect
	github.com/getsentry/sentry-go v0.31.1 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/snappy v0.0.5-0.20220116011046-fa5810519dcb // indirect
	github.com/google/btree v1.1.3 // indirect
	github.com/google/orderedcode v0.0.1 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/hashicorp/go-hclog v1.6.3 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-plugin v1.6.3 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/yamux v0.1.2 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/linxGnu/grocksdb v1.9.7 // indirect
	github.com/lmittmann/tint v1.0.7 // indirect
	github.com/lucasb-eyer/go-colorful v1.2.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-runewidth v0.0.15 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/oasisprotocol/curve25519-voi v0.0.0-20230904125328-1f23a7beb09a // indirect
	github.com/oklog/run v1.1.0 // indirect
	github.com/pelletier/go-toml/v2 v2.2.3 // indirect
	github.com/petermattis/goid v0.0.0-20240813172612-4fcff4a6cae7 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/sagikazarmark/locafero v0.4.0 // indirect
	github.com/sagikazarmark/slog-shim v0.1.0 // indirect
	github.com/sasha-s/go-deadlock v0.3.5 // indirect
	github.com/sourcegraph/conc v0.3.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	github.com/supranational/blst v0.3.14 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d // indirect
	github.com/tidwall/btree v1.7.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/zondax/hid v0.9.2 // indirect
	github.com/zondax/ledger-go v0.14.3 // indirect
	gitlab.com/yawning/tuplehash v0.0.0-20230713102510-df83abbf9a02 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/arch v0.13.0 // indirect
	golang.org/x/exp v0.0.0-20250106191152-7588d65b2ba8 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/term v0.30.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250122153221-138b5a5a4fd4 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	rsc.io/qr v0.2.0 // indirect
)

// Here are the short-lived replace from the Cosmos SDK
// Replace here are pending PRs, or version to be tagged
// replace (
// 	<temporary replace>
// )

// TODO remove after all modules have their own go.mods
replace (
	cosmossdk.io/x/bank => ./x/bank
	cosmossdk.io/x/staking => ./x/staking
)

// Below are the long-lived replace of the Cosmos SDK
replace (
	// use cosmos fork of keyring
	github.com/99designs/keyring => github.com/cosmos/keyring v1.2.0

	// replace broken goleveldb
	github.com/syndtr/goleveldb => github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
)

retract (
	// false start by tagging the wrong branch
	v0.50.0
	// revert fix https://github.com/cosmos/cosmos-sdk/pull/16331
	v0.46.12
	// subject to a bug in the group module and gov module migration
	[v0.46.5, v0.46.6]
	// subject to the dragonberry vulnerability
	// and/or the bank coin metadata migration issue
	[v0.46.0, v0.46.4]
	// subject to the dragonberry vulnerability
	[v0.45.0, v0.45.8]
	// do not use
	v0.43.0
)

replace github.com/depinnetwork/por-consensus => ../por-consensus

replace github.com/depinnetwork/iavl => ../iavl

replace github.com/depinnetwork/btcutil => ../btcutil

replace github.com/depinnetwork/bip39 => ../bip39

replace github.com/depinnetwork/db => ../db

replace github.com/depinnetwork/por-consensus/api/cometbft/abci => ../por-consensus/api/cometbft/abci

replace github.com/depinnetwork/por-consensus/api/cometbft/blocksync => ../por-consensus/api/cometbft/blocksync

replace github.com/depinnetwork/por-consensus/api/cometbft/consensus => ../por-consensus/api/cometbft/consensus

replace github.com/depinnetwork/por-consensus/api/cometbft/crypto => ../por-consensus/api/cometbft/crypto

replace github.com/depinnetwork/por-consensus/api/cometbft/libs/bits => ../por-consensus/api/cometbft/libs/bits

replace github.com/depinnetwork/por-consensus/api/cometbft/mempool => ../por-consensus/api/cometbft/mempool

replace github.com/depinnetwork/por-consensus/api/cometbft/p2p => ../por-consensus/api/cometbft/p2p

replace github.com/depinnetwork/por-consensus/api/cometbft/privval => ../por-consensus/api/cometbft/privval

replace github.com/depinnetwork/por-consensus/api/cometbft/services/block => ../por-consensus/api/cometbft/services/block

replace github.com/depinnetwork/por-consensus/api/cometbft/services/block_results => ../por-consensus/api/cometbft/services/block_results

replace github.com/depinnetwork/por-consensus/api/cometbft/services/pruning => ../por-consensus/api/cometbft/services/pruning

replace github.com/depinnetwork/por-consensus/api/cometbft/services/version => ../por-consensus/api/cometbft/services/version

replace github.com/depinnetwork/por-consensus/api/cometbft/state => ../por-consensus/api/cometbft/state

replace github.com/depinnetwork/por-consensus/api/cometbft/statesync => ../por-consensus/api/cometbft/statesync

replace github.com/depinnetwork/por-consensus/api/cometbft/store => ../por-consensus/api/cometbft/store

replace github.com/depinnetwork/por-consensus/api/cometbft/types => ../por-consensus/api/cometbft/types

replace github.com/depinnetwork/por-consensus/api/cometbft/version => ../por-consensus/api/cometbft/version

replace github.com/depinnetwork/depin-sdk/store => ./store
