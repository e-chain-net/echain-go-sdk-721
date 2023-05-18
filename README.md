# EChain-go-sdk-721

适配标准``ERC721``合约的 ``EChain-Server`` 签名SDK，提供标准``ERC721``合约中如下方法的签名方法：
```
function mint(address to, uint256 tokenId);
function transferFrom(address from,address to,uint256 tokenId); 
function burn(uint256 tokenId);
```
# 测试用例
见 ``test`` 包

|  示例名称   | 描述  |
|  ----  | ----  |
| account_test.go  | 生成随机账户地址、私钥 |
| sdk_test.go  | 对标准`Erc721`合约进行铸造、转移、销毁进行签名，得到交易哈希、签名后的交易体 |

# 与EChainServer的结合使用
见 https://github.com/e-chain-net/echain-server-go-demo