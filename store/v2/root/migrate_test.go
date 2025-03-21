package root

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	corestore "cosmossdk.io/core/store"
	coretesting "cosmossdk.io/core/testing"
	"cosmossdk.io/log"
	"github.com/depinnetwork/depin-sdk/store/v2"
	"github.com/depinnetwork/depin-sdk/store/v2/commitment"
	"github.com/depinnetwork/depin-sdk/store/v2/commitment/iavl"
	dbm "github.com/depinnetwork/depin-sdk/store/v2/db"
	"github.com/depinnetwork/depin-sdk/store/v2/pruning"
)

var storeKeys = []string{"store1", "store2", "store3"}

type MigrateStoreTestSuite struct {
	suite.Suite

	rootStore store.RootStore
}

func TestMigrateStoreTestSuite(t *testing.T) {
	suite.Run(t, &MigrateStoreTestSuite{})
}

func (s *MigrateStoreTestSuite) SetupTest() {
	testLog := log.NewTestLogger(s.T())
	nopLog := coretesting.NewNopLogger()

	mdb := dbm.NewMemDB()
	multiTrees := make(map[string]commitment.Tree)
	for _, storeKey := range storeKeys {
		prefixDB := dbm.NewPrefixDB(mdb, []byte(storeKey))
		multiTrees[storeKey] = iavl.NewIavlTree(prefixDB, nopLog, iavl.DefaultConfig())
	}
	orgSC, err := commitment.NewCommitStore(multiTrees, nil, mdb, testLog)
	s.Require().NoError(err)

	// apply changeset against the original store
	toVersion := uint64(200)
	keyCount := 10
	for version := uint64(1); version <= toVersion; version++ {
		cs := corestore.NewChangeset(version)
		for _, storeKey := range storeKeys {
			for i := 0; i < keyCount; i++ {
				cs.Add([]byte(storeKey), []byte(fmt.Sprintf("key-%d-%d", version, i)), []byte(fmt.Sprintf("value-%d-%d", version, i)), false)
			}
		}
		s.Require().NoError(orgSC.WriteChangeset(cs))
		_, err = orgSC.Commit(version)
		s.Require().NoError(err)
	}

	multiTrees1 := make(map[string]commitment.Tree)
	for _, storeKey := range storeKeys {
		multiTrees1[storeKey] = iavl.NewIavlTree(dbm.NewMemDB(), nopLog, iavl.DefaultConfig())
	}
	sc, err := commitment.NewCommitStore(multiTrees1, nil, dbm.NewMemDB(), testLog)
	s.Require().NoError(err)

	pm := pruning.NewManager(sc, nil)

	// assume no storage store, simulate the migration process
	s.rootStore, err = New(dbm.NewMemDB(), testLog, orgSC, pm, nil)
	s.Require().NoError(err)
}

func (s *MigrateStoreTestSuite) TestMigrateState() {
	err := s.rootStore.LoadLatestVersion()
	s.Require().NoError(err)
	originalLatestVersion, err := s.rootStore.GetLatestVersion()
	s.Require().NoError(err)

	// check if the Query fallback to the original SC
	for version := uint64(1); version <= originalLatestVersion; version++ {
		for _, storeKey := range storeKeys {
			for i := 0; i < 10; i++ {
				res, err := s.rootStore.Query([]byte(storeKey), version, []byte(fmt.Sprintf("key-%d-%d", version, i)), true)
				s.Require().NoError(err)
				s.Require().Equal([]byte(fmt.Sprintf("value-%d-%d", version, i)), res.Value)
			}
		}
	}

	// continue to apply changeset against the original store
	latestVersion := originalLatestVersion + 1
	keyCount := 10
	for ; latestVersion < 2*originalLatestVersion; latestVersion++ {
		cs := corestore.NewChangeset(latestVersion)
		for _, storeKey := range storeKeys {
			for i := 0; i < keyCount; i++ {
				cs.Add([]byte(storeKey), []byte(fmt.Sprintf("key-%d-%d", latestVersion, i)), []byte(fmt.Sprintf("value-%d-%d", latestVersion, i)), false)
			}
		}
		_, err = s.rootStore.Commit(cs)
		s.Require().NoError(err)

		// check if the migration is completed
		ver, err := s.rootStore.GetLatestVersion()
		s.Require().NoError(err)
		if ver == latestVersion {
			break
		}

		// add some delay to simulate the consensus process
		time.Sleep(100 * time.Millisecond)
	}

	// check if the migration is successful
	version, err := s.rootStore.GetLatestVersion()
	s.Require().NoError(err)
	s.Require().Equal(latestVersion, version)

	// query against the migrated store
	for version := uint64(1); version <= latestVersion; version++ {
		for _, storeKey := range storeKeys {
			for i := 0; i < keyCount; i++ {
				targetVersion := version
				if version < originalLatestVersion {
					targetVersion = originalLatestVersion
				}
				res, err := s.rootStore.Query([]byte(storeKey), targetVersion, []byte(fmt.Sprintf("key-%d-%d", version, i)), true)
				s.Require().NoError(err)
				s.Require().Equal([]byte(fmt.Sprintf("value-%d-%d", version, i)), res.Value)
			}
		}
	}

	// apply changeset against the migrated store
	for version := latestVersion + 1; version <= latestVersion+10; version++ {
		cs := corestore.NewChangeset(version)
		for _, storeKey := range storeKeys {
			for i := 0; i < keyCount; i++ {
				cs.Add([]byte(storeKey), []byte(fmt.Sprintf("key-%d-%d", version, i)), []byte(fmt.Sprintf("value-%d-%d", version, i)), false)
			}
		}
		_, err = s.rootStore.Commit(cs)
		s.Require().NoError(err)
	}

	version, err = s.rootStore.GetLatestVersion()
	s.Require().NoError(err)
	s.Require().Equal(latestVersion+10, version)
}
