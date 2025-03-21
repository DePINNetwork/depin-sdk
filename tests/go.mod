module github.com/cosmos/cosmos-sdk/tests

go 1.23.5

require (
	cosmossdk.io/api v0.8.2
	cosmossdk.io/collections v1.1.0
	cosmossdk.io/core v1.0.0
	cosmossdk.io/core/testing v0.0.2
	cosmossdk.io/depinject v1.1.0
	cosmossdk.io/log v1.5.0
	cosmossdk.io/math v1.5.0
	cosmossdk.io/store v1.10.0-rc.1
	cosmossdk.io/x/tx v1.1.0
	github.com/cometbft/cometbft v1.0.1 // indirect
	github.com/cosmos/cosmos-proto v1.0.0-beta.5
	github.com/cosmos/gogoproto v1.7.0
	github.com/spf13/cobra v1.8.1 // indirect
	github.com/stretchr/testify v1.10.0
	go.uber.org/mock v0.5.0
	google.golang.org/grpc v1.71.0
	google.golang.org/protobuf v1.36.4
	gotest.tools/v3 v3.5.1
	pgregory.net/rapid v1.1.0
)

require (
	cosmossdk.io/runtime/v2 v2.0.0-20240911143651-72620a577660
	cosmossdk.io/server/v2/stf v1.0.0-beta.2
	cosmossdk.io/store/v2 v2.0.0-beta.1
	github.com/cometbft/cometbft/api v1.0.0 // indirect
	github.com/cosmos/cosmos-db v1.1.1
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/google/go-cmp v0.6.0
	github.com/google/gofuzz v1.2.0
	gitlab.com/yawning/secp256k1-voi v0.0.0-20230925100816-f2616030848b
)

require (
	cosmossdk.io/x/accounts v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/accounts/defaults/base v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/accounts/defaults/lockup v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/accounts/defaults/multisig v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/authz v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/bank v0.0.0-20240226161501-23359a0b6d91
	cosmossdk.io/x/consensus v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/distribution v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/evidence v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/feegrant v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/gov v0.0.0-20231113122742-912390d5fc4a
	cosmossdk.io/x/group v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/mint v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/protocolpool v0.0.0-20230925135524-a1bc045b3190
	cosmossdk.io/x/slashing v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/staking v0.0.0-00010101000000-000000000000
	cosmossdk.io/x/upgrade v0.0.0-00010101000000-000000000000
)

require (
	buf.build/gen/go/cometbft/cometbft/protocolbuffers/go v1.36.4-20241120201313-68e42a58b301.1 // indirect
	buf.build/gen/go/cosmos/gogo-proto/protocolbuffers/go v1.36.4-20240130113600-88ef6483f90f.1 // indirect
	cloud.google.com/go v0.118.0 // indirect
	cloud.google.com/go/auth v0.13.0 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.6 // indirect
	cloud.google.com/go/compute/metadata v0.6.0 // indirect
	cloud.google.com/go/iam v1.3.1 // indirect
	cloud.google.com/go/storage v1.43.0 // indirect
	cosmossdk.io/client/v2 v2.0.0-beta.6 // indirect
	cosmossdk.io/errors v1.0.1 // indirect
	cosmossdk.io/errors/v2 v2.0.0 // indirect
	cosmossdk.io/schema v1.0.0 // indirect
	cosmossdk.io/server/v2/appmanager v1.0.0-beta.2 // indirect
	cosmossdk.io/x/epochs v0.0.0-20240522060652-a1ae4c3e0337 // indirect
	github.com/DataDog/zstd v1.5.6 // indirect
	github.com/aws/aws-sdk-go v1.55.5 // indirect
	github.com/aybabtme/uniplot v0.0.0-20151203143629-039c559e5e7e // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/go-netrc v0.0.0-20140422174119-9fd32a8b3d3d // indirect
	github.com/bits-and-blooms/bitset v1.10.0 // indirect
	github.com/bvinc/go-sqlite-lite v0.6.1 // indirect
	github.com/bytedance/sonic v1.12.8 // indirect
	github.com/bytedance/sonic/loader v0.2.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/cloudwego/base64x v0.1.5 // indirect
	github.com/cockroachdb/apd/v3 v3.2.1 // indirect
	github.com/cockroachdb/errors v1.11.3 // indirect
	github.com/cockroachdb/fifo v0.0.0-20240816210425-c5d0cb0b6fc0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20241215232642-bb51bb14a506 // indirect
	github.com/cockroachdb/pebble v1.1.2 // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/cockroachdb/tokenbucket v0.0.0-20230807174530-cc333fc44b06 // indirect
	github.com/cosmos/iavl v1.3.4 // indirect
	github.com/cosmos/iavl/v2 v2.0.0-alpha.4 // indirect
	github.com/cosmos/ics23/go v0.11.0 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/emicklei/dot v1.6.2 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/getsentry/sentry-go v0.30.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.3 // indirect
	github.com/google/s2a-go v0.1.9 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.4 // indirect
	github.com/googleapis/gax-go/v2 v2.14.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.16.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-getter v1.7.8 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-metrics v0.5.4 // indirect
	github.com/hashicorp/go-safetemp v1.0.0 // indirect
	github.com/hashicorp/go-version v1.7.0 // indirect
	github.com/hashicorp/golang-lru v1.0.2 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/kocubinski/costor-api v1.1.1 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/linxGnu/grocksdb v1.9.7 // indirect
	github.com/manifoldco/promptui v0.9.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-testing-interface v1.14.1 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_golang v1.20.5 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.62.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/rogpeppe/go-internal v1.13.1 // indirect
	github.com/rs/zerolog v1.33.0 // indirect
	github.com/spf13/cast v1.7.1 // indirect
	github.com/spf13/pflag v1.0.6 // indirect
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d // indirect
	github.com/tendermint/go-amino v0.16.0 // indirect
	github.com/tidwall/btree v1.7.0 // indirect
	github.com/twitchyliquid64/golang-asm v0.15.1 // indirect
	github.com/ulikunitz/xz v0.5.12 // indirect
	gitlab.com/yawning/tuplehash v0.0.0-20230713102510-df83abbf9a02 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.58.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.58.0 // indirect
	go.opentelemetry.io/otel v1.34.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	golang.org/x/arch v0.13.0 // indirect
	golang.org/x/crypto v0.32.0 // indirect
	golang.org/x/exp v0.0.0-20250106191152-7588d65b2ba8 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/oauth2 v0.25.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	golang.org/x/time v0.9.0 // indirect
	google.golang.org/api v0.216.0 // indirect
	google.golang.org/genproto v0.0.0-20250106144421-5f5ef82da422 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20250106144421-5f5ef82da422 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250122153221-138b5a5a4fd4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)

// Here are the short-lived replace from the SimApp
// Replace here are pending PRs, or version to be tagged
// replace (
// 	<temporary replace>
// )

// SimApp on main always tests the latest extracted SDK modules importing the sdk
replace (
	cosmossdk.io/client/v2 => ../client/v2
	cosmossdk.io/indexer/postgres => ../indexer/postgres
	cosmossdk.io/runtime/v2 => ../runtime/v2
	cosmossdk.io/server/v2/appmanager => ../server/v2/appmanager
	cosmossdk.io/server/v2/stf => ../server/v2/stf
	cosmossdk.io/store/v2 => ../store/v2
	cosmossdk.io/tools/benchmark => ../tools/benchmark
	cosmossdk.io/x/accounts => ../x/accounts
	cosmossdk.io/x/accounts/defaults/base => ../x/accounts/defaults/base
	cosmossdk.io/x/accounts/defaults/lockup => ../x/accounts/defaults/lockup
	cosmossdk.io/x/accounts/defaults/multisig => ../x/accounts/defaults/multisig
	cosmossdk.io/x/authz => ../x/authz
	cosmossdk.io/x/bank => ../x/bank
	cosmossdk.io/x/circuit => ../x/circuit
	cosmossdk.io/x/consensus => ../x/consensus
	cosmossdk.io/x/distribution => ../x/distribution
	cosmossdk.io/x/epochs => ../x/epochs
	cosmossdk.io/x/evidence => ../x/evidence
	cosmossdk.io/x/feegrant => ../x/feegrant
	cosmossdk.io/x/gov => ../x/gov
	cosmossdk.io/x/group => ../x/group
	cosmossdk.io/x/mint => ../x/mint
	cosmossdk.io/x/nft => ../x/nft
	cosmossdk.io/x/protocolpool => ../x/protocolpool
	cosmossdk.io/x/slashing => ../x/slashing
	cosmossdk.io/x/staking => ../x/staking
	cosmossdk.io/x/upgrade => ../x/upgrade
)

// Below are the long-lived replace for tests.
replace (
	cosmossdk.io/core/testing => ../core/testing
	github.com/99designs/keyring => github.com/cosmos/keyring v1.2.0
	// We always want to test against the latest version of the SDK.
	github.com/cosmos/cosmos-sdk => ../.
)
