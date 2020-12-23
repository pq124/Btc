package bccommand

import (
	"BtWeb/utils"
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
)
const  VERSION  =0x00

func GenerateKey(curve elliptic.Curve)(*ecdsa.PrivateKey,error)  {

	return ecdsa.GenerateKey(curve,rand.Reader)
}
//获取非压缩格式的公钥
func GetUnCompresspub(curve elliptic.Curve,pri *ecdsa.PrivateKey)[]byte  {
	return elliptic.Marshal(curve,pri.X,pri.Y)
}

//生成一个比特币的地址
func GetAddress() string {
	//1.SHA256
	curve:=elliptic.P256()
	pri,_:=GenerateKey(curve)
	pub:=GetUnCompresspub(curve,pri)

	hash256:=SHA256Hash(pub)

	ripemd160:= Ripemd160Hash(hash256)

	versionRipemd := append([]byte{VERSION},ripemd160...)

	//双hash
	hash1:=SHA256Hash(versionRipemd)
	hash2:=SHA256Hash(hash1)

	check:=hash2[0:4]

	add:=append(versionRipemd,check...)
	return utils.Encode(add)

}
//进行ripemd160Hash计算
func Ripemd160Hash(msg []byte)[]byte  {
	ripemd:=ripemd160.New()
	ripemd.Write(msg)
	return ripemd.Sum(nil)
}

//Sha256Hash计算
func SHA256Hash(msg []byte) []byte {
	sha256Hash:=sha256.New()
	sha256Hash.Write(msg)
	return sha256Hash.Sum(nil)

}

//校验给定的比特币的地址是否正确
func CheckAdd(add string) bool {
	//反编码
	deAddBytes := utils.Decode(add)
	//截取校验位
	length := len(deAddBytes)

	deCheck := deAddBytes[length-4:]
	//计算校验位
	//过去反编码去除后4位内容
	versionRipemd160 := deAddBytes[:length-4]

	//双hash
	sha256Hash := sha256.New()
	sha256Hash.Write(versionRipemd160)
	hash1 := sha256Hash.Sum(nil)

	sha256Hash.Reset()
	sha256Hash.Write(hash1)
	hash2 := sha256Hash.Sum(nil)

	check := hash2[:4]

	return bytes.Compare(deCheck, check) == 0

}