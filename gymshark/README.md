### TEST WITH UI 
http://gym-shark-ui.s3-website-us-east-1.amazonaws.com/

<br/>

### TEST WITH CURL
### Calculate packs required
curl http://54.197.220.230/calculate-packs\?orderQuantity\={{totalOrderQty}}

### Update packs
curl -X POST -H "Content-Type: application/json" -d '{
"packSizes": [250,500,1000,2000,5000]
}' http://54.197.220.230/update-pack-sizes

<br/>

### LOCAL RUN
`cd package_service`

`go run .`


## General Info

#### UI is a simple html which is serving from aws s3 bucket

#### Package service is the backend and running in an AWS EC2 instance

#### I used one of my existing docker hub repo to push the image. https://hub.docker.com/repository/docker/emreisler/sample-gateway