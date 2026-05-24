package fsIo

import (
	"zarinloosli.com/hangouts-wrapped/model"
	"zarinloosli.com/hangouts-wrapped/state/stats"
	"zarinloosli.com/hangouts-wrapped/util"
)

type DirectoryHandle struct {
	FSHandle
}

func (handle DirectoryHandle) Entries() []model.FSAgnosticHandle {
	fsEntries := []model.FSAgnosticHandle{}
	directoryHandle, err := handle.BrowserHandle.AsDirectoryHandle()
	util.PanicIfError(err)

	for _, browserEntry := range directoryHandle.Entries() {
		stats.IncrementStat(stats.FilesFound)
		PathToFSHandle[browserEntry.Path()] = FSHandle{browserEntry}
		fsEntries = append(fsEntries, FSHandle{browserEntry})
	}
	return fsEntries
}

func (handle DirectoryHandle) GetEntry(name string) (model.FSAgnosticHandle, error) {
	directoryHandle, err := handle.BrowserHandle.AsDirectoryHandle()
	if err != nil {
		return nil, err
	}

	entry, err := directoryHandle.GetEntry(name)
	stats.IncrementStat(stats.FilesFound)
	return DirectoryHandle{FSHandle{entry}}, err
}

var _ model.FSAgnosticDirectoryHandle = DirectoryHandle{} // Compile-time inheritance check
