<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <title>短网址生成</title>
    <!-- Bootstrap -->
    <link href="/static/bootstrap-3.3.7-dist/css/bootstrap.min.css" rel="stylesheet">
    <link href="/static/css/style.css" rel="stylesheet">
</head>
<body>
<nav class="navbar navbar-inverse navbar-fixed-top">
    <div class="container">
        <div class="navbar-header">
            <button type="button" class="navbar-toggle collapsed" data-toggle="collapse" data-target="#navbar"
                    aria-expanded="false" aria-controls="navbar">
                <span class="sr-only">Toggle navigation</span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
                <span class="icon-bar"></span>
            </button>
            <a class="navbar-brand" href="#">跨考教育</a>
        </div>
        <div id="navbar" class="collapse navbar-collapse">
            <ul class="nav navbar-nav">
                <li class="active"><a href="#">首页</a></li>
                <li><a href="#about">报表</a></li>
            </ul>
        </div><!--/.nav-collapse -->
    </div>
</nav>
<div class="container body">
    <div class="row ">
        <div class="col-md-8 col-md-offset-2">
            <div class="box box-success">
                <div class="page-header">
                    <h4>输入将要缩短的长网址
                    </h4>
                </div>
                <div class="box-body">
                    <form action=""  method="post">
                        <div class="input-group input-group-lg">
                            <input type="text" name="url" class="form-control" value="" />
                            <span class="input-group-btn">
                                <button type="submit" class="btn btn-success btn-flat">获取短网址</button>
                            </span>
                        </div>
                        <blockquote style="margin-top: 10px; word-break: break-all; word-wrap: break-word;">
                            <footer>短网址：{{.shortUrl}}</footer>
                        </blockquote>
                        <hr>
                        <p>短网址二维码</p>
                        <div id="qrcode" style="width:240px;height:240px;margin: 0 auto;">
                            {{if .shortUrl}}
                            <img src="/qrcode?url={{.shortUrl}}" alt="">
                            {{end}}
                        </div>
                    </form>
                </div>
            </div>
        </div>
    </div>
</div>

</body>
</html>