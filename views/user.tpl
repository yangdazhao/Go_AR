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
          width: 90%;
        }
     </style>
  </head>

  <body>
    <nav class="navbar navbar-inverse navbar-fixed-top">
      <div class="container">
        <div class="navbar-header">
          <a class="navbar-brand" href="#">数据统计分析</a>
        </div>
      </div>
    </nav>

    <div class="jumbotron">
<!--
      <div class="container">
        <div id="main" style="width: 1200px;height:600px;"></div>
      </div>
-->
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
             <div id="task" class="center-block" style="width: 1200px;height:750px;"></div>
        </div>
        <div class="item">
			<div id="taxpayer" class="center-block" style="width: 1200px;height:750px;"></div>
        </div>
    </div>
    <!-- 轮播（Carousel）导航 -->
  <a class="carousel-control left" href="#myCarousel" data-slide="prev"> 
  <span _ngcontent-c3="" aria-hidden="true" class="glyphicon glyphicon-chevron-left"></span></a>
  <a class="carousel-control right" href="#myCarousel" data-slide="next"><span _ngcontent-c3="" aria-hidden="true" class="glyphicon glyphicon-chevron-right"></span></a>
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
    </div> <!-- /container -->
	<script src="https://cdn.bootcss.com/jquery/3.4.0/jquery.min.js"></script>
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
	<script type="text/javascript">
            // 基于准备好的dom，初始化echarts实例
            var myChart = echarts.init(document.getElementById('success'));
			myChart.setOption(
			{
				title: {
					top:'30',
					subtext: '数据来自申报网关'
				},
				tooltip: {
					trigger: 'axis',
				},
				grid: {
					top:'100',
					left: '3%',
					right: '4%',
					bottom: '5%',
					containLabel: true
				},
				xAxis: {
					type: 'category',
				},
				yAxis: {
					type: 'value',
				},
				series: [
				],
			}
            );
		myChart.on('click', function (params) {
			alert(JSON.stringify(params));
    if (params.componentType === 'markPoint') {
        // 点击到了 markPoint 上
        if (params.seriesIndex === 5) {
            // 点击到了 index 为 5 的 series 的 markPoint 上。
        }
    }
    else if (params.componentType === 'series') {
		alert(params.seriesType);
    }
});
	 //异步加载数据
	$.post('/statisticalEx').done(function (data) {
		// 填入数据
		myChart.setOption({
		    title: {
		        text: data.title,
            },
             legend: {
                data: data.legend,
                orient:'vertical',
                padding: 5,
                top: 20,
                left: '70%'
             },
			xAxis: {
				data: data.categories
			},
			series:  data.series
		});
		

		//alert(data.series);
	});
		
		
		function InitChart( id )
		{
			// 基于准备好的dom，初始化echarts实例
            var taxpayerChart = echarts.init(document.getElementById(id));
			taxpayerChart.setOption(
			{
            				title: {
            					subtext: '数据来自申报网关'
            				},
            				tooltip: {
            					trigger: 'axis',
            					axisPointer: {
            						type: 'shadow'
            					}
            				},
            				legend: {
            				},
            				grid: {

            					left: '10%',
            					right: '20%',
            					bottom: '3%',
            					containLabel: true
            				},
            				xAxis: {
            					type: 'value',
            				},
            				yAxis: {
            					type: 'category',
            				},
				    		//dataZoom: [        {            type: 'inside'        }    ],
            				series: [
            					{
            						type: 'bar',
									//itemStyle: {                						normal: {color: 'rgba(0,0,0,0.05)'}						}
            					}

            				],
            			}
            );
			
			
	// 异步加载数据
	$.post('/statistical/' + id).done(function (data) {
		// 填入数据
		taxpayerChart.setOption({
		    title: {
		        text:data.title,
            }			,
			yAxis: {
				data: data.categories
			},
			series: [{
				name: data.name,
				data: data.data
			}]
		});
	});
		}
		InitChart("task");
		InitChart("taxpayer");
		

        </script>
  </body>
</html>