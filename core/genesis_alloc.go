
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
//版权所有2017 Go Ethereum作者
//此文件是Go以太坊库的一部分。
//
//Go-Ethereum库是免费软件：您可以重新分发它和/或修改
//根据GNU发布的较低通用公共许可证的条款
//自由软件基金会，或者许可证的第3版，或者
//（由您选择）任何更高版本。
//
//Go以太坊图书馆的发行目的是希望它会有用，
//但没有任何保证；甚至没有
//适销性或特定用途的适用性。见
//GNU较低的通用公共许可证，了解更多详细信息。
//
//你应该收到一份GNU较低级别的公共许可证副本
//以及Go以太坊图书馆。如果没有，请参见<http://www.gnu.org/licenses/>。

package core

//包含内置Genesis块的Genesis分配的常量。
//它们的内容是由RLP编码的（地址、平衡）元组列表。
//使用mkalloc.go创建/更新它们。

//诺林：拼写错误
/*
const testnetAllocData = "\xf9\x03\xa4\u0080\x01\xc2\x01\x01\xc2\x02\x01\xc2\x03\x01\xc2\x04\x01\xc2\x05\x01\xc2\x06\x01\xc2\a\x01\xc2\b\x01\xc2\t\x01\xc2\n\x80\xc2\v\x80\xc2\f\x80\xc2\r\x80\xc2\x0e\x80\xc2\x0f\x80\xc2\x10\x80\xc2\x11\x80\xc2\x12\x80\xc2\x13\x80\xc2\x14\x80\xc2\x15\x80\xc2\x16\x80\xc2\x17\x80\xc2\x18\x80\xc2\x19\x80\xc2\x1a\x80\xc2 \x1b \x80 \xc2 \x1c \x80 \xc2 \x1d \x80 \xc2 \x1e \x80 \xc2 \x1f \x80 \xc2 \x80 \xc2！\x80\xc2\"\x80\xc2#\x80\xc2$\x80\xc2%\x80\xc2&\x80\xc2'\x80\xc2(\x80\xc2)\x80\xc2*\x80\xc2+\x80\xc2,\x80\xc2-\x80\xc2.\x80\xc2/\x80\xc20\x80\xc21\x80\xc22\x80\xc23\x80\xc24\x80\xc25\x80\xc26\x80\xc27\x80\xc28\x80\xc29\x80\xc2:\x80\xc2;\x80\xc2<\x80\xc2=\x80\xc2>\x80\xc2?\x80\xc2@\x80\xc2A\x80\xc2B\x80\xc2C\x80\xc2D\x80\xc2E\x80\xc2F\x80\xc2G\x80\xc2H\x80\xc2I\x80\xc2J\x80\xc2K\x80\xc2L\x80\xc2M\x80\xc2N\x80\xc2O\x80\xc2P\x80\xc2Q\x80\xc2R\x80\xc2S\x80\xc2T\x80\xc2U\x80\xc2V\x80\xc2W\x80\xc2X\x80\xc2Y\x80\xc2Z\x80\xc2[\x80\xc2\\\x80\xc2]\x80\xc2^\x80\xc2_\x80\xc2`\x80\xc2a\x80\xc2b\x80\xc2c\x80\xc2d\x80\xc2e\x80\xc2f\x80\xc2g\x80\xc2h\x80\xc2i\x80\xc2j\x80\xc2k\x80\xc2l\x80\xc2m\x80\xc2n\x80\xc2o\x80\xc2p\x80\xc2q\x80\xc2r\x80\xc2s\x80\xc2t\x80\xc2u\x80\xc2v\x80\xc2w\x80\xc2x\x80\xc2y\x80\xc2z\x80\xc2{\x80\xc2|\x80\xc2}\x80\xc2~\x80\xc2\u007f\x80\u00c1\x80\x80\u00c1\x81\x80\u00c1\x82\x80\u00c1\x83\x80\u00c1\x84\x80\u00c1\x85\x80\u00c1\x86\x80\u00c1\x87\x80\u00c1\x88\x80\u00c1\x89\x80\u00c1\x8a\x80\u00c1\x8b\x80\u00c1\x8c\x80\u00c1\x8d\x80\u00c1\x8e\x80\u00c1\x8f\x80\u00c1\x90\x80\u00c1\x91\x80\u00c1\x92\x80\u00c1\x93\x80\u00c1\x94\x80\u00c1\x95\x80\u00c1\x96\x80\u00c1\x97\x80\u00c1\x98\x80\u00c1\x99\x80\u00c1\x9a\x80\u00c1\x9b\x80\u00c1\x9c\x80\u00c1\x9d\x80\u00c1\x9e\x80\u00c1\x9f\x80\u00c1\xa0\x80\u00c1\xa1\x80\u00c1\xa2\x80\u00c1\xa3\x80\u00c1\xa4\x80\u00c1\xa5\x80\u00c1\xa6\x80\u00c1\xa7\x80\u00c1\xa8\x80\u00c1\xa9\x80\u00c1\xaa\x80\u00c1\xab\x80\u00c1\xac\x80\u00c1\xad\x80\u00c1\xae\x80\u00c1\xaf\x80\u00c1\xb0\x80\u00c1\xb1\x80\u00c1\xb2\x80\u00c1\xb3\x80\u00c1\xb4\x80\u00c1\xb5\x80\u00c1\xb6\x80\u00c1\xb7\x80\u00c1\xb8\x80\u00c1\xb9\x80\u00c1\xba\x80\u00c1\xbb\x80\u00c1\xbc\x80\u00c1\xbd\x80\u00c1\xbe\x80\u00c1\xbf\x80\u00c1\xc0\x80\u00c1\xc1\x80\u00c1\u0080\u00c1\u00c0\u00c1\u0100\u00c1\u0140\u00c1\u0180\u00c1\u01c0\u00c1\u0200\u00c1\u0240\u00c1\u0280\u00c1\u02c0\u00c1\u0300\u00c1\u0340\u00c1\u0380\u00c1\u03c0\u00c1\u0400\u00c1\u0440\u00c1\u0480\u00c1\u04c0\u00c1\u0500\u00c1\u0540\u00c1\u0580\u00c1\u05c0\u00c1\u0600\u00c1\u0640\u00c1\u0680\u00c1\u06c0\u00c1\u0700\u00c1\u0740\u00c1\u0780\u00c1\u07c0\u00c1\xe0\x80\u00c1\xe1\x80\u00c1\xe2\x80\u00c1\xe3\x80\u00c1\xe4\x80\u00c1\xe5\x80\u00c1\xe6\x80\u00c1\xe7\x80\u00c1\xe8\x80\u00c1\xe9\x80\u00c1\xea\x80\u00c1\xeb\x80\u00c1\xec\x80\u00c1\xed\x80\u00c1\xee\x80\u00c1\xef\x80\u00c1\xf0\x80\u00c1\xf1\x80\u00c1\xf2\x80\u00c1\xf3\x80\u00c1\xf4\x80\u00c1\xf5\x80\u00c1\xf6\u00c1\u00c1\u00c1\xf0\xf0\x80\u00c1\xf8\x80\u00c1\xf9\x80\u00c1\x80\u00c1\xfc\x80\u00c1\x80\u00c1\xf80\xfffd\x80\x80\x80\x80\xfd\x80\x80\x80\xu00c1 \xfe \x80 \u00c1 \xff \x80 \u3507kt \xa8 \xbd \x15）f \xd6？pk\xae\x1f\xfe\xb0a\x19！\xe5 \x8d \f \x9f，\x9c \xd0ft \xed \xea@\x00 \x00 \x00”
const rinkebyAllocData = "\xf9\x03\xb7\u0080\x01\xc2\x01\x01\xc2\x02\x01\xc2\x03\x01\xc2\x04\x01\xc2\x05\x01\xc2\x06\x01\xc2\a\x01\xc2\b\x01\xc2\t\x01\xc2\n\x01\xc2\v\x01\xc2\f\x01\xc2\r\x01\xc2\x0e\x01\xc2\x0f\x01\xc2\x10\x01\xc2\x11\x01\xc2\x12\x01\xc2\x13\x01\xc2\x14\x01\xc2\x15\x01\xc2\x16\x01\xc2\x17\x01\xc2\x18\x01\xc2\x19\x01\xc2\x1a\x01\xc2 \x1b \x01 \xc2 \x1c \x01 \xc2 \x1d \x01 \xc2 \x1e \x01 \xc2 \x1f \x01 \xc2 \x01 \xc2！\x01\xc2\"\x01\xc2#\x01\xc2$\x01\xc2%\x01\xc2&\x01\xc2'\x01\xc2(\x01\xc2)\x01\xc2*\x01\xc2+\x01\xc2,\x01\xc2-\x01\xc2.\x01\xc2/\x01\xc20\x01\xc21\x01\xc22\x01\xc23\x01\xc24\x01\xc25\x01\xc26\x01\xc27\x01\xc28\x01\xc29\x01\xc2:\x01\xc2;\x01\xc2<\x01\xc2=\x01\xc2>\x01\xc2?\x01\xc2@\x01\xc2A\x01\xc2B\x01\xc2C\x01\xc2D\x01\xc2E\x01\xc2F\x01\xc2G\x01\xc2H\x01\xc2I\x01\xc2J\x01\xc2K\x01\xc2L\x01\xc2M\x01\xc2N\x01\xc2O\x01\xc2P\x01\xc2Q\x01\xc2R\x01\xc2S\x01\xc2T\x01\xc2U\x01\xc2V\x01\xc2W\x01\xc2X\x01\xc2Y\x01\xc2Z\x01\xc2[\x01\xc2\\\x01\xc2]\x01\xc2^\x01\xc2_\x01\xc2`\x01\xc2a\x01\xc2b\x01\xc2c\x01\xc2d\x01\xc2e\x01\xc2f\x01\xc2g\x01\xc2h\x01\xc2i\x01\xc2j\x01\xc2k\x01\xc2l\x01\xc2m\x01\xc2n\x01\xc2o\x01\xc2p\x01\xc2q\x01\xc2r\x01\xc2s\x01\xc2t\x01\xc2u\x01\xc2v\x01\xc2w\x01\xc2x\x01\xc2y\x01\xc2z\x01\xc2{\x01\xc2|\x01\xc2}\x01\xc2~\x01\xc2\u007f\x01\u00c1\x80\x01\u00c1\x81\x01\u00c1\x82\x01\u00c1\x83\x01\u00c1\x84\x01\u00c1\x85\x01\u00c1\x86\x01\u00c1\x87\x01\u00c1\x88\x01\u00c1\x89\x01\u00c1\x8a\x01\u00c1\x8b\x01\u00c1\x8c\x01\u00c1\x8d\x01\u00c1\x8e\x01\u00c1\x8f\x01\u00c1\x90\x01\u00c1\x91\x01\u00c1\x92\x01\u00c1\x93\x01\u00c1\x94\x01\u00c1\x95\x01\u00c1\x96\x01\u00c1\x97\x01\u00c1\x98\x01\u00c1\x99\x01\u00c1\x9a\x01\u00c1\x9b\x01\u00c1\x9c\x01\u00c1\x9d\x01\u00c1\x9e\x01\u00c1\x9f\x01\u00c1\xa0\x01\u00c1\xa1\x01\u00c1\xa2\x01\u00c1\xa3\x01\u00c1\xa4\x01\u00c1\xa5\x01\u00c1\xa6\x01\u00c1\xa7\x01\u00c1\xa8\x01\u00c1\xa9\x01\u00c1\xaa\x01\u00c1\xab\x01\u00c1\xac\x01\u00c1\xad\x01\u00c1\xae\x01\u00c1\xaf\x01\u00c1\xb0\x01\u00c1\xb1\x01\u00c1\xb2\x01\u00c1\xb3\x01\u00c1\xb4\x01\u00c1\xb5\x01\u00c1\xb6\x01\u00c1\xb7\x01\u00c1\xb8\x01\u00c1\xb9\x01\u00c1\xba\x01\u00c1\xbb\x01\u00c1\xbc\x01\u00c1\xbd\x01\u00c1\xbe\x01\u00c1\xbf\x01\u00c1\xc0\x01\u00c1\xc1\x01\u00c1\xc2\x01\u00c1\xc3\x01\u00c1\xc4\x01\u00c1\xc5\x01\u00c1\xc6\x01\u00c1\xc7\x01\u00c1\xc8\x01\u00c1\xc9\x01\u00c1\xca\x01\u00c1\xcb\x01\u00c1\xcc\x01\u00c1\xcd\x01\u00c1\xce\x01\u00c1\xcf\x01\u00c1\xd0\x01\u00c1\xd1\x01\u00c1\xd2\x01\u00c1\xd3\x01\u00c1\xd4\x01\u00c1\xd5\x01\u00c1\xd6\x01\u00c1\xd7\x01\u00c1\xd8\x01\u00c1\xd9\x01\u00c1\xda\x01\u00c1\xdb\x01\u00c1\xdc\x01\u00c1\xdd\x01\u00c1\xde\x01\u00c1\xdf\x01\u00c1\xe0\x01\u00c1\xe1\x01\u00c1\xe2\x01\u00c1\xe3\x01\u00c1\xe4\x01\u00c1\xe5\x01\u00c1\xe6\x01\u00c1\xe7\x01\u00c1\xe8\x01\u00c1\xe9\x01\u00c1\xea\x01\u00c1\xeb\x01\u00c1\xec\x01\u00c1\xed\x01\u00c1\xee\x01\u00c1\xef\x01\u00c1\xf0\x01\u00c1\xf1\x01\u00c1\xf2\x01\u00c1\xf3\x01\u00c1\xf4\x01\u00c1\xf5\x01\u00c1\xf6\x01\u00c1\xf7\x01\u00c1\xf8\x01\u00c1\xf9\x01\u00c1\xfa\x01\u00c1\xfb\x01\u00c1\xfc\x01\u00c1\xfd\x01\u00c1\xfe\x01\u00c1\xff\x01\xf6\x941\xb9\x8d\x14\x00{\xde\xe67)\x80\x86\x98\x8a\v\xbd1\x18E#\xa0 \x02 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00 \x00”
