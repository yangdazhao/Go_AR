<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">
    <title>Bootstrap Template Page for Go Web Programming</title>
    <!-- Bootstrap core CSS -->
    <link href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
  </head>

  <body>

    <nav class="navbar navbar-inverse navbar-fixed-top">
      <div class="container">
        <div class="navbar-header">
          <a class="navbar-brand" href="#">YangDazhao general infor</a>
        </div>
      </div>
    </nav>

    <div class="jumbotron">
      <div class="container">
        <h1>Hello, {{.Name}}</h1>
        <ul>
			<li>Name   : {{.Name}}</li>
        	<li>Id     : {{.Id}}</li>
        	<li>Country: {{.Country}}</li>
        </ul>
        <p><a class="btn btn-primary btn-lg" href="#" role="button">More &raquo;</a></p>
      </div>
    </div>

    <div class="container">
      <div class="row">
        <div class="col-md-4">
          <h2>Name</h2>
          <p>Name has the value of : {{.Name}} </p>
          <p><a class="btn btn-default" href="#" role="button">More &raquo;</a></p>
        </div>
        <div class="col-md-4">
          <h2>Id</h2>
          <p>Id has the value of : {{.Id}} </p>
          <p><a class="btn btn-default" href="#" role="button">More &raquo;</a></p>
       </div>
        <div class="col-md-4">
          <h2>Country</h2>
          <p>Country has the value of : {{.Country}} </p>
          <p><a class="btn btn-default" href="#" role="button">More &raquo;</a></p>
        </div>
      </div>
      <hr>
      <footer>
      <nav class="navbar navbar-inverse ">
        <div class="container">
          <div class="navbar-header">
            <p><a class="navbar-brand" href="#">Hello, {{.Name}}, Welcome to go 数据分析...</a></p>
          </div>
        </div>
      </nav>
      </footer>
    </div> <!-- /container -->
	<script src="https://cdn.bootcss.com/jquery/3.4.0/jquery.min.js"></script>
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
  </body>
</html>