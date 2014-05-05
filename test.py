import urllib2, sys, zlib

proxy, dst = sys.argv[1:]

proxy_support = urllib2.ProxyHandler({"http":proxy})
opener = urllib2.build_opener(proxy_support)
urllib2.install_opener(opener)

request = urllib2.Request(dst)
request.add_header('hello', 'world')
request.add_header('Accept-Encoding', 'gzip, deflate');
response = urllib2.urlopen(request)
print response.info()
print zlib.decompress(response.read(), 16+zlib.MAX_WBITS);


