<!DOCTYPE HTML>
<html>
        <head>
            <title>秒Po-北美留学生极简信息发布平台</title>
            <meta charset="UTF-8">
            <meta http-equiv="X-UA-Compatible" content="IE=edge">
            <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
            <link rel="stylesheet" type="text/css" href="/static/plugin/bootstrap/css/bootstrap.min.css">
            <link rel="stylesheet" type="text/css" href="/static/css/common.css">
            <link rel="stylesheet" type="text/css" href="/static/css/main.css">  
         </head>
         <body>
                <div class="container">
                        <div class="page-header row">
                                <div class="logo col-md-1">
                                       <a href="/">秒Po</a>
                                </div>
                                <div class="slogan col-md-10">
                                    <div class="row">
                                        <div class="col-md-5">
                                            <div class="input-group">
                                              <input type="text" class="form-control" placeholder="搜索, eg.单人间">
                                              <span class="input-group-btn">
                                                <button class="btn btn-default" type="button"><span class="glyphicon glyphicon-search" aria-hidden="true"></span></button>
                                              </span>
                                            </div>
                                        </div>
                                        <div class="col-md-5">
                                            {{range .cats}}
                                                <a href="/info/list?cid={{.Id}}" class="search_cats">{{.Name}}</a> 
                                            {{end}}
                                        </div>
                                        <div class="col-md-2">
                                            <a href="/info/create" class="btn btn-primary active" role="button">发布</a>
                                        </div>
                                    </div>
                                </div>
                        </div>
                        <div class="content">
                                {{.LayoutContent}}
                        </div>
                </div>

               

                <script src="/static/plugin/jquery/jquery-2.2.4.js"></script>
                <script src="/static/plugin/bootstrap/js/bootstrap.min.js"></script>
                <script src="/static/plugin/bootstrap/js/bootstrap3-typeahead.min.js"></script>
                <script src="/static/js/common.js"></script>
                <script src="/static/js/main.js"></script>
         </body>
</html>