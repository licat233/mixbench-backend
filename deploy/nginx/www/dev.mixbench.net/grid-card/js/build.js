class App {
    list = [
        {
            name: "mixbench",
            description: "基于go-zero的微服务",
            picture: "./img/mixbench.jpg",
            link: "http://micro.mixbench.net"
        },
        {
            name: "phpRedisAdmin",
            description: "redis缓存管理",
            picture: "./img/redis.jpg",
            link: "http://redis.mixbench.net"
        },
        {
            name: "phpMyAdmin",
            description: "mysql数据库管理",
            picture: "./img/mysql.jpg",
            link: "http://mysql.mixbench.net"
        },
        {
            name: "filebrowser",
            description: "文件管理服务，用于浏览数据中心文件(ui)",
            picture: "./img/filebrowser.jpg",
            link: "http://filebrowser.mixbench.net"
        },
        {
            name: "jaeger",
            description: "go实现的分布式链路追踪系统(ui)",
            picture: "./img/jaeger.jpg",
            link: "http://jaeger.mixbench.net"
        },
        {
            name: "Prometheus",
            description: "Prometheus是一个系统和服务监控系统，它以给定的间隔从配置的目标收集度量，评估规则表达式，显示结果，并在观察到指定条件时触发警报。",
            picture: "./img/prometheus.jpg",
            link: "http://prometheus.mixbench.net"
        },
        {
            name: "grafana",
            description: "prometheus服务监控系统 (ui)，username:admin，password:adminadmin",
            picture: "./img/grafana.jpg",
            link: "http://grafana.mixbench.net"
        },
        {
            name: "Elasticsearch",
            description: "Elasticsearch 是一个分布式、RESTful 风格的搜索和数据分析引擎，能够解决不断涌现出的各种用例。作为 Elastic Stack 的核心，Elasticsearch 会集中存储您的数据，让您飞快完成搜索，微调相关性，进行强大的分析，并轻松缩放规模。",
            picture: "./img/elasticsearch.jpg",
            link: "http://elasticsearch.mixbench.net"
        },
        {
            name: "kibana",
            description: "对 Elasticsearch 索引中的数据进行搜索、查看、交互操作，已切换为中文",
            picture: "./img/kibana.jpg",
            link: "http://kibana.mixbench.net"
        },
        {
            name: "asynqmon",
            description: "go基于redis实现的消息队列asynq及管理面板 (ui)",
            picture: "./img/asynqmon.jpg",
            link: "http://asynqmon.mixbench.net"
        },
        {
            name: "CMAK",
            description: "CMAK（以前称为 Kafka Manager）是用于管理Apache Kafka集群的工具",
            picture: "./img/cmak.jpg",
            link: "http://kafka.mixbench.net"
        },
    ];
    tempE = document.getElementById("template");
    appE = document.getElementById("app");
    ulE = document.createElement("ul");
    constructor() {
        this.ulE.className="cards";
        this.buildCard();
        this.addEvent();
    }
    buildCard = ()=>{
        let tmps = this.list.map(item => {
            let html_text = this.tempE.innerHTML;
            Object.keys(item).forEach(key => {
                html_text = html_text.replace(`__${key}__`, item[key]);
            });
            return html_text;
        });
        this.ulE.insertAdjacentHTML("beforeend", tmps.join(""));
        this.appE.appendChild(this.ulE);
    }
    addEvent = ()=>{
        let btns = document.querySelectorAll(".card__btn");
        btns.forEach((btn)=>{
            btn.addEventListener("click",function(){
                window.open(this.dataset.link)
            })
        });
    }
}
new App;