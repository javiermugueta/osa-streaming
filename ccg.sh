#!/bin/sh
# jmu, jun 2021
#
# packages
#
go get github.com/oracle/oci-go-sdk/common github.com/oracle/oci-go-sdk/streaming github.com/oracle/oci-go-sdk/common github.com/eclipse/paho.mqtt.golang github.com/google/uuid io/ioutil github.com/oracle/oci-go-sdk/identity context fmt io/ioutil os github.com/oracle/oci-go-sdk/common time math/rand github.com/ShiraazMoollatjie/goluhn github.com/icrowley/fake
#
# build
#
go build ccg.go
#
# run
#
export stream="ocid1.stream.oc1.eu-frankfurt-1.amaa...wgsa"
export tenancy="ocid1.tenancy.oc1..aaaaa...vga"
export user="ocid1.user.oc1..aaaaa...ha"
export region="eu-frankfurt-1"
export fingerprint="85:...:d6"
export ppkfile=$(cat ./ppk)
export ppkpasswd="-"
#
./ccg 
#