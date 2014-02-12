// Copyright 2009 Thomas Jager <mail@jager.no>  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package irc

import (
	"crypto/tls"
	"log"
	"net"
	"time"
)

type Nick struct {
	Nick string
	User string	
}

type Nicklist struct {
	Nick []Nick

}

//Initilize nicklist for channel
func (irccon *irc.Connection) Nicklist(channel string) irc.Nicklist {

}
