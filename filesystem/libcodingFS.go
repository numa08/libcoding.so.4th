package filesystem
import (
    "github.com/hanwen/go-fuse/fuse/pathfs"
    "github.com/hanwen/go-fuse/fuse"
    "github.com/hanwen/go-fuse/fuse/nodefs"
    "github.com/numa08/libcoding.so.4th/libcoding"
)

type LibcodingFs struct {
    pathfs.FileSystem
    _performers []libcoding.Performer
}

func (this *LibcodingFs) performers() []libcoding.Performer {
    if this == nil {
        return nil
    }
    if this._performers != nil {
        return this._performers
    }
    this._performers, _ = libcoding.LoadPerformers()
    return this._performers
}

func (this *LibcodingFs) GetAttr(name string, context *fuse.Context) (*fuse.Attr, fuse.Status) {
    performers := this.performers()
    performer := search(performers, name)
    if performer == nil {
        return nil, fuse.ENOENT
    }
    return &fuse.Attr{
        Mode: fuse.S_IFREG | 0644, Size: uint64(len(name)),
    }, fuse.OK
}

func (this *LibcodingFs) OpenDir(name string, context *fuse.Context) (c []fuse.DirEntry, code fuse.Status) {
    performers := this.performers()

    if name == "" {
        entry := make([]fuse.DirEntry, len(performers))
        for i := 0; i < len(performers); i++ {
            e := fuse.DirEntry{
                Name: performers[i].Name,
                Mode: fuse.S_IFREG,
            }
            entry[i] = e
        }

        return entry, fuse.OK
    }

    return nil, fuse.ENOENT
}

func (this *LibcodingFs) Open(name string, flags uint32, context *fuse.Context) (file nodefs.File, code fuse.Status) {
    if flags&fuse.O_ANYWRITE != 0 {
        return nil, fuse.EPERM
    }
    performers := this.performers()
    performer := search(performers, name)
    if performer == nil {
        return nil, fuse.ENOENT
    }
    return nodefs.NewDataFile([]byte(performer.Name)), fuse.OK
}

func search(pf []libcoding.Performer, n string) *libcoding.Performer {
    for _, p := range pf {
        if p.Name == n {
            return &p
        }
    }

    return nil
}