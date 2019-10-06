<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <title>机器视图</title>

    <!-- Bootstrap -->
    <!-- Latest compiled and minified CSS -->
    <link rel="stylesheet" href="/static/bootstrap/3.4.1/css/bootstrap.min.css">
    <link rel="stylesheet" href="/static/bootstrap-table/bootstrap-table.min.css"/>

    <!-- HTML5 shim and Respond.js for IE8 support of HTML5 elements and media queries -->
    <!-- WARNING: Respond.js doesn't work if you view the page via file:// -->
    <!--[if lt IE 9]>
    <script src="https://oss.maxcdn.com/html5shiv/3.7.3/html5shiv.min.js"></script>
    <script src="https://oss.maxcdn.com/respond/1.4.2/respond.min.js"></script>
    <![endif]-->
</head>
<body>
    <div class="container">
        <table id="myMac" class="table table-hover"></table>
    </div>
    {{template "footer.html" .}}
    <script>
        function ViewData() {
            $("#myMac").bootstrapTable("destroy");
            $("#myMac").bootstrapTable({
                method: 'post',
                contentType: "application/json",//必须要有！！！！
                url: "/Machine",                //要请求数据的文件路径
                height: 'auto',                 //高度调整
                toolbar: '#toolbar',            //指定工具栏
                striped: true,                  //是否显示行间隔色
                dataField: "res",//bootstrap table 可以前端分页也可以后端分页，这里
                //我们使用的是后端分页，后端分页时需返回含有total：总记录数,这个键值好像是固定的
                //rows： 记录集合 键值可以修改  dataField 自己定义成自己想要的就好
                pageNumber: 1, //初始化加载第一页，默认第一页
                pagination: true,//是否分页
                queryParamsType: 'limit',//查询参数组织方式
                queryParams: "",//请求服务器时所传的参数
                sidePagination: 'server',//指定服务器端分页
                showRefresh: true,//刷新按钮
                showColumns: true,
                sortable:   true,                       //是否启用排序
                sortOrder: "asc",                       //排序方式
                clickToSelect: true,//是否启用点击选中行
                toolbarAlign: 'right',//工具栏对齐方式
                buttonsAlign: 'right',//按钮对齐方式
                columns: [
                    {
                        title: 'IP地址',
                        field: 'IP',
                        width: '100px',
                    },
                    {
                        title: 'Mac地址',
                        field: 'Mac',
                        width: '200px',
                    },
                    {
                        title: '最后心跳时间',
                        field: 'Time',
                        width: '150px',
                    },
                    {
                        title: '用户名密码',
                        field: 'uid',
                        width: '200px',
                    },
                    {
                        title: '当前任务税号',
                        field: 'TaxpayerId',
                        width: '200px',
                    },
                    {
                        title: '任务ID',
                        field: 'TableSetId',
                        width: '40px',
                    },
                    {
                        title: '任务类型',
                        field: 'TaskType',
                        width: '40px',
                    }
                ],
                locale: 'zh-CN',//中文支持,
                responseHandler: function (res) {    //在ajax获取到数据，渲染表格之前，修改数据源
                    return res;
                }
            });
        }
        window.onload = ViewData;
        //ViewData();
        setInterval(ViewData,30000);
    </script>
    <!-- jQuery (necessary for Bootstrap's JavaScript plugins) -->
    <script src="/static/jquery/3.4.1/jquery.min.js"></script>
    <!-- Include all compiled plugins (below), or include individual files as needed -->
    <!-- Latest compiled and minified JS -->
    <script src="/static/bootstrap/3.4.1/js/bootstrap.min.js"></script>
    <script src="/static/bootstrap-table/bootstrap-table.min.js"></script>
    <script src="https://cdn.bootcss.com/jquery.serializeJSON/2.9.0/jquery.serializejson.min.js"></script>
</body>
</html>