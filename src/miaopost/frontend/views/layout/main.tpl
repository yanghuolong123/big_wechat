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
                                <div class="slogan col-md-10">
                                        <span>中国学生极简信息发布平台</span>
                                </div>
                        </div>
                        <div class="content row">
                            <div class="col-md-9">{{.LayoutContent}}</div>
                             <div class="col-md-3">{{template "layout/side.tpl" .}}</div>   
                        </div>
                </div>

               
                {{template "layout/footer.tpl" .}}
                {{template "layout/base-script.tpl" .}}
                {{template "layout/wxshare.tpl" .}}
         </body>
</html>