package utils

import (
    "RemindMe/global"
    "crypto/aes"
    "crypto/cipher"
    "crypto/sha1"
    "encoding/base64"
    "encoding/binary"
    "encoding/hex"
    "errors"
    "sort"
)

var (
    _token       = "bgUsotxSHXiGc1q3"
    _aesKey      = "XoHdGOrIC2nmMf2nDfxdwwdJSTFIQqzGpKar7EetNek"
    errSignature = errors.New("signature error")
)

func CheckSignature(signature, timestamp, nonce, msg string) bool {
    var ss = []string{_token, timestamp, nonce, msg}
    sort.Strings(ss)

    var sortedString string
    for _, s := range ss {
        sortedString += s
    }

    var sum = sha1.Sum([]byte(sortedString))
    var sign = hex.EncodeToString(sum[:])

    return signature == sign
}

func PKCS5UnPadding(origData []byte) []byte {
    length := len(origData)
    unpadding := int(origData[length-1])
    return origData[:(length - unpadding)]
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }

    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
    origData := make([]byte, len(crypted))
    blockMode.CryptBlocks(origData, crypted)
    origData = PKCS5UnPadding(origData)
    return origData, nil
}

func decryptEventData(msg string) (rawMsg, received string, err error) {
    aesMsg, err := base64.StdEncoding.DecodeString(msg)
    if err != nil {
        global.Log.Sugar().Error(err)
        return "", "", err
    }

    aesKey, err := base64.StdEncoding.DecodeString(_aesKey + "=")
    if err != nil {
        global.Log.Sugar().Error(err)
        return "", "", err
    }

    randMsg, err := AesDecrypt(aesMsg, aesKey)
    if err != nil {
        global.Log.Sugar().Error(err)
        return "", "", err
    }

    content := randMsg[16:]
    length := int(binary.BigEndian.Uint32(content[0:4]))

    rawMsg = string(content[4 : length+4])
    received = string(content[length+4:])

    return rawMsg, received, nil
}

func DecryptWeComEventMsg(sign, timestamp, nonce, data string) (msg, received string, err error) {
    if !CheckSignature(sign, timestamp, nonce, data) {
        return "", "", errSignature
    }

    msg, received, err = decryptEventData(data)
    if err != nil {
       return "", "", err
    }
    return
}
