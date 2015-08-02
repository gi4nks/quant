package quant

import (
    "bytes"
    "compress/gzip"
	"io/ioutil"
)

func Compress(value string) (string, error) {
    var b bytes.Buffer
    gz := gzip.NewWriter(&b)
	
	defer gz.Flush()
	defer gz.Close()
    
	if _, err := gz.Write([]byte(value)); err != nil {
        return "", err
    }
	/*
    if err := gz.Flush(); err != nil {
        return "", err
    }
    if err := gz.Close(); err != nil {
        return "", err
    }
	*/
    return b.String(), nil
}


func Uncompress(value string) (string, error) {
	buf := bytes.NewBufferString(value)
	
	gz, err := gzip.NewReader(buf)
	
	if err != nil {
        return "", err
    }
	
	defer gz.Close()
	
	b, err := ioutil.ReadAll(gz)
    if err != nil {
    	return "", err
    }
	
    return string(b), nil
}