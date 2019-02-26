// Copyright 2018 kdevb0x Ltd. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause license
// The full license text can be found in the LICENSE file.

package passman

import (
	"log"

	gocb "gopkg.in/couchbase/gocb.v1"
)
type dbConn struct {
	gocb.Cluster
}
func dbConnect(db string) (dbConn, error) {
	cluster, err := gocb.Connect(db)
	if err != nil {
		log.Println(err)
		return err
	}
	defer cluster.Close
	return nil
}

func main() {

}
