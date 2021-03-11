# yangon

Self-use Development scaffolding

![](https://img.shields.io/badge/windowns10-Development-d0d1d4)
![](https://img.shields.io/badge/golang-1.16-blue)
![](https://img.shields.io/badge/version-1.0.1-r)

## :rocket:Installation

```
go get -u github.com/coder2z/yangon
```

or

```
git clone https://github.com.cnpmjs.org/coder2z/yangon.git

cd yangon

go install
```

## :anchor:Usage

### Verify that the installation is complete

input:
```shell
λ yangon --help
```

output:
```console
Usage:
  Yangon [command]

Available Commands:
  go          db,handle,server,route code production
  help        Help about any command
  new         Generate app scaffolding
  version     app version

Flags:
  -h, --help   help for Yangon

Use "Yangon [command] --help" for more information about a command.
```

### Create project

input:

```shell
λ yangon new -a appname -p projectname

λ cd projectname
```

output:

```                                     
├─build                                   
│  └─appname       //   dockerfile                       
├─cmd              //   app main                       
│  └─appname                              
│      └─app                              
├─config           //   config                       
├─deployments      //   k8sfile                      
│  └─appname                              
│      └─templates                        
├─internal         //   business code                       
│  └─appname                              
│      ├─api                              
│      │  └─v1                            
│      │      ├─handle                    
│      │      ├─middleware                
│      │      └─registry                  
│      ├─map                              
│      ├─model                            
│      ├─services                         
│      └─validator                        
├─pkg              //   Public package                         
│  ├─constant                             
│  ├─rand                                 
│  ├─recaptcha                            
│  ├─response                             
│  └─rpc                                  
├─scripts          //   construct                          
│  └─appname                              
└─test                                    
```

###Generate CRUD code

modify config

```shell
vim config/config.toml
```

Set up database connection config

```shell
[mysql.main]
    tablePrefix = ""
    host = "127.0.0.1"
    username = "root"
    password = "root"
    dbName = ""
    type = "mysql"
    debug = true
```

input:

```shell
λ yangon go -a appname -p projectname -v v1
```

output:
```
├─internal
│  └─appname
│      ├─api
│      │  └─v1
│      │      ├─handle          //code Here     
│      │      ├─middleware      
│      │      └─registry        //code Here    
│      ├─map                    //code Here    
│      ├─model                  //code Here    
│      │  ├─access_token
│      │  └─user
│      ├─services               //code Here    
│      │  ├─access_token
│      │  └─user
│      └─validator
```

## :tada:Contribute code

Open source projects are inseparable from everyone’s support. If you have a good idea, encountered some bugs and fixed
them, and corrected the errors in the document, please submit a Pull Request~

1. Fork this project to your own repo
2. Clone the project in the past, that is, the project in your warehouse, to your local
3. Modify the code
4. Push to your own library after commit
5. Initiate a PR (pull request) request and submit it to the `provide` branch
6. Waiting to merge

## :closed_book:License

Distributed under MIT License, please see license file within the code for more details.