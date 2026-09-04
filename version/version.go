// Package version описывает версию с тремя числовыми компонентами.
package version

import (
	"cmp"
	"errors"
	"fmt"
)

// Version хранит основную, дополнительную и исправляющую части версии.
// Предрелизные метки и метаданные сборки в эту модель не входят.
type Version struct {
	Major uint32
	Minor uint32
	Patch uint32
}

var ErrOverflow = errors.New("version component overflow")

func (v Version) Validate() error          { return nil }
func (v Version) String() string           { return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch) }
func (v Version) IsZero() bool             { return v == Version{} }
func (v Version) Equal(other Version) bool { return v == other }

// Compare orders versions lexicographically by major, minor, then patch.
func (v Version) Compare(other Version) (int, error) {
	if result := cmp.Compare(v.Major, other.Major); result != 0 {
		return result, nil
	}
	if result := cmp.Compare(v.Minor, other.Minor); result != 0 {
		return result, nil
	}
	return cmp.Compare(v.Patch, other.Patch), nil
}
func (v Version) NextPatch() (Version, error) {
	if v.Patch == ^uint32(0) {
		return Version{}, ErrOverflow
	}
	v.Patch++
	return v, nil
}

// NextMinor increments minor and resets patch.
func (v Version) NextMinor() (Version, error) {
	if v.Minor == ^uint32(0) {
		return Version{}, ErrOverflow
	}
	v.Minor++
	v.Patch = 0
	return v, nil
}

// NextMajor increments major and resets minor and patch.
func (v Version) NextMajor() (Version, error) {
	if v.Major == ^uint32(0) {
		return Version{}, ErrOverflow
	}
	v.Major++
	v.Minor = 0
	v.Patch = 0
	return v, nil
}
