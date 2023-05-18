package abi

func EncodeMint(to string,tokenId string) string{
	formatter := new(Formatter)

	method := "mint(address,uint256)"
	methodSig := formatter.ToMethodFormat(method)
	addressFmt := formatter.ToAddressFormat(to)
	tokenIdFmt := formatter.ToIntegerFormat(tokenId, 64)

	input := methodSig + addressFmt + tokenIdFmt
	return input
}

func EncodeTransferFrom(from string,to string,tokenId string) string{
	formatter := new(Formatter)

	method := "transferFrom(address,address,uint256)"
	methodSig := formatter.ToMethodFormat(method)
	addressFromFmt := formatter.ToAddressFormat(from)
	addressToFmt := formatter.ToAddressFormat(to)
	tokenIdFmt := formatter.ToIntegerFormat(tokenId, 64)

	input := methodSig + addressFromFmt + addressToFmt+ tokenIdFmt
	return input
}

func EncodeBurn(tokenId string) string{
	formatter := new(Formatter)

	method := "burn(uint256)"
	methodSig := formatter.ToMethodFormat(method)
	tokenIdFmt := formatter.ToIntegerFormat(tokenId, 64)

	input := methodSig + tokenIdFmt
	return input
}

func EncodeOwnerOf(tokenId string)string{
	formatter := new(Formatter)

	method := "ownerOf(uint256)"
	methodSig := formatter.ToMethodFormat(method)
	tokenIdFmt := formatter.ToIntegerFormat(tokenId, 64)
	input := methodSig + tokenIdFmt
	return input
}


