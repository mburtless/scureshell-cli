package errorHandler

import (
	"net"
	"net/url"
	"log"
	"os"
)

func Handle(err error) {
	if ue, ok := err.(*url.Error); ok {
		//handle connection refused error
		switch uet := ue.Err.(type) {
			//if oe, ok := ue.Err.(*net.OpError); ok {
		case *net.OpError:
				switch oet := uet.Err.(type) {
				case *os.SyscallError:
					//if se, ok := oe.Err.(*os.SyscallError); ok {
						if oet.Err.Error() == "connection refused" {
							log.Fatalf("Error: Connection refused when attempting to connect to scureshell server at %s", uet.Addr.(*net.TCPAddr))
						}
					//}
			}
		default:
			if ue.Op == "parse" {
				log.Fatalf("Error: Invalid URL provided for scureshell server - \"%s\"", ue.URL)
			}
		}
	}
}
