package main

import (
    "github.com/elazarl/goproxy"
    "log"
    "net/http"
    "io/ioutil"
    "bytes"
    "compress/gzip"
)

func NewResponse(r *http.Response) *http.Response {
    resp := &http.Response{}
    resp.Request = r.Request
    resp.TransferEncoding = []string{"gzip"}
    resp.Header = r.Header
    resp.StatusCode = r.StatusCode
    resp.Header.Set("Content-Encoding", "gzip")
    defer r.Body.Close()
    body, err := ioutil.ReadAll(r.Body)
    if err != nil {
        return nil
    }
    log.Println(r.Header)
    buf := new(bytes.Buffer)
    gzw, _ := gzip.NewWriterLevel(buf, 5)
    defer gzw.Close()
    gzw.Write(body)
    log.Println(len(body))
    resp.ContentLength = int64(buf.Len())
    resp.Body = ioutil.NopCloser(buf)
    return resp
}

func main() {
    proxy := goproxy.NewProxyHttpServer()
    proxy.Verbose = true
    proxy.OnRequest().DoFunc(
    func(r *http.Request,ctx *goproxy.ProxyCtx)(*http.Request,*http.Response) {
        if r.Header.Get("hello") != "world" {
           return r, goproxy.NewResponse(r, goproxy.ContentTypeText,http.StatusForbidden, "Don't waste your time!")
        }
        r.Header.Del("hello")
        return r,nil
    })
    proxy.OnResponse().DoFunc(
    func(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response{
       return NewResponse(resp)
    })
    log.Fatal(http.ListenAndServe(":8081", proxy))
}
