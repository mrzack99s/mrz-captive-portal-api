# mrz-captive-portal-api

>	Author: Chatdanai Phakaket <br>
>	Email: zchatdanai@gmail.com 

MRZ-Captive-portal is distributed system captive portal and mrz-captive-portal-api is a one part of system

## How to deploy

1. Change an api environments value in **config.yaml** file
2. Create a namespace
```
    kubectl apply -f https://raw.githubusercontent.com/mrzack99s/mrz-captive-portal-api/master/namespace.yaml
```
3. Apply the mainfests 
```
    kubectl apply -f https://raw.githubusercontent.com/mrzack99s/mrz-captive-portal-api/master/mainfests.yaml
```
If you're want to change the port and API ip address. You can follow in below.

- Download mainfests.yaml
```
curl -fsSL https://raw.githubusercontent.com/mrzack99s/mrz-captive-portal-api/master/mainfests.yaml > mainfest.yaml
```

- Change a service value
```
spec:
  loadBalancerIP: < IP Address >
  ports:
  - port: <Port>
    protocol: TCP
```

If you're want to change the replica set. Follow this below
```
spec:
  replicas: < Replica set number >
```

## License

Copyright (c) 2020 - Chatdanai Phakaket

	

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)
