package main

import (
	"encoding/base64"
	"log"
	"io/ioutil"
	"net"

	"github.com/gorilla/mux"
	"github.com/gorilla/context"
	"github.com/gorilla/sessions"
	"github.com/gorilla/http"
)

type ConnHandler
