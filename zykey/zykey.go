package zykey

import (
    "net/http"
    "time"
    "io/ioutil"
    "math/rand"
    "io"
    "os"
    "strconv"
)



func GetKey(url string ) (int64, error) {
	
    rand.Seed(time.Now().Unix())
    tmp_file := ".li_" + strconv.Itoa(rand.Intn(1000000))
    out, err := os.Create(tmp_file)
    if err != nil {
        return 0, err
    }
    defer os.Remove(tmp_file)
    defer out.Close()

    res, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer res.Body.Close()

    _, err = io.Copy(out, res.Body)
    if err != nil {
        return 0, err
    }
    contents, err := ioutil.ReadFile(tmp_file); 
    if err == nil{
        val, _  :=  strconv.ParseInt(string(contents),10,64)
        return val, nil
    } else {
    	return 0, err
    }
}

