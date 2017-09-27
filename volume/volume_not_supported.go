package volume

import (
	"github.com/libopenstorage/openstorage/api"
)

var (
	// BlockNotSupported is a default (null) block driver implementation.  This can be
	// used by drivers that do not want to (or care about) implementing the attach,
	// format and detach interfaces.
	BlockNotSupported = &blockNotSupported{}
	// SnapshotNotSupported is a null snapshot driver implementation. This can be used
	// by drivers that do not want to implement the snapshot interface
	SnapshotNotSupported = &snapshotNotSupported{}
	// IONotSupported is a null IODriver interface
	IONotSupported = &ioNotSupported{}
	// StatsNotSupported is a null stats driver implementation. This can be used
	// by drivers that do not want to implement the stats interface.
	StatsNotSupported = &statsNotSupported{}
	// BackupNotSupported is a null stats driver implementation. This can be used
	// by drivers that do not want to implement the backup interface.
	BackupNotSupported = &backupNotSupported{}
)

type blockNotSupported struct{}

func (b *blockNotSupported) Attach(volumeID string, attachOptions map[string]string) (string, error) {
	return "", ErrNotSupported
}

func (b *blockNotSupported) Detach(volumeID string, options map[string]string) error {
	return ErrNotSupported
}

type snapshotNotSupported struct{}

func (s *snapshotNotSupported) Snapshot(volumeID string, readonly bool, locator *api.VolumeLocator) (string, error) {
	return "", ErrNotSupported
}

func (s *snapshotNotSupported) Restore(volumeID, snapshotID string) error {
	return ErrNotSupported
}

type ioNotSupported struct{}

func (i *ioNotSupported) Read(volumeID string, buffer []byte, size uint64, offset int64) (int64, error) {
	return 0, ErrNotSupported
}

func (i *ioNotSupported) Write(volumeID string, buffer []byte, size uint64, offset int64) (int64, error) {
	return 0, ErrNotSupported
}

func (i *ioNotSupported) Flush(volumeID string) error {
	return ErrNotSupported
}

type statsNotSupported struct{}

// Stats returns stats
func (s *statsNotSupported) Stats(
	volumeID string,
	cumulative bool,
) (*api.Stats, error) {
	return nil, ErrNotSupported
}

// UsedSize returns allocated size
func (s *statsNotSupported) UsedSize(volumeID string) (uint64, error) {
	return 0, ErrNotSupported
}

// GetActiveRequests gets active requests
func (s *statsNotSupported) GetActiveRequests() (*api.ActiveRequests, error) {
	return nil, nil
}

type backupNotSupported struct{}

func (b *backupNotSupported) GetCloudBkupCatalog(cloudVol string, credID string) ([]byte, error)  {
	return nil, ErrNotSupported
}

func (b *backupNotSupported) GetCloudBkupMetadata(cloudVol string, credID string) (map[string]string, error) {
	return nil, ErrNotSupported
}

func (b *backupNotSupported) CloudBackup(volumeID string, snapID string, credID string, fullBkup bool, scheduled bool) error {
	return ErrNotSupported
}

func (b *backupNotSupported) CloudRestore(dstVol string, cloudVol string, credID string, nodeID string) (string, error) {
	return "", ErrNotSupported
}

func (b *backupNotSupported) ListCloudSnaps(srcVol string, clusterID string, credID string, all bool) ([]*api.CloudSnapInfo, error) {
	return nil, ErrNotSupported
}

func (b *backupNotSupported) DeleteCloudSnaps(clusterID string, volstring, credID string) error {
	return ErrNotSupported
}
func (b *backupNotSupported) CloudBackupStatusFromCache(volumeID string, local bool) (map[string]*api.CloudSnapStatus, error) {
	return nil, ErrNotSupported
}

func (b *backupNotSupported) ChangeStateForCloudBackup(volumeID string, reqState string) error {
	return ErrNotSupported
}

func (b *backupNotSupported) CreateCloudBackupSchedule(schedInfo api.CloudsnapScheduleInfo) (string, error) {
	return "", ErrNotSupported
}

func (b *backupNotSupported) UpdateCloudBackupSchedule(uuid string, schedInfo api.CloudsnapScheduleInfo) error {
	return ErrNotSupported
}

func (b *backupNotSupported) ListCloudBackupSchedules() (map[string]api.CloudsnapScheduleInfo, error) {
	return nil, ErrNotSupported
}

func (b *backupNotSupported) DeleteCloudBackupSchedule(uuid string) error {
	return ErrNotSupported
}