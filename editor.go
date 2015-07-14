package main

import "fmt"
import "os"
import "sync"
import "github.com/codegangsta/cli"

func p( t string) {
  fmt.Println("> ", t)
}

func main() {

     app := cli.NewApp()
     app.Name    = "editor"
     app.Usage   = "Web-based text editor for MagickBox bucket files.\n\n" +
                   "   This program opens an application reachable by\n" +
                   "   a local web-browser (localhost:9090). The interface\n" +
                   "   allows to edit MagickBox related text files."
     app.Version = "0.0.1"
     app.Author  = "Hauke Bartsch"
     app.Email   = "HaukeBartsch@gmail.com"
     app.Flags = []cli.Flag {
       cli.BoolFlag {
         Name:  "verbose",
         Usage: "Generate verbose output with intermediate files.",
       },
     }

     app.Commands = []cli.Command {
       {
         Name: "open",
         ShortName: "o",
         Usage: "Start the editor",
         Description: ".",
         Flags: []cli.Flag{
           cli.BoolFlag {
             Name: "init",
             Usage: "Initialize missing files",
           },
         },
         Action: func(c *cli.Context) {
             verbose     := c.GlobalBool("verbose")
             init_flag   := !c.Bool("init") 
             if (verbose) {
               p("verbose on")
               p("open editor")
             }
             var wg sync.WaitGroup
             wg.Add(1)
             go webinterface( verbose, init_flag, wg )
             wg.Wait()
         },
       },
     }
     app.Run(os.Args)
}
