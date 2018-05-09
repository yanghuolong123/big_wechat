<!DOCTYPE HTML>
<html>
        <head>
            {{template "layout/base-style.tpl" .}}
         </head>
         <body>
                <div class="container">
                        <div class="page-header row">
                                <div class="logo col-md-2">
                                       <a href="/">秒Po-UTD</a>
                                </div>
                                <div class="slogan col-md-8">
                                        <span>中国学生极简信息发布平台</span>
                                </div>
                                <div class="col-md-2">
                                {{if .user}}
                                    <a  href="/user" class="btn btn-info">
                                                      <span class="glyphicon glyphicon-user" aria-hidden="true"></span>  {{.user.Nickname}}
                                    </a>  
                                {{else}}
                                    <button id="loginBtn" class="btn btn-info" type="button">
                                              <span aria-hidden="true" class="glyphicon glyphicon-log-in"></span> 登录
                                    </button>
                                {{end}}
                                </div>
                        </div>
                        <div class="content row">
                            <div class="col-md-9">{{.LayoutContent}}</div>
                             <div class="col-md-3">{{template "layout/side.tpl" .}}</div>   
                        </div>
                </div>

                {{template "login/login.tpl" .}} 
                {{template "layout/footer.tpl" .}}
                {{template "layout/base-script.tpl" .}}
                {{if .isWeixin}}
                {{template "layout/wxshare.tpl" .}}
                {{end}}
         </body>
</html>