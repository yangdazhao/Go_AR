<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <!-- The above 3 meta tags *must* come first in the head; any other head content must come *after* these tags -->
    <title>定位一下</title>

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
    <form action="./" method="POST" role="form" id="taskinfo">
        <legend>任务详情查询</legend>
        <div class="row">
            <div class="col-xs-4 col-sm-4 col-md-4 col-lg-4">
                <div class="form-group">
                    <label for="">流水号或者公司名称或者税号</label>
                    <input 
						type="text" 
						class="form-control" 
						name="serialnumber" 
						id="serialnumber"
                        placeholder="流水号长度为32位">
                </div>
            </div>
            <div class="col-xs-3 col-sm-3 col-md-3 col-lg-3">
                <button type="button" class="btn btn-primary" onclick="QueryEx()">定位一下</button>
            </div>
        </div>
    </form>
    <table id="mytab" class="table table-hover"></table>
</div>

    <div class="footer">
        <a src="https://pan.bigfintax.com/s/cWhi9cRYqg6W6rN">工具下载</a>
    </div>
<script>
function checknum(value) {
    var Regx = /^[A-Za-z0-9]*$/;
    if (Regx.test(value)) {
        return true;
    }else {
        return false;
    }
}

function QueryEx()
{
	var KeyWord = jQuery("#serialnumber").val().trim();
	if(checknum(KeyWord))
	{
		if( KeyWord.length == 32)
		{
			query({"serialnumber":KeyWord});
		}
		else if( KeyWord.length == 18 ){
			query({"taxpayerid":KeyWord});
		}
		else{
			query({});
		}
	}
	else
	{
		query({"taxpayerid":KeyWord});
	}
}
		
function query(Params) {
    $('#mytab').bootstrapTable({
        method: 'post',
        contentType: "application/json",//必须要有！！！！
        url: "TaskQuery",//要请求数据的文件路径
        height: 'auto',//高度调整
        toolbar: '#toolbar',//指定工具栏
        striped: true, //是否显示行间隔色
        dataField: "res",//bootstrap table 可以前端分页也可以后端分页，这里
        //我们使用的是后端分页，后端分页时需返回含有total：总记录数,这个键值好像是固定的
        //rows： 记录集合 键值可以修改  dataField 自己定义成自己想要的就好
        pageNumber: 1, //初始化加载第一页，默认第一页
        pagination: true,//是否分页
        queryParamsType: 'limit',//查询参数组织方式
        queryParams: Params,//请求服务器时所传的参数
        sidePagination: 'server',//指定服务器端分页
        pageSize: 10,//单页记录数
        pageList: [5, 10, 20, 30],//分页步进值
        showRefresh: false,//刷新按钮
        showColumns: true,
        clickToSelect: true,//是否启用点击选中行
        toolbarAlign: 'right',//工具栏对齐方式
        buttonsAlign: 'right',//按钮对齐方式
        toolbar: '#toolbar',//指定工作栏
        columns: [
            {
                title: '序列号',
                field: 'SerialNumber',
                width: '100px',
                //sortable: true
            },
            {
                title: '公司名',
                field: 'Company.CompanyName',
                //sortable: true
            },
            {
                title: '税种id',
                field: 'TableSetID',
            },
            {
                title: '开始时间',
                field: 'Created',
                width: '200px',
            },
            {
                title: '任务类型',
                field: 'Type',
                width: '40px',
            },
            {
                title: '状态',
                field: 'Message',
                align: 'center',
				width: '40px',
            },
            {
                title: '详情',
                field: 'TaskJson',
                align: 'center',
                formatter: aFormatter //添加超链接的方法
            }
        ],
        locale: 'zh-CN',//中文支持,
        responseHandler: function (res) {    //在ajax获取到数据，渲染表格之前，修改数据源
            return res;
        }
    });

    //三个参数，value代表该列的值
    function aFormatter(value, row, index) {
		if(value.length){
		var uri = "https://cabinet.bigfintax.com/";
        uri += value;
        return ['<a href="'+ uri +'">Json</a> | <a href="'+ uri +'">日志</a>'].join("")
		}
		return "";
    }

    //请求服务数据时所传参数
    function queryParams(params) {
        return {
            //每页多少条数据
            pageSize: params.limit,
            //请求第几页
            pageIndex: params.pageNumber,
            Name: $('#search_name').val(),
            Tel: $('#search_tel').val()
        }
    }

    //查询按钮事件
    $('#search_btn').click(function () {
        $('#mytab').bootstrapTable('refresh', {url: '../index.php/admin/index/userManagement'});
    })
}
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