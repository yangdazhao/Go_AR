<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>任务统计</title>
    <script src="/static/js/echarts.min.js"></script>
    <script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.js"></script>
</head>
<body>
 <!-- 为 ECharts 准备一个具备大小（宽高）的 DOM -->
    <div id="main" style="width: 1600px;height:800px;"></div>
        <script type="text/javascript">
            // 基于准备好的dom，初始化echarts实例
            var myChart = echarts.init(document.getElementById('main'));
			myChart.setOption(
			{
				title: {
					subtext: '数据来自申报网关'
				},
				tooltip: {
					trigger: 'axis',
				},
				grid: {
					left: '3%',
					right: '4%',
					bottom: '3%',
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
			// 异步加载数据
	// $.post('{{.Param}}').done(function (data) {
	$.post('/statisticalEx').done(function (data) {
		// 填入数据
		myChart.setOption({
		    title: {
		        text: data.title,
            },
             legend: {
                data: data.legend,
             },
			xAxis: {
				data: data.categories
			},
			series:  data.series
		});
	});
        </script>	
</body>
</html>