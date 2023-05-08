package backup

import (
	v1 "github.com/rancher/backup-restore-operator/pkg/apis/resources.cattle.io/v1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestBackupFilename(t *testing.T) {
	// Expected output: TestName-TestNamespace-2023-05-08T13-40-33-04-00
	// The timestamp is based on time.Now() so we don't check for it explicitly
	namespace := "TestNamespace"
	backup_name := "TestName"

	mock_handler := handler{
		kubeSystemNS: namespace,
	}

	backup := &v1.Backup{}
	backup.SetName(backup_name)
	filename, err := mock_handler.generateBackupFilename(backup)
	print(filename + "\n")
	require.NoError(t, err, "Error when generating backup filename")
	wantNamespace := regexp.MustCompile(`\b` + namespace + `\b`)
	assert.True(t, wantNamespace.MatchString(filename), "Expected namespace in generated filename")
	wantName := regexp.MustCompile(`\b` + backup_name + `\b`)
	assert.True(t, wantName.MatchString(filename), "Expected backup name in generated filename")
}
