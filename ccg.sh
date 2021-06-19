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
export stream="ocid1.stream.oc1.eu-frankfurt-1.amaaaaaaxwhvexya6tvz4er3h4rkldezrvvw4ibtye5grk3ikwlng4k3wgsa"
export tenancy="ocid1.tenancy.oc1..aaaaaaaafwn2xifqqy2ces6zkuxa6terq4ftsm5c5hun3p7if3s7t2uj2vga"
export user="ocid1.user.oc1..aaaaaaaaxzdnwy2xc2sk7bkilul6w7gpwryxfjdxre42rzwg3gcpgmhjeeha"
export region="eu-frankfurt-1"
export fingerprint="85:b7:1f:8d:af:13:42:81:0b:ad:ea:7c:18:6d:f3:d6"
export ppkfile=$(cat ./myppk)
export ppkpasswd="-"
#
./ccg 
#