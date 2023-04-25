

安装prometheus + grafana    mac下
https://cloud.tencent.com/developer/article/1690610

启动： ./prometheus --config.file=/usr/local/etc/prometheus.yml

添加环境变量
https://blog.csdn.net/cc18868876837/article/details/126388438

启动： grafana-server --config=/usr/local/etc/grafana/grafana.ini --homepath /usr/local/share/grafana --packaging=brew cfg:default.paths.logs=/usr/local/var/log/grafana cfg:default.paths.data=/usr/local/var/lib/grafana cfg:default.paths.plugins=/usr/local/var/lib/grafana/plugins


安装pushgateway
./pushgateway 这里解压之后打开会说未验证  需要在mac的设置-》隐私-》通用里面设置允许
https://blog.csdn.net/Liiipseoroinis/article/details/119277546
向pushgatway中添加数据
echo "some_test 1219" | curl --data-binary @- http://localhost:9091/metrics/job/some_job

上面这个pushgetway其实可以装可以不装，这个一般只是用在短命服务中左右个数据的通道，一般的服务是直接发送数据到prometheus



配置prometheus抓取本机的指标数据
  - job_name: 'my_job': my_job指的是采集指标的名称

scrape_configs:
  - job_name: 'my_job'
    static_configs:
      - targets: ['localhost:8080']







