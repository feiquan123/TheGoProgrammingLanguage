package musicManager

import "testing"

func TestOps(t *testing.T) {
	mm := NewMusicManager()
	if mm == nil {
		t.Error("NewMusicManager failed.")
	}
	if mm.Len() != 0 {
		t.Error("NewMusicManager failed, not empty.")
	}
	m0 := &MusicEntry{
		Id:     "1",
		Name:   "My Heart Will Go On",
		Artist: "Celion Dion",
		Style:  "Pop",
		Source: "http://qbox.me/24501234",
		Type:   "MP3",
	}
	mm.Add(m0)

	if mm.Len() != 1 {
		t.Error("MusicManger.Add() failed.")
	}

	m := mm.Find(m0.Name)
	if m == nil {
		t.Error("MusicManager.Find() failed.")
	}
	if m.Id != m0.Id || m.Name != m0.Name || m.Artist != m0.Artist || m.Style != m0.Style || m.Source != m0.Source || m.Type != m0.Type {
		t.Error("MusicManager.Find() failed. Found item mismatch.")
	}

	m, err := mm.Get(0)
	if m == nil {
		t.Error("MusicManager.Get() failed.", err)
	}

	m = mm.Remove(0)
	if mm.Len() != 0 || m == nil {
		t.Error("MusicManager.Remove() failed.", err)
	}
}
