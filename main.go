package main
import (
    "flag"
    "log"
    "github.com/hanwen/go-fuse/fuse/pathfs"
    "github.com/hanwen/go-fuse/fuse/nodefs"
    "github.com/numa08/libcoding.so.4th/filesystem"
)

func main() {
    flag.Parse()
    if len(flag.Args()) < 1 {
        log.Fatal("Usage:\n libcoding4 MOUNTPOINT")
    }
    nfs := pathfs.NewPathNodeFs(&filesystem.LibcodingFs{FileSystem: pathfs.NewDefaultFileSystem()}, nil)
    server, _, err := nodefs.MountRoot(flag.Arg(0), nfs.Root(), nil)
    if err != nil {
        log.Fatalf("Mount fail: %v\n", err)
    }
    server.Serve()
}