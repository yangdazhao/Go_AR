<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>申报网关</title>
        <!-- Bootstrap core CSS -->
        <link href="//netdna.bootstrapcdn.com/twitter-bootstrap/2.3.2/css/bootstrap-combined.min.css" rel="stylesheet">
        <link href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" rel="stylesheet">
            <style type="text/css">
              body {
                padding-top: 40px;
                padding-bottom: 40px;
                background-color: #f5f5f5;
              }

              .form-signin {
                max-width: 300px;
                padding: 19px 29px 29px;
                margin: 0 auto 20px;
                background-color: #fff;
                border: 1px solid #e5e5e5;
                -webkit-border-radius: 5px;
                   -moz-border-radius: 5px;
                        border-radius: 5px;
                -webkit-box-shadow: 0 1px 2px rgba(0,0,0,.05);
                   -moz-box-shadow: 0 1px 2px rgba(0,0,0,.05);
                        box-shadow: 0 1px 2px rgba(0,0,0,.05);
              }
              .form-signin .form-signin-heading,
              .form-signin .checkbox {
                margin-bottom: 10px;
              }
              .form-signin input[type="text"],
              .form-signin input[type="password"] {
                font-size: 16px;
                height: auto;
                margin-bottom: 15px;
                padding: 7px 9px;
              }

              #login img{
                margin: 10px 0;
              }
              #login .center {
                text-align: center;
              }

              #login .login {
                max-width: 300px;
              	margin: 35px auto;
              }

              #login .login-form{
                padding:0px 25px;
              }
            </style>
</head>
<body>
    <div id="login" class="container">
      <div class="row-fluid">
        <div class="span12">
          <div class="login well well-small">
            <div class="center">
              <img src="http://www.easydatalink.com/images/logo0.png" alt="logo">
            </div>
            <h2 class="form-signin-heading">请登录</h2>
            <form action="/home/login" style="" class="login-form" id="UserLoginForm" method="post" accept-charset="utf-8">
              <div class="control-group">
                <div class="input-prepend">
                  <span class="add-on"><i class="icon-user"></i></span>
                  <input name="Username" required="required" placeholder="Username" maxlength="255" type="text" id="Username">
                </div>
              </div>
              <div class="control-group">
                <div class="input-prepend">
                  <span class="add-on"><i class="icon-lock"></i></span>
                  <input name="Password" required="required" placeholder="Password" type="password" id="Password">
                </div>
              </div>
              <div class="control-group">
                <label id="remember-me">
                  <input type="checkbox" name="data[User][remember_me]" value="1" id="UserRememberMe"> Remember Me?</label>
              </div>
              <div class="control-group">
                <input class="btn btn-primary btn-large btn-block" type="submit" value="登录">
              </div>
            </form>
          </div><!--/.login-->
        </div><!--/.span12-->
      </div><!--/.row-fluid-->
    </div><!--/.container-->
    <script src="https://cdn.bootcss.com/jquery/3.4.0/jquery.min.js"></script>
    <script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
</body>
</html>