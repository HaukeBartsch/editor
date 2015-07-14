Simple web-based text editor for MagickBox bucket files
=========================================================

Editor is an application used to edit files on the (bucket-) local file system. The interface is
rendered as a web-site.

Build instructions
--------------------

go-bindata (https://github.com/jteeuwen/go-bindata.git) is used to package the web-site code into the
golang executable. The package needs to be installed to compile successfully. 

```
  /bin/go-bindata views/...
  go build
```

Test the application by calling:

```
  ./editor open
```
