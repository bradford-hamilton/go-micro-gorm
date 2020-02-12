package json

import jsoniter "github.com/json-iterator/go"

// API gets set to jsoniter with it's "ConfigCompatibleWithStandardLibrary" configuration. It is several
// times faster then encoding/json and is a noconf fully compatible (with encoding/json) API
var API jsoniter.API = jsoniter.ConfigCompatibleWithStandardLibrary
