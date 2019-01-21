
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//版权所有2018 Go Ethereum作者
//此文件是Go以太坊的一部分。
//
//Go以太坊是免费软件：您可以重新发布和/或修改它
//根据GNU通用公共许可证的条款
//自由软件基金会，或者许可证的第3版，或者
//（由您选择）任何更高版本。
//
//Go以太坊的分布希望它会有用，
//但没有任何保证；甚至没有
//适销性或特定用途的适用性。见
//GNU通用公共许可证了解更多详细信息。
//
//你应该已经收到一份GNU通用公共许可证的副本
//一起去以太坊吧。如果没有，请参见<http://www.gnu.org/licenses/>。

package main

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/node"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/swarm/fuse"
	"gopkg.in/urfave/cli.v1"
)

func mount(cliContext *cli.Context) {
	args := cliContext.Args()
	if len(args) < 2 {
		utils.Fatalf("Usage: swarm fs mount --ipcpath <path to bzzd.ipc> <manifestHash> <file name>")
	}

	client, err := dialRPC(cliContext)
	if err != nil {
		utils.Fatalf("had an error dailing to RPC endpoint: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mf := &fuse.MountInfo{}
	mountPoint, err := filepath.Abs(filepath.Clean(args[1]))
	if err != nil {
		utils.Fatalf("error expanding path for mount point: %v", err)
	}
	err = client.CallContext(ctx, mf, "swarmfs_mount", args[0], mountPoint)
	if err != nil {
		utils.Fatalf("had an error calling the RPC endpoint while mounting: %v", err)
	}
}

func unmount(cliContext *cli.Context) {
	args := cliContext.Args()

	if len(args) < 1 {
		utils.Fatalf("Usage: swarm fs unmount --ipcpath <path to bzzd.ipc> <mount path>")
	}
	client, err := dialRPC(cliContext)
	if err != nil {
		utils.Fatalf("had an error dailing to RPC endpoint: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mf := fuse.MountInfo{}
	err = client.CallContext(ctx, &mf, "swarmfs_unmount", args[0])
	if err != nil {
		utils.Fatalf("encountered an error calling the RPC endpoint while unmounting: %v", err)
	}
fmt.Printf("%s\n", mf.LatestManifest) //
}

func listMounts(cliContext *cli.Context) {
	client, err := dialRPC(cliContext)
	if err != nil {
		utils.Fatalf("had an error dailing to RPC endpoint: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mf := []fuse.MountInfo{}
	err = client.CallContext(ctx, &mf, "swarmfs_listmounts")
	if err != nil {
		utils.Fatalf("encountered an error calling the RPC endpoint while listing mounts: %v", err)
	}
	if len(mf) == 0 {
		fmt.Print("Could not found any swarmfs mounts. Please make sure you've specified the correct RPC endpoint\n")
	} else {
		fmt.Printf("Found %d swarmfs mount(s):\n", len(mf))
		for i, mountInfo := range mf {
			fmt.Printf("%d:\n", i)
			fmt.Printf("\tMount point: %s\n", mountInfo.MountPoint)
			fmt.Printf("\tLatest Manifest: %s\n", mountInfo.LatestManifest)
			fmt.Printf("\tStart Manifest: %s\n", mountInfo.StartManifest)
		}
	}
}

func dialRPC(ctx *cli.Context) (*rpc.Client, error) {
	var endpoint string

	if ctx.IsSet(utils.IPCPathFlag.Name) {
		endpoint = ctx.String(utils.IPCPathFlag.Name)
	} else {
		utils.Fatalf("swarm ipc endpoint not specified")
	}

	if endpoint == "" {
		endpoint = node.DefaultIPCEndpoint(clientIdentifier)
	} else if strings.HasPrefix(endpoint, "rpc:") || strings.HasPrefix(endpoint, "ipc:") {
//与geth的向后兼容性<1.5，这需要
//这些前缀。
		endpoint = endpoint[4:]
	}
	return rpc.Dial(endpoint)
}
