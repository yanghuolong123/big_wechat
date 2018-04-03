<!DOCTYPE HTML>
<html>
        <head>
            <title>秒Po-中国学生极简信息发布平台</title>
            <meta charset="UTF-8">
            <meta http-equiv="X-UA-Compatible" content="IE=edge">
            <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
            <meta name="keywords" content="“秒Po”是中国学生的极简信息发布平台。平台以各校区独立运营、免注册极简发布等模式，让学生间的信息交互更加简洁、高效和相对安全，有效解决聊天群信息分数、大众平台繁琐杂乱等问题。解决学生间的买卖二手，房屋租赁，求租，出租单间，卖车，买车，办卡，开电灯问题。常见二手物品包括床垫，床架，书桌，台灯，洗衣机，烘干机等。我们将会在主要学校开通秒Po，包括UTD,UT, UIUC, UW,SMU, USC, PU, NEU, Columbia, OSU, UCLA, Indiana University, Berkeley, NYU, PSU, ASU, UMAA, Boston, IIT, Rutgers.">
            <link rel="stylesheet" type="text/css" href="/static/plugin/bootstrap/css/bootstrap.min.css">
            <link rel="stylesheet" type="text/css" href="/static/css/common.css">
            <link rel="stylesheet" type="text/css" href="/static/css/main.css">  
         </head>
         <body>
                <div class="container">
                        <div class="page-header row">
                                <div class="logo col-md-2">
                                       <a href="/">秒Po-UTD</a>
                                </div>
                                <div class="slogan col-md-10">
                                    <div class="row">
                                        <div class="col-md-5">
                                            <div class="input-group">
                                              <input type="text" class="form-control" id="search" placeholder="搜索, eg. 室友，moving">
                                              <span class="input-group-btn">
                                                <button class="btn btn-default search-btn" type="button"><span class="glyphicon glyphicon-search" aria-hidden="true"></span></button>
                                              </span>
                                            </div>
                                        </div>
                                        <div class="col-md-5">
                                            {{range .cats}}
                                                <a href="/info/list?cid={{.Id}}" class="search_cats">{{.Name}}</a> 
                                            {{end}}
                                        </div>
                                        <div class="col-md-2">
                                            <div class="dropdown">
                                                  <button class="btn btn-primary dropdown-toggle" type="button" id="publish-btn" data-toggle="dropdown" aria-haspopup="true" aria-expanded="true">
                                                    发布    
                                                    <span class="caret"></span>
                                                  </button>
                                                  <ul class="dropdown-menu" aria-labelledby="publish-btn">
                                                    {{range .cats}}
                                                     <li><a href="/info/create?cid={{.Id}}">{{.Name}}</a></li>
                                                    {{end}}
                                                  </ul>
                                            </div>
                                            <!--<a href="/info/create" class="btn btn-primary active" role="button">发布</a>-->
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