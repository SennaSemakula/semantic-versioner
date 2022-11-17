package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewVersion(t *testing.T) {
	v, err := NewVersion("v1.2.3")
	require.NotNil(t, v)
	require.NoError(t, err)

	cases := map[string]struct {
		input   string
		wantErr bool
	}{
		"version string missing": {
			input:   "",
			wantErr: true,
		},
		"version string valid": {
			input:   "v1.0.0",
			wantErr: false,
		},
		"version prefix missing": {
			input:   "1.0.0",
			wantErr: true,
		},
		"major version missing": {
			input:   "v1.0",
			wantErr: true,
		},
		"patch version missing": {
			input:   "v1.10",
			wantErr: true,
		},
		"empty major version": {
			input:   "v2..12",
			wantErr: true,
		},
		"empty minor version": {
			input:   "v2..12",
			wantErr: true,
		},
		"empty patch version": {
			input:   "v2.2..",
			wantErr: true,
		},
		"missing all versions but prefix": {
			input:   "v",
			wantErr: true,
		},
		"4 different versions": {
			input:   "v1.2.3.4",
			wantErr: true,
		},
	}
	for desc, tc := range cases {
		t.Run(desc, func(t *testing.T) {
			_, err := NewVersion(tc.input)
			t.Log(err)
			assert.Equal(t, tc.wantErr, err != nil)
		})
	}
}

func TestVersionString(t *testing.T) {
	v, err := NewVersion("v1.2.3")
	require.NotNil(t, v)
	require.NoError(t, err)
	assert.Equal(t, "v1.2.3", v.String())
}

func TestMajor(t *testing.T) {
	v, err := NewVersion("v1.2.3")
	require.NotNil(t, v)
	require.NoError(t, err)
	assert.Equal(t, 1, v.Major())
	assert.Equal(t, "1", v.MajorString())
}

func TestMinor(t *testing.T) {
	v, err := NewVersion("v1.2.3")
	require.NotNil(t, v)
	require.NoError(t, err)
	assert.Equal(t, 2, v.Minor())
	assert.Equal(t, "2", v.MinorString())
}

func TestPatch(t *testing.T) {
	v, err := NewVersion("v1.2.3")
	require.NotNil(t, v)
	require.NoError(t, err)
	assert.Equal(t, 3, v.Patch())
	assert.Equal(t, "3", v.PatchString())
}

func TestSemVersionString(t *testing.T) {
	assert.Equal(t, semversion(0).String(), "major")
	assert.Equal(t, semversion(1).String(), "minor")
	assert.Equal(t, semversion(2).String(), "patch")
}
