package include

import "flag"

func init() {
	ConfigPath = flag.String("c", "", "configuration path")
	flag.Parse()
	LoadConfig()
	SharedHTTPClient = CreateHTTPClient()
}
