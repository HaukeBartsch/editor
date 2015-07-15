Simple web-based text editor for MagickBox bucket files
=========================================================

Editor is an application used to edit files on the (bucket-) local file system. The interface is
rendered as a web-site. The only files editable with this application currently are info.json, work.sh, and
db-plugin.code. Using a golang (statically-linked) executable no web-server like apache or nginx need to
be installed inside the bucket.

Build instructions
--------------------

go-bindata (https://github.com/jteeuwen/go-bindata.git) is used to package the web-site code into the
golang executable. The package needs to be installed (here into /bin/go-bindata) and executed before compiling successfully. 

```
  /bin/go-bindata views/...
  go build
```

Test the application by calling:

```
  ./editor open
```
which will open a web-server that listens to port 9090 on the local machine. Connect to it using http://localhost:9090.

<img src="img/screenshot-bucketeditor.png" alt="Screenshot of bucket editor">

