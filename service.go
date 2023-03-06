package main

type service struct {
	// 端口
	Port int `default:"8080" json:"port"`
	// 协议
	Protocol string `default:"tcp" json:"protocol" validate:"oneof=tcp udp"`

	// 端口列表
	Ports []*port `json:"ports"`
}
