# ipfs-driver

This is a Ipfs driver for Glacier to integrate Ipfs's feature! It implements Glacier's Standard Storage interface.

# Example

- Run a local IPFS node

```
ipfs init
ipfs daemon
```

- Test Driver

```
go test -v
```

- List files 

```
ipfs files ls /glc001
hello.txt
```

- Stat file

```
ipfs files stat /glc001/hello.txt
QmZ8wL7UgEocrxJmom8ABbamwU9G3XEi2PaUTZ1HHWTfS9
Size: 11
CumulativeSize: 121
ChildBlocks: 2
Type: file
```

- Cat file

```
ipfs cat QmZ8wL7UgEocrxJmom8ABbamwU9G3XEi2PaUTZ1HHWTfS9
hello ipfs!
```

- Create a Filcoin storage deal

```
lotus client deal QmZ8wL7UgEocrxJmom8ABbamwU9G3XEi2PaUTZ1HHWTfS9 t01000 <price> <duration>
```



# Ref

- Gateway: https://ipfs.github.io/public-gateway-checker/
- Client: https://docs.ipfs.tech/reference/kubo-rpc-cli/
- Guide: https://github.com/ipfs/kubo/blob/master/client/rpc/README.md
- API: https://docs.ipfs.tech/reference/kubo/rpc/#api-v0-files-stat
- Lotus: https://lotus.filecoin.io/tutorials/lotus/import-data-from-ipfs/