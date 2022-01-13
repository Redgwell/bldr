package builtin

import (
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestDockerBuildSetConfig(t *testing.T) {
	targetName := "name-of-target"

	params := map[string]interface{}{
		"path":    "src/ui/Dockerfile",
		"include": []string{"src/ui/**"},
	}

	p := &DockerBuild{}

	p.SetConfig(nil, targetName, nil, params)

	assert.Equal(t, targetName, p.Name)
	assert.Equal(t, params["path"], p.Path)
	assert.Equal(t, params["include"], p.Include)
}

func TestDockerBuildIsAffectedByDiff(t *testing.T) {
	logger, _ := test.NewNullLogger()
	p := &DockerBuild{
		Include: []string{"src/ui/**", "src/other/file.txt"},
		logger:  logger,
	}

	diffFilePaths := []string{"src/ui/something.go"}
	isAffected := p.isAffectedByDiff(diffFilePaths)
	assert.True(t, isAffected)

	diffFilePaths = []string{"src/other/file.txt"}
	isAffected = p.isAffectedByDiff(diffFilePaths)
	assert.True(t, isAffected)

	diffFilePaths = []string{"src/other/another-file.txt", "src/stuff/things.go"}
	isAffected = p.isAffectedByDiff(diffFilePaths)
	assert.False(t, isAffected)
}
