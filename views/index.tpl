<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">
    <title>神州云合--数据分析</title>
    <!-- Bootstrap core CSS -->
    <link href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
    <script src="https://cdn.bootcss.com/echarts/4.0.0/echarts.min.js"></script>
    <style type="text/css">
        .container {
            width: 100%;
        }
    </style>
</head>

<body>
<div class="container">
    <ul class="nav nav-tabs">
        <li role="presentation" class="active"><a href="/index">首页</a></li>
        <li role="presentation"><a href="/task/day">任务【天】</a></li>
        <li role="presentation"><a href="/task/month">任务【月】</a></li>
        <li class="dropdown"><a href="#" class="dropdown-toggle" data-toggle="dropdown" role="button"
                                aria-haspopup="true" aria-expanded="false">省份<span class="caret"></span></a>
            <ul class="dropdown-menu">
                <li><a href="/task?taskid=3700301">山东</a></li>
                <li><a href="/task?taskid=3400301">安徽</a></li>
                <li><a href="/task?taskid=3100301">上海</a></li>
                <li><a href="/task?taskid=1100301">北京</a></li>
                <li><a href="/task?taskid=3100301">上海</a></li>
                <li><a href="/task?taskid=3100301">上海</a></li>
                <li><a href="/task?taskid=3100301">上海</a></li>
                <li><a href="/task?taskid=3100301">上海</a></li>
                <li><a href="/task?taskid=3100301">上海</a></li>
                <li><a href="/task?taskid=3100301">上海</a></li>
                <li><a href="/task?taskid=3100301">上海</a></li>
                <li><a href="/task?taskid=3100301">上海</a></li>
                <li><a href="/task?taskid=3100301">上海</a></li>
            </ul>
        </li>
    </ul>
</div>
<div class="jumbotron">
    <div class="container">
        <div id="myCarousel" class="carousel slide">
            <!-- 轮播（Carousel）指标 -->
            <ol class="carousel-indicators">
                <li data-target="#myCarousel" data-slide-to="0" class="active"></li>
                <li data-target="#myCarousel" data-slide-to="1"></li>
                <li data-target="#myCarousel" data-slide-to="2"></li>
            </ol>
            <!-- 轮播（Carousel）项目 -->
            <div class="carousel-inner">
                <div class="item active">
                    <div id="success" class="center-block" style="width: 1200px;height:750px;"></div>
                </div>
                <div class="item">
                    <div id="hzw" class="center-block" style="width: 1200px;height:750px;"></div>
                </div>
                <div class="item">
                    <div id="jzt" class="center-block" style="width: 1200px;height:750px;"></div>
                </div>
                <div class="item">
                    <div id="task" class="center-block" style="width: 1200px;height:750px;"></div>
                </div>
                <div class="item">
                    <div id="taxpayer" class="center-block" style="width: 1200px;height:750px;"></div>
                </div>
            </div>
            <!-- 轮播（Carousel）导航 -->
            <a class="carousel-control left" href="#myCarousel" data-slide="prev">
                <span _ngcontent-c3=""
                      aria-hidden="true"
                      class="glyphicon glyphicon-chevron-left"></span></a>
            <a class="carousel-control right" href="#myCarousel" data-slide="next">
                <span _ngcontent-c3=""
                      aria-hidden="true"
                      class="glyphicon glyphicon-chevron-right"></span></a>
        </div>
    </div>
</div>
<hr>
<footer>
    <nav class="navbar navbar-inverse ">
        <div class="container">
            <div class="navbar-header">
                <p><a class="navbar-brand" href="#">Hello {{.Name}}, Welcome to go 数据分析...</a></p>
            </div>
        </div>
    </nav>
</footer>
</div>
<!-- /container -->
<script src="/static/jquery/3.4.1/jquery.min.js"></script>
<script src="/static/bootstrap/3.4.1/js/bootstrap.min.js"></script>
<script type="text/javascript">
    function InitChartsEx(id, group) {
        URI = "";
        if (group.length) {
            URI = "/statisticalEx?group=" + group;
        } else {
            URI = "/statisticalEx";
        }
        $.post(URI).done(function (data) {
            echarts.init(document.getElementById(id)).setOption({
                title: {
                    text: data.title,
                    top: '30',
                    subtext: '数据来自申报网关'
                },
                tooltip: {
                    trigger: 'axis',
                },
                grid: {
                    top: '100',
                    left: '3%',
                    right: '4%',
                    bottom: '5%',
                    containLabel: true
                },
                legend: {
                    data: data.legend,
                    orient: 'vertical',
                    padding: 5,
                    top: 20,
                    left: '70%'
                },
                xAxis: {
                    type: 'category',
                    data: data.categories
                },
                yAxis: {
                    type: 'value',
                },
                series: data.series
            });
        });
    }

    function InitChart(id) {
        $.post('/statistical/' + id).done(function (data) {
                var tt = echarts.init(document.getElementById(id));
                tt.setOption(data);
                tt.on('click', function (params) {
                    console.log(params.name);
                });
            }
        );
    }

    InitChart("task");
    InitChart("taxpayer");
    InitChartsEx("success", "");
    InitChartsEx("hzw", "孩子王");
    InitChartsEx("jzt", "九州通医药集团");
    // InitChartsEx("gj", "高济集团");
    // InitChartsEx("fdc", "百安居");
</script>
</body>
</html>