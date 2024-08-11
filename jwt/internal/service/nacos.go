package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"sync"
)

var Nacos = nacosService{}

type nacosService struct{
	ctx      context.Context
	log      *glog.Logger
	scfg     []constant.ServerConfig
	ccfg     *constant.ClientConfig

	serviceName   string
	serviceHost   string
	servicePort   uint64

	impClient naming_client.INamingClient
}

func NewNacosService(ctx context.Context,
						glog *glog.Logger,
						scfg []constant.ServerConfig,
						ccfg constant.ClientConfig,
					) *nacosService {
	nacosSrv := &nacosService{
		ctx:         ctx,
		log:         glog,
		scfg:        scfg,
		ccfg:        &ccfg,

	}
	nacosSrv.impClient = nacosSrv.GetNacosClients(nacosSrv.ccfg,nacosSrv.scfg)

	if nacosSrv.impClient==nil{
		glog.Fatal(ctx,"创建nacosService对象失败")
	}
	return nacosSrv
}

func (n *nacosService) Register(serviceName string,serviceHost string, servicePort uint64){
	if len(n.scfg)<1 {
		n.scfg = []constant.ServerConfig{
			{
				IpAddr: "127.0.0.1",
				Port:   8848,
			},
		}
	}

	if n.ccfg ==nil {
		n.ccfg  = &constant.ClientConfig{
			NamespaceId: "",
		}
	}

	if n.impClient==nil {
		n.impClient = n.GetNacosClients(n.ccfg,n.scfg)
	}

	// Register the service
	success, err := n.impClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          serviceHost,
		Port:        servicePort,
		ServiceName: serviceName,

		GroupName:   "bankend",
		ClusterName: "bankend-a",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "changsha"},
	})

	n.log.Info(n.ctx,"info:",serviceHost,"-",servicePort,"-",serviceName)
	if err != nil {
		n.log.Fatal(n.ctx,"系统错误:", err)
	}

	if success {
		// Registration successful
		n.log.Info(n.ctx,"创建nacos服务注册成功")
	} else {
		// Registration failed
		n.log.Fatal(n.ctx,"创建nacos服务注册失败:", err)
	}

	// Shutdown the client when done
	//client.CloseClient()
}

func (n *nacosService) GetConfig(dataId string, group string) (string,error) {
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  n.ccfg,
			ServerConfigs: n.scfg,
		},
	)

	if err != nil {
		n.log.Fatalf(n.ctx,"创建配置管理客户端错误",err)
	}

	config, err := client.GetConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
	})

	if err != nil {
		n.log.Fatalf(n.ctx,"获取配置错误",err)
	}

	if config == "" {
		n.log.Fatalf(n.ctx,"获取配置为空")
	}

	//  热加载到goframe 配置文件中
	err = client.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetContent(config,"default.yaml")
			g.Cfg().GetAdapter().(*gcfg.AdapterFile).SetFileName("default.yaml")
		},
	})

	return config,err
}

func (n *nacosService) GetNacosClients(ccfg *constant.ClientConfig,scfg []constant.ServerConfig) naming_client.INamingClient{
	var (
		once sync.Once
		client  naming_client.INamingClient
	)

	once.Do(func() {
		client, _ = clients.NewNamingClient(
			vo.NacosClientParam{
				ClientConfig:  ccfg,
				ServerConfigs: scfg,
			},
		)
	})

	return client
}






